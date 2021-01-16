package core
 
import (
	"errors"
    "crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)
 
// ValidateKeyAndNonce validate length
func ValidateKeyAndNonce(key, nonce []byte) error{
	if l := len(key); l != 32 {
		return errors.New("key length invalid")
	}

	if l := len(nonce); l != 12 {
		return errors.New("nonce length invalid")
	}

	return nil
}

// Encrypt encrypt function for AES-256-GCM
func Encrypt(plaintext, key, nonce string) (string,error){
    bPlaintext := []byte(plaintext)

	bKey, err := base64.StdEncoding.DecodeString(key)
    if err != nil {
		return "" , err
	}
	bNonce, err := base64.StdEncoding.DecodeString(nonce)
    if err != nil {
		return "" , err
	}
	if err := ValidateKeyAndNonce(bKey, bNonce); err != nil {
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

// Decrypt decrypt function for AES-256-GCM
func Decrypt(ciphertext, key, nonce string) (string,error){
    bCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
    if err != nil {
		return "" , err
	}

	bKey, err := base64.StdEncoding.DecodeString(key)
    if err != nil {
		return "" , err
	}
	bNonce, err := base64.StdEncoding.DecodeString(nonce)
    if err != nil {
		return "" , err
	}
	if err := ValidateKeyAndNonce(bKey, bNonce); err != nil {
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
  
	bPlainText, err := aesgcm.Open(nil, bNonce, bCiphertext, nil)
	if err != nil {
		return "", err
	}
	plainText := string(bPlainText)
	
	return plainText , nil
}
