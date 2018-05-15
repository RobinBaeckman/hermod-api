package controller

import (
	"net/http"

	"github.com/RobinBaeckman/hermod-api/application/admin/admin/presenter"
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

func Store(w http.ResponseWriter, r *http.Request) (err error) {
	db := mongo.DB.With(mongo.DB.Session.Copy())
	defer db.Session.Close()
	i := NewInteractor(&mongo.AdminDB{db}, w)
	ind, err := mapCreateRequest(r)
	if err != nil {
		return err
	}
	if err := i.Store(ind); err != nil {
		return err
	}

	return
}

func Show(w http.ResponseWriter, r *http.Request) (err error) {
	db := mongo.DB.With(mongo.DB.Session.Copy())
	defer db.Session.Close()
	i := NewInteractor(&mongo.AdminDB{db}, w)
	ind, err := mapGetRequest(r)
	if err != nil {
		return err
	}
	if err := i.Show(ind); err != nil {
		return err
	}

	return nil
}

func Index(w http.ResponseWriter, r *http.Request) (err error) {
	db := mongo.DB.With(mongo.DB.Session.Copy())
	defer db.Session.Close()
	i := NewInteractor(&mongo.AdminDB{db}, w)
	if err := i.Index(); err != nil {
		return err
	}

	return
}
