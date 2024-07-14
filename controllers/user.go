package controllers

import (
	"e_real_estate/models"
	"e_real_estate/services"
	"e_real_estate/utils"
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
)

type UserController struct { 
	UserService *services.UserService
}
func Test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"StatusCode": "200", "Message": "Properties fetched successfully"}`))
}


func NewUserController (service *services.UserService) *UserController {
	return &UserController{UserService: service}
}


func (c UserController) CreateAccount(w http.ResponseWriter, r *http.Request) {
	// parse user body into userPayload
	w.Header().Set("Content-Type", "application/json")
	var userPayload models.UserPayload
	err := json.NewDecoder(r.Body).Decode(&userPayload) 

	if err != nil { 
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	//validate the payload
		emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

		if !emailRegex.MatchString(userPayload.Email) {
			utils.RespondWithError(w, http.StatusBadRequest, "Please provide a valid email address")
			return
		} else if !utils.Validator(w, userPayload.Password, "Password", 6) || 
		!utils.Validator(w, userPayload.FirstName, "First Name", 3) || 
		!utils.Validator(w, userPayload.LastName, "Last Name", 3) {
			return
		}

		
		// hash password
		hashedPassword, _ := utils.HashPassword(userPayload.Password, 6)

		userPayload.Password = string(hashedPassword)

		createdUser, err := c.UserService.CreateUser(userPayload)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, err.Error())
			return
		}
		
		json.NewEncoder(w).Encode(createdUser)

}

func (c UserController) Login(w http.ResponseWriter, r *http.Request) {
	//parse user body into userPayload
	w.Header().Set("Content-Type", "application/json")
	var loginPayload models.LoginPayload
	err := json.NewDecoder(r.Body).Decode(&loginPayload)

	if err != nil { 
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	//validate the payload
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	if !emailRegex.MatchString(loginPayload.Email) {
		utils.RespondWithError(w, http.StatusBadRequest, "Please provide a valid email address")
		return
	} else if !utils.Validator(w, loginPayload.Password, "Password", 6){
		return
	}


	LoginUser, err := c.UserService.Login(loginPayload)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	
	json.NewEncoder(w).Encode(LoginUser)

}

func (c UserController) GetUserProfile(w http.ResponseWriter, r *http.Request) {
	//get the email parameter from the query string

	// email := r.URL.Query().Get("email")
	email := mux.Vars(r)["email"]

	//validate the payload
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	if !emailRegex.MatchString(email) {
		utils.RespondWithError(w, http.StatusBadRequest, "Please provide a valid email address")
		return
	}

	getUser, err := c.UserService.GetUser(email)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	
	json.NewEncoder(w).Encode(getUser)

}

func (c UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	getAllUsers, err := c.UserService.GetAllUsers()
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	json.NewEncoder(w).Encode(getAllUsers)
}