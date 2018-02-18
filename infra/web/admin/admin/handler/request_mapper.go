package handler

import (
	"encoding/json"
	"net/http"

	"github.com/RobinBaeckman/hermod-api/customerr"
	"github.com/RobinBaeckman/hermod-api/usecase/admin/admin"
	"github.com/gorilla/mux"
)

func mapCreateRequest(r *http.Request) (admin.InputData, error) {
	ind := admin.InputData{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&ind)
	if err != nil {
		return ind, &customerr.App{err, "Missing parameters", 404}
	}
	defer r.Body.Close()

	return ind, nil
}

func mapGetRequest(r *http.Request) (admin.InputData, error) {
	vars := mux.Vars(r)
	id := vars["id"]
	ind := admin.InputData{
		ID: id,
	}

	return ind, nil
}
