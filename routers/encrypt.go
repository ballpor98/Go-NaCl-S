package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// EncryptRequestModel encrypt api request model
type EncryptRequestModel struct {
	Plaintext	string `json:"plaintext" validate:"required"`
	Nonce		string `json:"nonce" validate:"required"`
}

// EncryptResponseModel encrypt api response model
type EncryptResponseModel struct {
	Ciphertext	string `json:"ciphertext"`
	Nonce		string `json:"nonce"`
}

// Encrypt router
func Encrypt(c *gin.Context) {
	var request EncryptRequestModel
	if err := c.ShouldBindJSON(&request); err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  return
	}

	response := EncryptResponseModel{Ciphertext: "QWx9S3BdA3980PKyNlF4YubC2Ko5", Nonce: request.Nonce}

	c.JSON(http.StatusOK, response)
}
