package utils

import (
	"e_real_estate/models"
	"encoding/json"
	"net/http"

	"golang.org/x/time/rate"
)

func RateLimiter(next http.Handler) http.Handler{
	limiter := rate.NewLimiter(2,4)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow(){
			message:= models.ErrorResponse{
				Error: "Request failed, The API is at capacity, try again later",
				StatusCode: 429,
			}
			w.WriteHeader(http.StatusTooManyRequests)
			json.NewEncoder(w).Encode(message)
			return
		} 
		next.ServeHTTP(w, r)
	})
}