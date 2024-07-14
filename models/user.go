package models

type Role string

const (
	UserRole Role = "user"
	OwnerRole Role = "owner"
)
type UserPayload struct {
	Id int `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role Role `json:"role"`
}

type UserResponseData struct {
	Id int `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Role Role `json:"role"`
}
type CreatedUserResponse struct {
	Message string `json:"message"`
	Token string `json:"token"`
	Data interface{} `json:"data"`
	StatusCode int `json:"statusCode"`
}

type LoginPayload struct {
	Email string
	Password string
}
type LoginResponse struct {
	email string
	password string
}

