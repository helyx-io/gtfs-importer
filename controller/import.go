package controller

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/fatih/stopwatch"
	"github.com/akinsella/go-playground/database/mysql"
	"github.com/akinsella/go-playground/utils"
	"github.com/akinsella/go-playground/models"
	"github.com/goinggo/workpool"

)

const (
	url = "http://localhost/data/gtfs_paris_20140502.zip"
	zipFilename = "/Users/akinsella/Desktop/gtfs_paris_20140502.zip"
	folderFilename = "/Users/akinsella/Desktop/gtfs_paris_20140502"
)

type ImportController struct { }

func (importController *ImportController) Init(r *mux.Router) {
	r.HandleFunc("/", importController.Import)
}

func (ac *ImportController) Import(w http.ResponseWriter, r *http.Request) {

	var err error

	defer func() {
		if r := recover(); r != nil {
			err, _ = r.(error)
			http.Error(w, err.Error(), 500)
			return
		}
	}()

	sw := stopwatch.Start(0)

	w.Header().Set("Content-Type", "text/html")

	w.Write([]byte("Importing agencies ..."))
	w.Write([]byte("<br/>"))

	w.Write([]byte(fmt.Sprintf(" - Downloading zip file from url: '%v' to file path: '%v' ...", url, zipFilename)))
	w.Write([]byte("<br/>"))

	writtenBytes, err := utils.DownloadFileFromURL(url, zipFilename)
	utils.FailOnError(err, fmt.Sprintf("Could not download file from url: '%v' to file path: '%v'", url, zipFilename))

	w.Write([]byte(fmt.Sprintf(" - Downloaded zip file: '%v' - %v bytes - ElapsedTime: %v", zipFilename, writtenBytes, sw.ElapsedTime())))
	w.Write([]byte("<br/>"))

	w.Write([]byte(fmt.Sprintf(" - Unzipping file: '%v' to directory: '%v' ...", zipFilename, folderFilename)))
	w.Write([]byte("<br/>"))

	swZip := stopwatch.Start(0)

	err = utils.Unzip(zipFilename, folderFilename)
	utils.FailOnError(err, fmt.Sprintf("Could unzip filename: '%v' to folder: '%v'", zipFilename, folderFilename))

	w.Write([]byte(fmt.Sprintf(" - Unzipped file: '%v' to directory: '%v' - ElapsedTime: %v - Duration: %v", zipFilename, folderFilename, sw.ElapsedTime(), swZip.ElapsedTime())))
	w.Write([]byte("<br/>"))

	stopsFilename := fmt.Sprintf("%v/stop_times.txt", folderFilename)

	w.Write([]byte(fmt.Sprintf(" - Reading file: '%v'", stopsFilename)))
	w.Write([]byte("<br/>"))

	swReadFile := stopwatch.Start(0)

	workPool := workpool.New(32, 10000)

	offset := 0

	db, err := mysql.InitDb(2, 100);
	utils.FailOnError(err, "Could not open database")
	defer db.Close()

	gtfs := mysql.CreateMySQLGTFSRepository(db)
	stopTimes := gtfs.StopTimes()
	stopTimes.RemoveAllByAgencyKey("RATP")

	gtfsFile := models.GTFSFile{stopsFilename}

	for lines := range gtfsFile.LinesIterator() {

		offset++
		taskName := fmt.Sprintf("ChunkImport-%d", offset)
		task := stopTimes.CreateImportTask(taskName, lines, workPool)

		err := workPool.PostWork("import", &task)

		utils.FailOnError(err, fmt.Sprintf("Could not post work with offset: %d", offset))
	}

	w.Write([]byte(fmt.Sprintf(" - 	Read file: '%v' - ElapsedTime: %v - Duration: %v", stopsFilename, sw.ElapsedTime(), swReadFile.ElapsedTime())))
	w.Write([]byte("<br/>"))
}
