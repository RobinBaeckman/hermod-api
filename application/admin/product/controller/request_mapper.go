package controller

import (
	"encoding/json"

	usecase "github.com/RobinBaeckman/hermod-api/usecase/admin/product"
)

func (r Req) mapRequest() usecase.ReqBody {
	rb := usecase.ReqBody{}
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&rb)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	return rb
}
