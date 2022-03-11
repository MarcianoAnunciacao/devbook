package middlewares

import (
	"log"
	"net/http"
	"webapp/src/cookies"
)

func Logger(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunction(w, r)
	}
}

func IsThereACookieOnRequest(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := cookies.ReadCookie(r); err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		nextFunction(w, r)
	}
}
