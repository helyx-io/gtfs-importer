package controller

////////////////////////////////////////////////////////////////////////////////////////////////
/// Imports
////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/helyx-io/gtfs-playground/database"
	"github.com/helyx-io/gtfs-playground/utils"
	"github.com/helyx-io/gtfs-playground/config"
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Structures
////////////////////////////////////////////////////////////////////////////////////////////////

type StopTimeController struct { }


////////////////////////////////////////////////////////////////////////////////////////////////
/// Variables
////////////////////////////////////////////////////////////////////////////////////////////////

var (
	stopTimeRepository database.GTFSStopTimeRepository
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// StopTime Controller
////////////////////////////////////////////////////////////////////////////////////////////////

func (stopTimeController *StopTimeController) Init(r *mux.Router) {
	stopTimeRepository = config.GTFS.StopTimes().(database.GTFSStopTimeRepository)

	r.HandleFunc("/", stopTimeController.StopTimes)
}

func (ac *StopTimeController) StopTimes(w http.ResponseWriter, r *http.Request) {
	stopTimes, err := stopTimeRepository.FindAll()

	if err != nil {
		http.Error(w, err.Error(), 500)
	} else if stopTimes == nil {
		http.Error(w, "No stopTime found", 500)
	} else {
		utils.SendJSON(w, stopTimes.ToJSONStopTimes())
	}
}

