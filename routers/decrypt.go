package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// DecryptRequestModel encrypt api request model
type DecryptRequestModel struct {
	Ciphertext	string `json:"ciphertext" validate:"required"`
	Nonce		string `json:"nonce" validate:"required"`
}

// DecryptResponseModel encrypt api response model
type DecryptResponseModel struct {
	Plaintext	string `json:"plaintext"`
	Nonce		string `json:"nonce"`
}

// Decrypt router
func Decrypt(c *gin.Context) {
	var request DecryptRequestModel
	if err := c.ShouldBindJSON(&request); err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  return
	}

	response := DecryptResponseModel{Plaintext: "hello", Nonce: request.Nonce}

	c.JSON(http.StatusOK, response)
}
