package server

import "net/http"

func Middleware(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		handler.ServeHTTP(w, r)
	}
}
