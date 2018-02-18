package view

import (
	"encoding/json"
	"net/http"

	"github.com/RobinBaeckman/hermod-api/adapter/admin/admin/presenter"
)

type Viewer struct {
	http.ResponseWriter
}

func (v Viewer) ViewCreated(vm *presenter.ViewModel) (err error) {
	jbs, err := json.Marshal(vm)
	if err != nil {
		return err
	}
	v.Header().Set("Content-Type", "application/json")
	v.Write(jbs)

	return
}

func (v Viewer) View(vm *presenter.ViewModel) (err error) {
	jbs, err := json.Marshal(vm)
	if err != nil {
		return err
	}
	v.Header().Set("Content-Type", "application/json")
	v.Write(jbs)

	return
}

func (v Viewer) ViewAll(vm *[]presenter.ViewModel) (err error) {
	jbs, err := json.Marshal(vm)
	if err != nil {
		return err
	}
	v.Header().Set("Content-Type", "application/json")
	v.Write(jbs)

	return
}
