package controller

import (
	"net/http"

	"github.com/RobinBaeckman/hermod-api/application/admin/product/presenter"
	"github.com/RobinBaeckman/hermod-api/infra/mongo"
	"github.com/RobinBaeckman/hermod-api/infra/web/admin/product/view"
	"github.com/RobinBaeckman/hermod-api/usecase/admin/product"
	"github.com/RobinBaeckman/hermod-api/validate"
	"github.com/gorilla/sessions"
	mgo "gopkg.in/mgo.v2"
)

type App struct {
	CStore *sessions.CookieStore
	DB     *mgo.Database
}

func NewInteractor(db *mongo.ProductDB, w http.ResponseWriter) *product.Interactor {
	return &product.Interactor{
		product.Repository(db),
		presenter.Presenter{view.Viewer{w}},
	}
}

func (a *App) Store(w http.ResponseWriter, r *http.Request) (err error) {
	db := a.DB.With(a.DB.Session.Copy())
	defer db.Session.Close()
	i := NewInteractor(&mongo.ProductDB{db}, w)

	ind, err := mapStoreRequest(r)
	if err != nil {
		return err
	}

	if err := validate.Check(&ind); err != nil {
		return err
	}

	if err := i.Store(ind); err != nil {
		return err
	}

	return
}

func (a *App) Show(w http.ResponseWriter, r *http.Request) (err error) {
	db := a.DB.With(a.DB.Session.Copy())
	defer db.Session.Close()
	i := NewInteractor(&mongo.ProductDB{db}, w)
	ind, err := mapShowRequest(r)
	if err != nil {
		return err
	}

	if err := validate.Check(&ind); err != nil {
		return err
	}

	if err := i.Show(ind); err != nil {
		return err
	}

	return nil
}

func (a *App) Index(w http.ResponseWriter, r *http.Request) (err error) {
	db := a.DB.With(a.DB.Session.Copy())
	defer db.Session.Close()
	i := NewInteractor(&mongo.ProductDB{db}, w)
	if err := i.Index(); err != nil {
		return err
	}

	return
}
