package controller

import (
	"encoding/json"
	"net/http"

	"github.com/RobinBaeckman/hermod-api/usecase/admin/auth"
)

func mapAuthRequest(r *http.Request) (auth.InputDataAuth, error) {
	ind := auth.InputDataAuth{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&ind)
	if err != nil {
		return ind, err
	}

	defer r.Body.Close()

	return ind, nil
}
