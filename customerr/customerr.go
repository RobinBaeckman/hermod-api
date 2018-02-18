package customerr

import (
	"log"
	"net/http"
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
	if e := fn(w, r); e != nil {
		switch v := e.(type) {
		case *App:
			http.Error(w, v.Message, v.Code)
			log.Println(*v)
		default:
			http.Error(w, v.Error(), 500)
			log.Println(v)
		}
	}
}
