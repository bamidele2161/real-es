package models

type PropertyPayload struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Address string `json:"address"`
	Amount float64 `json:"amount"`
	OwnedBy int `json:"owned_by"`
}

type PropertyResponseData struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Address string `json:"address"`
	Amount float64 `json:"amount"`
	OwnedBy int `json:"owned_by"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type CreatedPropertyResponse struct {
	Message string `json:"message"`
	Data PropertyResponseData `json:"data"`
	StatusCode int `json:"statusCode"`
}

type CreatedPropertyListResponse struct {
	Message string `json:"message"`
	Data []PropertyResponseData `json:"data"`
	StatusCode int `json:"statusCode"`
}
