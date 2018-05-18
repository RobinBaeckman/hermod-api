package middleware

import (
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/spf13/viper"
)

type Adapter func(http.Handler) http.Handler

func Login(cStore *sessions.CookieStore) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, _ := cStore.Get(r, viper.GetString("session.cookie_name"))

			// Check if user is authenticated
			if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
				h.ServeHTTP(w, r)
				return
			}
		})
	}
}

// Compatible with http.HandlerFunc
func Auth(cStore *sessions.CookieStore) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, _ := cStore.Get(r, viper.GetString("session.cookie_name"))

			// Check if user is authenticated
			if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			h.ServeHTTP(w, r)
		})
	}
}

func Notify(logger *log.Logger) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
			h.ServeHTTP(w, r)
		})
	}
}
