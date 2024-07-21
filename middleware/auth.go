package middleware

import (
	"fmt"
	"net/http"

	// "os"
	"strings"
	// "github.com/golang-jwt/jwt"
)



func Auth(next http.Handler) http.Handler {
	// secretKey := os.Getenv("TOKEN_SECRET_KEY")
	// var mySigningKey = []byte(secretKey)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//get the token from the header
		tokenString := strings.Replace(r.Header.Get("Authorization"), "Bearer", "", 1)

		// if tokenString == "" {
		// 	http.Error(w, "Missing or empty token", http.StatusUnauthorized)
		// 	return
		// }

		// //verify the token
		// token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 	if _, ok := token.Method.(*jwt.SigningMethodHMAC);
		// 	!ok {
		// 		return nil, fmt.Errorf("Unexpected signing method")
		// 	}
		// 	return mySigningKey, nil
		// })

		// if claims, ok := token.Claims.(jwt.MapClaims);
		// ok && token.Valid{
		// 	userID := claims["user"].(string)
		// 	r.Header.Set("User", userID)
		// }

		// if err != nil {
		// 	http.Error(w, "Invalid token or token expired", http.StatusUnauthorized)
		// }
		fmt.Println("token",tokenString)
		next.ServeHTTP(w, r)
	})

}