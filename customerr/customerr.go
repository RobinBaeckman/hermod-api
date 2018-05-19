package customerr

import (
	"log"
	"net/http"
	"os"

	"github.com/spf13/viper"
)

type Check func(http.ResponseWriter, *http.Request) error

func (e *App) Error() string { return e.Message }

type App struct {
	Err     error
	Message string
	Code    int
}

// error handler
func (fn Check) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	logger := log.New(os.Stdout, viper.GetString("app.log_prefix"), 3)
	if e := fn(w, r); e != nil {
		switch v := e.(type) {
		case *App:
			http.Error(w, v.Message, v.Code)
			if v.Err != nil {
				logger.Printf("Status: %v, Message: %v, Error: %v", v.Code, v.Message, v.Err)
			} else {
				logger.Printf("Status: %v, Message: %v", v.Code, v.Message)
			}
		default:
			http.Error(w, v.Error(), 500)
			logger.Printf("Status: %v, Error: %v", 500, v.Error())
		}
	}
}
