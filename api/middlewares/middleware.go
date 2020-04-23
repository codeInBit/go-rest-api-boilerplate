package middlewares

import (
	"errors"
	"net/http"

	customAuth "github.com/codeinbit/go-shop/api/auth"
	"github.com/codeinbit/go-shop/api/utilities"
)

func SetMiddlewareJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	}
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := customAuth.TokenValid(r)
		if err != nil {
			utilities.ERROR(w, http.StatusUnauthorized, errors.New("unauthorized"))
			return
		}
		next(w, r)
	}
}
