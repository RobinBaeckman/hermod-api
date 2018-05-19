package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	admincontroller "github.com/RobinBaeckman/hermod-api/application/admin/admin/controller"
	authcontroller "github.com/RobinBaeckman/hermod-api/application/admin/auth/controller"
	productcontroller "github.com/RobinBaeckman/hermod-api/application/admin/product/controller"
	"github.com/RobinBaeckman/hermod-api/customerr"
	"github.com/RobinBaeckman/hermod-api/infra/middleware"
	"github.com/RobinBaeckman/hermod-api/infra/mongo"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/spf13/viper"
)

func main() {
	parseConfig()
	db := mongo.NewDB()
	cStore := sessions.NewCookieStore([]byte(viper.GetString("session.auth_key")))
	logger := log.New(os.Stdout, viper.GetString("app.log_prefix"), 3)
	authApp := &authcontroller.App{CStore: cStore, DB: db}
	adminApp := &admincontroller.App{CStore: cStore, DB: db}
	productApp := &productcontroller.App{CStore: cStore, DB: db}
	r := mux.NewRouter()

	// Product
	r.Handle("/api/v1/products", Adapt(
		customerr.Check(productApp.Store),
		middleware.Auth(cStore),
		middleware.Notify(logger),
	)).Methods("POST")
	r.Handle("/api/v1/products/{id}", Adapt(
		customerr.Check(productApp.Show),
		middleware.Auth(cStore),
		middleware.Notify(logger),
	)).Methods("GET")
	r.Handle("/api/v1/products", Adapt(
		customerr.Check(productApp.Index),
		middleware.Auth(cStore),
		middleware.Notify(logger),
	)).Methods("GET")

	// Admin user
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

	// Auth
	r.Handle("/api/v1/admin/auth", Adapt(
		customerr.Check(authApp.Auth),
		middleware.Login(cStore),
		middleware.Notify(logger),
	)).Methods("POST")
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
