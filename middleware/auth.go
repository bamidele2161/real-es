package middleware

import (
	"fmt"
	"net/http"
	"strings"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//get the token from the header
		tokenString := strings.Replace(r.Header.Get("Authorization"), "Bearer", "", 1)

		fmt.Println("token",tokenString)
		next.ServeHTTP(w, r)
	})

}