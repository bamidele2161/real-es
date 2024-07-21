package routes

import (
	"e_real_estate/controllers"
	"e_real_estate/middleware"
	"e_real_estate/services"
	"e_real_estate/utils"
	"net/http"

	"github.com/gorilla/mux"
)

func ChainMiddlewares(handler http.Handler, middlwares ...func(http.Handler) http.Handler) http.Handler {
	for _, middlware := range middlwares {
		handler = middlware(handler)
	}
	return handler
}


func PropertyRouter(propertyService *services.PropertyService) *mux.Router {
	propertyController := controllers.NewPropertyController(propertyService)

	router := mux.NewRouter()


	//PROPERTY routes
	router.HandleFunc("/properties/create", propertyController.CreateProperty).Methods("POST")
	router.Handle("/properties", ChainMiddlewares(http.HandlerFunc(propertyController.GetAllProperties), middleware.Auth, utils.RateLimiter)).Methods("GET")
	router.HandleFunc("/properties/{id}", propertyController.GetProperty).Methods("GET")
	router.HandleFunc("/properties/{id}", propertyController.DeleteProperty).Methods("DELETE")

	return router
}