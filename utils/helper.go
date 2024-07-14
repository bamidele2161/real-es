package utils

import (
	"e_real_estate/models"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)


func HashPassword (password string, cost int) ([]byte, error) {
	hashedPassword, _ :=bcrypt.GenerateFromPassword([]byte(password), cost)
	return hashedPassword, nil
}

func ComparePassword (password string, hashedPassword []byte) ( error) {
	err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))

	if err != nil{
		return err
	}
	return nil
}
func RespondWithError (w http.ResponseWriter, statusCode int, errMessage string)  {
	response := models.ErrorResponse{
		Error: errMessage,
		StatusCode: statusCode,
	}

	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response)
}

func Validator (w http.ResponseWriter, fieldValue, fieldName string, minLength int) bool {
	if len(fieldValue) < minLength {
		RespondWithError(w, 400, fieldName + " must be at least " + strconv.Itoa(minLength) + " characters")
		return false
	}
	return true
}


func CreateToken (id string) (string, error) {
	godotenv.Load(".env")
	secretKey := os.Getenv("TOKEN_SECRET_KEY")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
	jwt.MapClaims{
		"id": id,
		"exp": time.Now().Add(time.Hour *24).Unix(),
	})

	tokenKey := []byte(secretKey)
	tokenString, err := token.SignedString(tokenKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}