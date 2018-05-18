package view

import (
	"encoding/json"
	"net/http"

	"github.com/RobinBaeckman/hermod-api/application/admin/auth/presenter"
	"github.com/gorilla/sessions"
	"github.com/spf13/viper"
)

type Viewer struct {
	W     http.ResponseWriter
	R     *http.Request
	Store *sessions.CookieStore
}

func (v Viewer) View(vm presenter.ViewModel) error {
	jData, err := json.Marshal(vm)
	if err != nil {
		return err
	}

	session, _ := v.Store.Get(v.R, viper.GetString("session.cookie_name"))
	session.Values["authenticated"] = true
	session.Values["role"] = "admin"
	session.Save(v.R, v.W)
	v.W.Header().Set("Content-Type", "application/json")
	v.W.Write(jData)

	return err
}
