package controller

import (
	"encoding/json"
	"net/http"

	"github.com/RobinBaeckman/hermod-api/application/admin/auth/presenter"
	"github.com/RobinBaeckman/hermod-api/domain/admin"
	"github.com/RobinBaeckman/hermod-api/infra/mongo"
	"github.com/RobinBaeckman/hermod-api/infra/web/admin/auth/view"
	"github.com/RobinBaeckman/hermod-api/usecase/admin/auth"
	"github.com/gorilla/sessions"
	"github.com/spf13/viper"
	mgo "gopkg.in/mgo.v2"
)

type App struct {
	CStore *sessions.CookieStore
	DB     *mgo.Database
}

func NewInteractor(r admin.Repository, p auth.Presenter) *auth.Interactor {
	return &auth.Interactor{
		r,
		p,
	}
}

func (a *App) Auth(w http.ResponseWriter, r *http.Request) (err error) {
	db := a.DB.With(a.DB.Session.Copy())
	defer db.Session.Close()
	i := NewInteractor(
		admin.Repository(mongo.NewAuthDB(db)),
		presenter.Presenter{view.Viewer{w, r, a.CStore}},
	)

	ind := &auth.InputData{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(ind)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if err := i.Auth(*ind); err != nil {
		return err
	}

	return
}

func (a *App) Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := a.CStore.Get(r, viper.GetString("session.cookie_name"))

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}
