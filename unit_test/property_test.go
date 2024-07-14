package tests

import (
	"bytes"
	"e_real_estate/controllers"
	"e_real_estate/models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)




type MockPropertyService struct{}
func (m *MockPropertyService) CreateProperty(propertyPayload models.PropertyPayload) (models.CreatedPropertyResponse, error) {
	return models.CreatedPropertyResponse{
		Message:    "Property created successfully",
		Data:       models.PropertyResponseData{Id: 1, Name: propertyPayload.Name},
		StatusCode: 200,
	}, nil
} 

func (m *MockPropertyService) GetAllProperties() (models.CreatedPropertyListResponse, error) {
	return models.CreatedPropertyListResponse{
		Message:    "Properties fetched successfully",
		Data: []models.PropertyResponseData{
			{Id: 1, Name: "Testing Apartment", Description: "", Address: "123 Street", Amount: 100000, OwnedBy: 1, CreatedAt: "", UpdatedAt: ""},
			{Id: 2, Name: "Testing Apartment", Description: "", Address: "123 Street", Amount: 100000, OwnedBy: 1, CreatedAt: "", UpdatedAt: ""}},
		StatusCode: 200,
	}, nil
}



func (m *MockPropertyService) GetProperty(id int) (models.CreatedPropertyResponse, error) {
	return models.CreatedPropertyResponse{
		Message:    "Property fetched successfully",
		Data: models.PropertyResponseData{Id: 1, Name: "Testing Apartment", Description: "", Address: "123 Street", Amount: 100000, OwnedBy: 1, CreatedAt: "", UpdatedAt: ""},
		StatusCode: 200,
	}, nil
}

func (m *MockPropertyService) DeleteProperty(id int) (models.CreatedPropertyResponse, error) {
	return models.CreatedPropertyResponse{
		Message:    "Property deleted successfully",
		StatusCode: 200,
	}, nil
}

func TestCreateProperty(t *testing.T) {
	mockService := &MockPropertyService{}
	controller := controllers.NewPropertyController((mockService))

	propertyPayload := models.PropertyPayload{
		Name: "Create Test", 
		Description: "This is a test for create endpoint",
		Address: "Lagos, Nigeria",
		Amount: 3000,
		OwnedBy: 2,
	}
	body, err := json.Marshal(propertyPayload)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest("POST", "/properties/create", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("this is the main error:", err)
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/properties/create", controller.CreateProperty).Methods("POST")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var response models.CreatedPropertyResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	assert.Equal(t, "Property created successfully", response.Message)
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "Create Test", response.Data.Name)
}

func TestGetProperty(t *testing.T) {
	mockService := &MockPropertyService{}
	controller := controllers.NewPropertyController(mockService)

	req, err := http.NewRequest("GET", "/properties/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/properties/{id}", controller.GetProperty).Methods("GET")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var response models.CreatedPropertyResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	property:= response.Data
	assert.Equal(t, "123 Street", property.Address)
}


func TestGetAllProperties(t *testing.T) {
	mockService := &MockPropertyService{}
	controller := controllers.NewPropertyController(mockService)

	req, err := http.NewRequest("GET", "/properties", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/properties", controller.GetAllProperties).Methods("GET")
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var response models.CreatedPropertyListResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)
	properties := response.Data
	assert.Len(t, properties, 2)
}


func TestDeleteProperty(t *testing.T) {
	mockService := &MockPropertyService{}
	controller := controllers.NewPropertyController(mockService)

	req, err := http.NewRequest("DELETE", "/properties/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/properties/{id}", controller.DeleteProperty).Methods("DELETE")
	router.ServeHTTP(rr, req)

	// assert.Equal(t, http.StatusNoContent, rr.Code)
}


