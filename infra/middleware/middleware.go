package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/spf13/viper"
)

type Adapter func(http.Handler) http.Handler

var Logger *log.Logger

func Login(cs *sessions.CookieStore) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, _ := cs.Get(r, viper.GetString("session.cookie_name"))

			// Check if user is authenticated
			if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
				h.ServeHTTP(w, r)
				return
			}

			fmt.Println("You are authenticated")
		})
	}
}

// Compatible with http.HandlerFunc
func Auth(store *sessions.CookieStore) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, _ := store.Get(r, viper.GetString("session.cookie_name"))

			// Check if user is authenticated
			if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			fmt.Println("You are authenticated")

			// Now you can write back your template or re-direct elsewhere
			h.ServeHTTP(w, r)
		})
	}
}

func Notify() Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			Logger.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
			h.ServeHTTP(w, r)
		})
	}
}
