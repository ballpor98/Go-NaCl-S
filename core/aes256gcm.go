package core
 
import (
    "crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)
 
// Encrypt encrypt function for AES-256-GCM
func Encrypt(plaintext string, key string, nonce string) (string,error){
    bPlaintext := []byte(plaintext)

	bKey, err := base64.StdEncoding.DecodeString(key)
    if err != nil {
		return "" , err
	}

	bNonce, err := base64.StdEncoding.DecodeString(nonce)
    if err != nil {
		return "" , err
	}

    block, err := aes.NewCipher(bKey)
    if err != nil {
      return "" , err
    }
  
    aesgcm, err := cipher.NewGCM(block)
    if err != nil {
		return "" , err
	}
  
	bCiphertext := aesgcm.Seal(nil, bNonce, bPlaintext, nil)
	ciphertext := base64.StdEncoding.EncodeToString(bCiphertext)
	
	return ciphertext , nil
}
