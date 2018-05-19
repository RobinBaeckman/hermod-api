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
	mgo "gopkg.in/mgo.v2"
)

func main() {
	parseConfig()
	db := newMongoDB()
	cStore := sessions.NewCookieStore([]byte(viper.GetString("session.auth_key")))
	logger := log.New(os.Stdout, viper.GetString("app.log_prefix"), 3)
	authApp := &authcontroller.App{CStore: cStore, DB: db}
	adminApp := &admincontroller.App{CStore: cStore, DB: db}
	r := mux.NewRouter()
	// API
	//r.Handle("/products", phandler.Create).Methods("POST")
	//r.Handle("/products/{id}", phandler.Get).Methods("GET")
	//r.Handle("/products", phandler.GetAll).Methods("GET")
	r.Handle("/api/v1/admin/auth", Adapt(
		customerr.Check(authApp.Auth),
		middleware.Login(cStore),
		middleware.Notify(logger),
	)).Methods("POST")
	r.Handle("/api/v1/admins", Adapt(
		customerr.Check(adminApp.Store),
		middleware.Auth(cStore),
		middleware.Notify(logger),
	)).Methods("POST")
	r.Handle("/api/v1/admins/{id}", Adapt(
		customerr.Check(adminApp.Show),
		middleware.Auth(cStore),
		middleware.Notify(logger),
	)).Methods("GET")
	r.Handle("/api/v1/admins", Adapt(
		customerr.Check(adminApp.Index),
		middleware.Auth(cStore),
		middleware.Notify(logger),
	)).Methods("GET")
	r.HandleFunc("/api/v1/logout", authApp.Logout).Methods("GET")

	http.Handle("/", r)

	logger.Println("Started")
	log.Fatal(http.ListenAndServe(viper.GetString("app.host")+":"+viper.GetString("app.port"), nil))
}

func Adapt(h http.Handler, adapters ...middleware.Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}

func parseConfig() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath(".")      // optionally look for config in the working directory
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		fmt.Errorf("Fatal error config file: %s \n", err)
	}
}

func newMongoDB() *mgo.Database {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("hermod")

	return c
}
