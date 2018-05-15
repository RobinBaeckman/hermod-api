package main

import (
	"log"
	"net/http"
	"os"

	admincontroller "github.com/RobinBaeckman/hermod-api/application/admin/admin/controller"
	authcontroller "github.com/RobinBaeckman/hermod-api/application/admin/auth/controller"
	"github.com/RobinBaeckman/hermod-api/customerr"
	"github.com/RobinBaeckman/hermod-api/infra/middleware"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func main() {
	r := mux.NewRouter()
	logger := log.New(os.Stdout, "server: ", log.Lshortfile)
	authcontroller.Store = store
	// API
	//r.Handle("/products", phandler.Create).Methods("POST")
	//r.Handle("/products/{id}", phandler.Get).Methods("GET")
	//r.Handle("/products", phandler.GetAll).Methods("GET")
	r.Handle("/api/v1/admin/auth", Adapt(
		customerr.Check(authcontroller.Auth),
		middleware.Login(store),
		middleware.Notify(logger),
	)).Methods("POST")
	r.Handle("/api/v1/admins", Adapt(
		customerr.Check(admincontroller.Store),
		middleware.Auth(store),
		middleware.Notify(logger),
	)).Methods("POST")
	r.Handle("/api/v1/admins/{id}", Adapt(
		customerr.Check(admincontroller.Show),
		middleware.Auth(store),
		middleware.Notify(logger),
	)).Methods("GET")
	r.Handle("/api/v1/admins", Adapt(
		customerr.Check(admincontroller.Index),
		middleware.Auth(store),
	)).Methods("GET")
	r.HandleFunc("/api/v1/logout", authcontroller.Logout).Methods("GET")

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

func Adapt(h http.Handler, adapters ...middleware.Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}
