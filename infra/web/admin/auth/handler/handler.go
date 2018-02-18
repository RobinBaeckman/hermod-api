package handler

import (
	"encoding/json"
	"net/http"

	"github.com/RobinBaeckman/hermod-api/adapter/admin/auth/presenter"
	"github.com/RobinBaeckman/hermod-api/domain/admin"
	"github.com/RobinBaeckman/hermod-api/infra/mongo"
	"github.com/RobinBaeckman/hermod-api/infra/web/admin/auth/view"
	"github.com/RobinBaeckman/hermod-api/usecase/admin/auth"
	"github.com/gorilla/sessions"
)

var Store *sessions.CookieStore

func NewInteractor(r admin.Repository, p auth.Presenter) *auth.Interactor {
	return &auth.Interactor{
		r,
		p,
	}
}

func Auth(w http.ResponseWriter, r *http.Request) (err error) {
	db := mongo.DB.With(mongo.DB.Session.Copy())
	defer db.Session.Close()
	i := NewInteractor(
		admin.Repository(mongo.NewAuthDB(db)),
		presenter.Presenter{view.Viewer{w, r, Store}},
	)

	ind := &auth.InputData{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(ind)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	i.Auth(*ind)

	return
}

func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "cookie-name")

	// Revoke users authentication
	session.Values["authenticated"] = false
	session.Save(r, w)
}
