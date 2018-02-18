package handler

import (
	"net/http"

	"github.com/RobinBaeckman/hermod-api/adapter/admin/admin/presenter"
	"github.com/RobinBaeckman/hermod-api/infra/mongo"
	"github.com/RobinBaeckman/hermod-api/infra/web/admin/admin/view"
	"github.com/RobinBaeckman/hermod-api/usecase/admin/admin"
)

func NewInteractor(db *mongo.AdminDB, w http.ResponseWriter) *admin.Interactor {
	return &admin.Interactor{
		admin.Repository(db),
		presenter.Presenter{view.Viewer{w}},
	}
}

func Create(w http.ResponseWriter, r *http.Request) (err error) {
	db := mongo.DB.With(mongo.DB.Session.Copy())
	defer db.Session.Close()
	i := NewInteractor(&mongo.AdminDB{db}, w)
	ind, err := mapCreateRequest(r)
	if err != nil {
		return err
	}
	if err := i.Create(ind); err != nil {
		return err
	}

	return
}

func Get(w http.ResponseWriter, r *http.Request) (err error) {
	db := mongo.DB.With(mongo.DB.Session.Copy())
	defer db.Session.Close()
	i := NewInteractor(&mongo.AdminDB{db}, w)
	ind, err := mapGetRequest(r)
	if err != nil {
		return err
	}
	if err := i.Get(ind); err != nil {
		return err
	}

	return nil
}

func GetAll(w http.ResponseWriter, r *http.Request) (err error) {
	db := mongo.DB.With(mongo.DB.Session.Copy())
	defer db.Session.Close()
	i := NewInteractor(&mongo.AdminDB{db}, w)
	if err := i.GetAll(); err != nil {
		return err
	}

	return
}
