package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

type Adapter func(http.Handler) http.Handler

func Login(cs *sessions.CookieStore) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			session, _ := cs.Get(r, "cookie-name")

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
			session, _ := store.Get(r, "cookie-name")

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

func Notify(logger *log.Logger) Adapter {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
			h.ServeHTTP(w, r)
		})
	}
}
