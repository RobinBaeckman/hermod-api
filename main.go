package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	admincontroller "github.com/RobinBaeckman/hermod-api/application/admin/admin/controller"
	authcontroller "github.com/RobinBaeckman/hermod-api/application/admin/auth/controller"
	"github.com/RobinBaeckman/hermod-api/customerr"
	"github.com/RobinBaeckman/hermod-api/infra/middleware"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}

var store = sessions.NewCookieStore([]byte(viper.GetString("session.auth_key")))
var logger = log.New(os.Stdout, "hermod-api: ", log.Ltime)

func main() {

	r := mux.NewRouter()
	//logger := log.New(os.Stdout, viper.GetString("app.log_prefix"), log.Lshortfile)
	authcontroller.Store = store
	middleware.Logger = logger
	// API
	//r.Handle("/products", phandler.Create).Methods("POST")
	//r.Handle("/products/{id}", phandler.Get).Methods("GET")
	//r.Handle("/products", phandler.GetAll).Methods("GET")
	r.Handle("/api/v1/admin/auth", Adapt(
		customerr.Check(authcontroller.Auth),
		middleware.Login(store),
		middleware.Notify(),
	)).Methods("POST")
	r.Handle("/api/v1/admins", Adapt(
		customerr.Check(admincontroller.Store),
		middleware.Auth(store),
		middleware.Notify(),
	)).Methods("POST")
	r.Handle("/api/v1/admins/{id}", Adapt(
		customerr.Check(admincontroller.Show),
		middleware.Auth(store),
		middleware.Notify(),
	)).Methods("GET")
	r.Handle("/api/v1/admins", Adapt(
		customerr.Check(admincontroller.Index),
		middleware.Auth(store),
	)).Methods("GET")
	r.HandleFunc("/api/v1/logout", authcontroller.Logout).Methods("GET")

	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(viper.GetString("app.host")+":"+viper.GetString("app.port"), nil))
}

func Adapt(h http.Handler, adapters ...middleware.Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}
