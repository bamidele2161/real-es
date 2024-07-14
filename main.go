package main

import (
	"fmt"

	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"

	db "e_real_estate/config"
	"e_real_estate/routes"
	"e_real_estate/services"
)

// LoggingMiddleware logs the details of each incoming request
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
		log.Printf("Completed %s in %v", r.RequestURI, time.Since(start))
	})
}


func main() {

	godotenv.Load(".env")
	portString := os.Getenv("PORT")

	if portString == "" {
		log.Fatal("PORT is not found in the environment")
	}

	// Database connection
	database, err := db.NewDb()
	if err != nil {
		fmt.Println(err)
	}
	database.Connect()
	defer database.Db.Close()
	err = database.Db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("We are connected to postgres")

	// Initialize services and controllers
	userService := services.NewUserService(database)
	propertyService := services.NewPropertyService(database)
	router:= mux.NewRouter()

	// Combine routers
	userRouter := routes.UserRouter(userService)
	propertyRouter := routes.PropertyRouter(propertyService)

	router.PathPrefix("/users").Handler(userRouter)
	router.PathPrefix("/properties").Handler(propertyRouter)
	
	// Add logging middleware
	router.Use(LoggingMiddleware)

	handler := cors.Default().Handler(router)

	server := &http.Server{
		Handler: handler,
		Addr:    ":" + portString,
	}

	log.Printf("Server starting on port %v", portString)
	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

