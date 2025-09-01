package api

import (
	"encoding/json"
	"net/http"
	"simengine/metrics"
	"simengine/simulator"
	"simengine/types"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func handleSimulation(w http.ResponseWriter, r *http.Request) {
	var problem types.Problem

	if err := json.NewDecoder(r.Body).Decode(&problem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	if err := validate.Struct(problem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	metric, err := simulator.EngineRun(problem)

	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	
	res, err := metrics.MetricsAggregator(&metric)
	
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	RespondWithJSON(w, 200, res)
}