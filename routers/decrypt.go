package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ballpor98/Go-NaCl-S/core"
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
	//read config
	secret := c.MustGet("AES_SECRET").(string)

	// validate and parse request
	var request DecryptRequestModel
	if err := c.ShouldBindJSON(&request); err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// handler
	plaintext, err := core.Decrypt(request.Ciphertext, secret, request.Nonce)
	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// response
	response := DecryptResponseModel{Plaintext: plaintext, Nonce: request.Nonce}
	c.JSON(http.StatusOK, response)
}
