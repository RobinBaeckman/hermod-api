package handler

import (
	"net/http"

	"github.com/RobinBaeckman/hermod-api/adapter/admin/product/presenter"
	"github.com/RobinBaeckman/hermod-api/domain/product"
	"github.com/RobinBaeckman/hermod-api/infra/mongo"
	"github.com/RobinBaeckman/hermod-api/infra/web/admin/product/view"
	usecase "github.com/RobinBaeckman/hermod-api/usecase/admin/product"
	"github.com/gorilla/mux"
)

type Req struct {
	*http.Request
}

func NewInteractor(r product.Repository, p product.Presenter) *usecase.Interactor {
	return &usecase.Interactor{
		r,
		p,
	}
}

func Create(w http.ResponseWriter, r *http.Request) {
	db := mongo.DB.With(mongo.DB.Session.Copy())
	defer db.Session.Close()
	i := NewInteractor(
		product.Repository(mongo.NewProductDB(db)),
		presenter.Presenter{view.Viewer{w}},
	)

	req := Req{r}
	rb := req.mapRequest()

	i.Create(rb)
}

func Get(w http.ResponseWriter, r *http.Request) {
	db := mongo.DB.With(mongo.DB.Session.Copy())
	defer db.Session.Close()
	i := NewInteractor(
		product.Repository(mongo.NewProductDB(db)),
		presenter.Presenter{view.Viewer{w}},
	)

	params := mux.Vars(r)
	id := params["id"]

	i.Get(id)
}

func GetAll(w http.ResponseWriter, r *http.Request) {
	db := mongo.DB.With(mongo.DB.Session.Copy())
	defer db.Session.Close()
	i := NewInteractor(
		product.Repository(mongo.NewProductDB(db)),
		presenter.Presenter{view.Viewer{w}},
	)

	i.GetAll()
}
