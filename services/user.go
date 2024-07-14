package services

import (
	db "e_real_estate/config"
	"e_real_estate/models"
	"e_real_estate/utils"
	"errors"
	"fmt"
)

type UserService struct {
	serverDb *db.Database
}

// creating new instance of UserService
func NewUserService(db *db.Database) *UserService { 
return &UserService{serverDb: db}
}

func (s UserService) CreateUser(userPayload models.UserPayload) (models.CreatedUserResponse, error) {
	//check db if user already exists

	row := s.serverDb.Db.QueryRow(`SELECT id, first_name, last_name, email, role FROM users WHERE email=$1`, userPayload.Email)
	existingUser := models.UserResponseData{}

	err := row.Scan(
		&existingUser.Id,
		&existingUser.FirstName, 
		&existingUser.LastName,
		&existingUser.Email,
		&existingUser.Role)


	if err == nil { 
		return models.CreatedUserResponse{}, errors.New("User with the given email alreay exists")
	}

	_, err = s.serverDb.Db.Exec(`Insert into users (first_name, last_name, email, password, role) values ($1, $2, $3, $4, $5)`, userPayload.FirstName, userPayload.LastName, userPayload.Email, userPayload.Password, userPayload.Role)
	
		if err != nil { 
			return models.CreatedUserResponse{}, errors.New("An error occured while creating user")
		}

	createdUser := models.UserResponseData{}
	affectedRow := s.serverDb.Db.QueryRow(`SELECT id, first_name, last_name, email, role FROM users WHERE email=$1`, userPayload.Email)
	
	err = affectedRow.Scan(&createdUser.Id, &createdUser.FirstName, &createdUser.LastName, &createdUser.Email, &createdUser.Role)

	if err != nil {
		return models.CreatedUserResponse{}, errors.New("Error occured while scanning user")
	}
	token, err := utils.CreateToken(createdUser.Email)
	if err != nil {
		return models.CreatedUserResponse{}, errors.New("Error occured")
	}

	response := models.CreatedUserResponse{
		Message: "User created successfully",
		Token: token,
		Data: createdUser,
		StatusCode: 200,
	}
	return response, nil
	

}
func (s UserService) Login(userPayload models.LoginPayload) (models.CreatedUserResponse, error){

	row := s.serverDb.Db.QueryRow(`SELECT id, first_name, last_name, email, password FROM users WHERE email = $1`, userPayload.Email)
	existingUser := models.UserPayload{}

	err := row.Scan(
		&existingUser.Id,
		&existingUser.FirstName, 
		&existingUser.LastName,
		&existingUser.Email,
		&existingUser.Password )

	if err != nil {
		return models.CreatedUserResponse{}, errors.New("User not found")
	}

	
	//compare password
	err = utils.ComparePassword(userPayload.Password, []byte(existingUser.Password))
	if err != nil { 
		return models.CreatedUserResponse{}, errors.New("Invalid credentials")
	}

	//generate token
	token, err := utils.CreateToken(existingUser.Email)
	if err != nil { 
		return models.CreatedUserResponse{}, errors.New("Error occured")}

		responseData := models.UserResponseData{
			Id :existingUser.Id,
			Email: existingUser.Email,
			FirstName : existingUser.FirstName,
			LastName :existingUser.LastName,
			Role :existingUser.Role,
		}
		
		//return user data to client
		response := models.CreatedUserResponse{
			Message: "Login successfully",
			Token: token,
			Data: responseData,
			StatusCode: 200,
		}
	return response, nil
}

func (s UserService) GetUser(email string) (models.CreatedUserResponse, error){

	row := s.serverDb.Db.QueryRow(`SELECT  id, first_name, last_name, email, role FROM users WHERE email = $1`, email)
	existingUser := models.UserResponseData{}

	err := row.Scan(
		&existingUser.Id,
		&existingUser.FirstName, 
		&existingUser.LastName,
		&existingUser.Email, 
		&existingUser.Role)

	if err != nil {
		return models.CreatedUserResponse{}, errors.New("User not found")
	}

		//return user data to client
		response := models.CreatedUserResponse{
			Message: "Login successfully",
			Data: existingUser,
			StatusCode: 200,
		}
	return response, nil
}

func (s UserService) GetAllUsers() (models.CreatedUserResponse, error) {
	rows, err:= s.serverDb.Db.Query("SELECT id, first_name, last_name, email, role FROM users")

	if err != nil { 
		return models.CreatedUserResponse{}, errors.New("Error occurred while querying users")
	}
	defer rows.Close()

	var  users []models.UserResponseData

	for rows.Next() {
	var user models.UserResponseData
	if err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.Role);
	err != nil {
		fmt.Println(err)
		return models.CreatedUserResponse{}, errors.New("Error occurred while scanning user")
	}
	users = append(users, user)
	}

	if len(users) == 0 {
		return models.CreatedUserResponse{}, errors.New("No users found!")
	}
	response := models.CreatedUserResponse{
		Message: "User fetched successfully",
		Data: users,
		StatusCode: 200,
	}

 return response, nil	
}