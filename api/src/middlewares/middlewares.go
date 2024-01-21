package middlewares

import (
	"log"
	"net/http"
	"social-car/src/auth"
	"social-car/src/responses"
)

// Authenticate is a function that takes an http.HandlerFunc as a parameter and returns an http.HandlerFunc.
//
// The Authenticate function is used to authenticate the request before calling the next handler function.
// It prints "Authenticating..." to the console and then calls the next handler function with the given http.ResponseWriter and http.Request.
func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := auth.ValidateToken(r); err != nil {
			responses.Error(w, http.StatusUnauthorized, err)
			return
		}
		next(w, r)
	}
}

// Logger logs a request and calls the next http.HandlerFunc.
//
// Logger takes a http.HandlerFunc as a parameter and returns a new http.HandlerFunc.
// The returned function logs "Request received" to the console and then calls the
// original http.HandlerFunc with the given http.ResponseWriter and http.Request.
func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("-> %s %s %s\n", r.Method, r.RequestURI, r.Host)
		next(w, r)
	}
}
