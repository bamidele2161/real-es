package services

import (
	db "e_real_estate/config"
	"e_real_estate/models"
	"errors"
)

type PropertyServiceInterface interface {
	CreateProperty(propertyPayload models.PropertyPayload) (models.CreatedPropertyResponse, error)
	GetProperty(id int) (models.CreatedPropertyResponse, error)
	GetAllProperties() (models.CreatedPropertyListResponse, error)
	DeleteProperty(id int) (models.CreatedPropertyResponse, error)
}
type PropertyService struct {
	serverDb *db.Database
}

func NewPropertyService(db *db.Database) *PropertyService {
	return &PropertyService{serverDb: db}
}

func (s PropertyService) CreateProperty(propertyPayload models.PropertyPayload) (models.CreatedPropertyResponse, error){
	//check for user
	row := s.serverDb.Db.QueryRow(`SELECT id, first_name, last_name, email, password FROM users WHERE id = $1`, propertyPayload.OwnedBy)
	existingUser := models.UserResponseData{}

	err := row.Scan(
		&existingUser.Id,
		&existingUser.FirstName, 
		&existingUser.LastName,
		&existingUser.Email,
		&existingUser.Role)

	if err != nil {
		return models.CreatedPropertyResponse{}, errors.New("User not found")
	}
	// save data

	result := s.serverDb.Db.QueryRow(`
	INSERT INTO properties (name, description, address, amount, owned_by) 
	VALUES ($1, $2, $3, $4, $5) RETURNING id`, 
	propertyPayload.Name, propertyPayload.Description, propertyPayload.Address, propertyPayload.Amount, propertyPayload.OwnedBy)

	var id int64
	err = result.Scan(&id)

		if err != nil { 
			return models.CreatedPropertyResponse{}, errors.New("An error occured while creating property")
		}
		createdProperty := models.PropertyResponseData{}
		affectedRow := s.serverDb.Db.QueryRow(`SELECT * FROM properties WHERE id = $1`, id)
		err = affectedRow.Scan(
			&createdProperty.Id, 
			&createdProperty.Name,
			&createdProperty.Description, 
			&createdProperty.Address, 
			&createdProperty.Amount, 
			&createdProperty.OwnedBy, 
			&createdProperty.CreatedAt, 
			&createdProperty.UpdatedAt)
		if err != nil {
			return models.CreatedPropertyResponse{}, errors.New("Error occured while scanning property")
		}
		response := models.CreatedPropertyResponse{
		Message: "Property created successfully",
		Data: createdProperty,
		StatusCode: 200,
	}
	return response, nil
}

func (s PropertyService) GetProperty(id int) (models.CreatedPropertyResponse, error) {
	row := s.serverDb.Db.QueryRow(`SELECT * FROM properties WHERE id = $1`, id)
	property := models.PropertyResponseData{}

	err := row.Scan(
		&property.Id, 
		&property.Name, 
		&property.Description, 
		&property.Address, 
		&property.Amount, 
		&property.OwnedBy, 
		&property.CreatedAt, 
		&property.UpdatedAt)

		if err != nil {
			return models.CreatedPropertyResponse{}, errors.New("Property not found")
		}

		response := models.CreatedPropertyResponse{
			Message: "Property created successfully",
			Data: property,
			StatusCode: 200,
		}
		return response, nil

}

func (s PropertyService) GetAllProperties() (models.CreatedPropertyListResponse, error) {

	rows, err := s.serverDb.Db.Query("SELECT * FROM properties")

	if err != nil {
		return models.CreatedPropertyListResponse{}, errors.New("Error occurred while querying properties")
	}
	defer rows.Close()

	var properties []models.PropertyResponseData

	for rows.Next() {
		var property models.PropertyResponseData
		if err := rows.Scan(
			&property.Id, 
			&property.Name, 
			&property.Description, 
			&property.Address, 
			&property.Amount, 
			&property.OwnedBy, 
			&property.CreatedAt, 
			&property.UpdatedAt); 
		err != nil {
			return models.CreatedPropertyListResponse{}, errors.New("Error occurred while scanning properties")
		}
		properties = append(properties, property)
	}
	if len(properties) == 0 {
		return models.CreatedPropertyListResponse{}, errors.New("No properties found!")
	}

	response:= models.CreatedPropertyListResponse{
		Message: "Properties fetched successfully",
		Data: properties,
		StatusCode: 200,
	}

	return response, nil

}

func (s PropertyService) DeleteProperty(id int) (models.CreatedPropertyResponse, error) {

	row, err := s.serverDb.Db.Exec("DELETE FROM properties WHERE id = $1", id)
	if err != nil {
		return models.CreatedPropertyResponse{}, errors.New("Error deleting property")
	}
	rowsAffected, err := row.RowsAffected()

		if err != nil {
			return models.CreatedPropertyResponse{}, errors.New("Error occured")
		}

		if rowsAffected == 0 {
			return models.CreatedPropertyResponse{}, errors.New("Property not found")
		}

		response := models.CreatedPropertyResponse{
			Message: "Property deleted successfully",
			StatusCode: 200,
		}
		return response, nil

}