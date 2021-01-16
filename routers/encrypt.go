package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ballpor98/Go-NaCl-S/core"
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
	//read config
	secret := c.MustGet("AES_SECRET").(string)

	// validate and parse request
	var request EncryptRequestModel
	if err := c.ShouldBindJSON(&request); err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	  return
	}

	// handler
	ciphertext, err := core.Encrypt(request.Plaintext, secret, request.Nonce)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// response
	response := EncryptResponseModel{Ciphertext: ciphertext, Nonce: request.Nonce}
	c.JSON(http.StatusOK, response)
}
