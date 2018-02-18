package main

import (
	"log"
	"net/http"
	"os"

	"github.com/RobinBaeckman/hermod-api/customerr"
	"github.com/RobinBaeckman/hermod-api/infra/middleware"
	adminhandler "github.com/RobinBaeckman/hermod-api/infra/web/admin/admin/handler"
	authhandler "github.com/RobinBaeckman/hermod-api/infra/web/admin/auth/handler"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func main() {
	r := mux.NewRouter()
	logger := log.New(os.Stdout, "server: ", log.Lshortfile)
	authhandler.Store = store
	// API
	//r.Handle("/products", phandler.Create).Methods("POST")
	//r.Handle("/products/{id}", phandler.Get).Methods("GET")
	//r.Handle("/products", phandler.GetAll).Methods("GET")
	r.Handle("/api/admin/auth", Adapt(
		customerr.Check(authhandler.Auth),
		middleware.Login(store),
		middleware.Notify(logger),
	)).Methods("POST")
	r.Handle("/api/admins", Adapt(
		customerr.Check(adminhandler.Create),
		middleware.Auth(store),
		middleware.Notify(logger),
	)).Methods("POST")
	r.Handle("/api/admins/{id}", Adapt(
		customerr.Check(adminhandler.Get),
		middleware.Auth(store),
		middleware.Notify(logger),
	)).Methods("GET")
	r.Handle("/api/admins", Adapt(customerr.Check(adminhandler.GetAll), middleware.Auth(store))).Methods("GET")
	r.HandleFunc("/api/logout", authhandler.Logout).Methods("GET")

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}

func Adapt(h http.Handler, adapters ...middleware.Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}
