package routes

import (
	"e_real_estate/controllers"
	"e_real_estate/middleware"
	"e_real_estate/services"
	"net/http"

	"github.com/gorilla/mux"
)


func UserRouter(userService *services.UserService) *mux.Router {
	userController := controllers.NewUserController(userService)

	router := mux.NewRouter()


	// AUTH routes
	router.HandleFunc("/users/test/test", controllers.Test).Methods("GET")
	router.HandleFunc("/users/register", userController.CreateAccount).Methods("POST")
	router.HandleFunc("/users/login", userController.Login).Methods("POST")
	router.Handle("/users/{email}",middleware.Auth(http.HandlerFunc(userController.GetUserProfile))).Methods("GET")
	router.HandleFunc("/users", userController.GetAllUsers).Methods("GET")
	
	return router
}
