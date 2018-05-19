package controller

import (
	"encoding/json"
	"net/http"

	"github.com/RobinBaeckman/hermod-api/usecase/admin/product"
	"github.com/gorilla/mux"
)

func mapStoreRequest(r *http.Request) (product.InputDataStore, error) {
	ind := product.InputDataStore{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&ind)
	if err != nil {
		return ind, err
	}

	defer r.Body.Close()

	return ind, nil
}

func mapShowRequest(r *http.Request) (product.InputDataShow, error) {
	vars := mux.Vars(r)
	id := vars["id"]
	ind := product.InputDataShow{
		ID: id,
	}

	return ind, nil
}
