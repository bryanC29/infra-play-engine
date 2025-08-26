package api

import (
	"encoding/json"
	"net/http"
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

	simulator.EngineRun(problem.Design)

	RespondWithJSON(w, 200, map[string]string{"msg":"done"})
}