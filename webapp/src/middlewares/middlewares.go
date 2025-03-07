package middlewares

import (
	"log"
	"net/http"
	"webapp/src/cookies"
)

// writes request info on terminal
func Logger(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("\n %s %s %s", r.Method, r.RequestURI, r.Host)
		nextFunction(w, r)
	}
}

// verifies if cookies exist
func Authenticate(nextFunction http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if _, error := cookies.Read(r); error != nil {
			http.Redirect(w, r, "/login", 302)
			return
		}
		nextFunction(w, r)

	}
}
