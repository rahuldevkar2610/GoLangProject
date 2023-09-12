package middlewares

import (
	"HotelAggregatorService/auth"
	"context"
	"fmt"
	"net/http"
)

const AUTHINFO_KEY = "authinfo"

func AuthMiddleware(auth auth.Auth, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, pass, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Not Authorized")
			return
		}
		authinfo, err := auth.Authenticate(user, pass)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "Not Authorized")
			return
		}

		req := r.WithContext(context.WithValue(r.Context(), AUTHINFO_KEY, authinfo))

		next(w, req)
	}
}

func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		authinfo, ok := r.Context().Value(AUTHINFO_KEY).(*auth.AuthInfo)
		if ok {
			fmt.Printf("user %v with role %v making request to %v\n", authinfo.User, authinfo.Role, r.URL)
		} else {
			//not *AuthInfo or may be AuthInfo not present in context
			fmt.Println("Request to ", r.URL)
		}

		next(w, r)
	}
}
