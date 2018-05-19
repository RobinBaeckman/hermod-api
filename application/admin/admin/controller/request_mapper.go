package controller

import (
	"encoding/json"
	"net/http"

	"github.com/RobinBaeckman/hermod-api/usecase/admin/admin"
	"github.com/gorilla/mux"
)

func mapStoreRequest(r *http.Request) (admin.InputDataStore, error) {
	ind := admin.InputDataStore{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&ind)
	if err != nil {
		return ind, err
	}

	defer r.Body.Close()

	return ind, nil
}

func mapShowRequest(r *http.Request) (admin.InputDataShow, error) {
	vars := mux.Vars(r)
	id := vars["id"]
	ind := admin.InputDataShow{
		ID: id,
	}

	return ind, nil
}
