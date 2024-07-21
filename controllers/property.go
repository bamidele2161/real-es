package controllers

import (
	"e_real_estate/models"
	"e_real_estate/services"
	"e_real_estate/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

type PropertyController struct {
	PropertyService services.PropertyServiceInterface
}

func NewPropertyController(service services.PropertyServiceInterface) *PropertyController {
	return &PropertyController{PropertyService: service}
}

func (c PropertyController) CreateProperty(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var propertyPayload []models.PropertyPayload
	err := json.NewDecoder(r.Body).Decode(&propertyPayload)

	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	
	var wg sync.WaitGroup
	responses := make(chan models.CreatedPropertyResponse, len(propertyPayload))

	for _, payload := range propertyPayload {
		wg.Add(1)
		go func(payload models.PropertyPayload) {
			defer wg.Done()

			response, err := c.PropertyService.CreateProperty(payload)

			if err != nil {
				utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
				return
			} else {
				responses <- response
			}

		}(payload)

	}
	

	wg.Wait()
	close(responses)

	var result []models.CreatedPropertyResponse
	for response := range responses {
		result = append(result, response)
	}
	json.NewEncoder(w).Encode(result)
}


func (c PropertyController) GetProperty(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	propertyId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Error getting the property ID", http.StatusBadRequest)
		return
	}
	getProperty, err := c.PropertyService.GetProperty(propertyId)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(getProperty)
}

func (c PropertyController) GetAllProperties(w http.ResponseWriter, r *http.Request) {
	getAllProperties, err := c.PropertyService.GetAllProperties()

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(getAllProperties)
}

func (c PropertyController) DeleteProperty(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	propertyId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Error getting the property ID", http.StatusBadRequest)
		return
	}
	getProperty, err := c.PropertyService.DeleteProperty(propertyId)

	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(getProperty)
}