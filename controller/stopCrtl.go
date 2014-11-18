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

type StopController struct { }


////////////////////////////////////////////////////////////////////////////////////////////////
/// Variables
////////////////////////////////////////////////////////////////////////////////////////////////

var (
	stopRepository database.GTFSStopRepository
)


////////////////////////////////////////////////////////////////////////////////////////////////
/// Stop Controller
////////////////////////////////////////////////////////////////////////////////////////////////

func (stopController *StopController) Init(r *mux.Router) {
	stopRepository = config.GTFS.Stops().(database.GTFSStopRepository)

	r.HandleFunc("/", stopController.Stops)
}

func (ac *StopController) Stops(w http.ResponseWriter, r *http.Request) {
	stops, err := stopRepository.FindAll()

	if err != nil {
		http.Error(w, err.Error(), 500)
	} else if stops == nil {
		http.Error(w, "No stop found", 500)
	} else {
		utils.SendJSON(w, stops.ToJSONStops())
	}
}
