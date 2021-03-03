package interfaces

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

// CryptoHandler struct utilizada na implementacao dos metodos.
type CryptoHandler struct {
	SecretKey string
}

// NewCryptoHandler retorna uma instancia do cryptohandler.
func NewCryptoHandler() *CryptoHandler {
	return &CryptoHandler{
		SecretKey: os.Getenv("ENCRYPT_KEY"),
	}
}

// EncryptString retorna a string encriptada.
func (ch *CryptoHandler) EncryptString(text string) (string, error) {
	plaintext := []byte(text)

	block, err := aes.NewCipher([]byte(ch.SecretKey))
	if err != nil {
		log.Println(err)
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// DecryptString retorna a string decriptada.
func (ch *CryptoHandler) DecryptString(encryptedText string) (string, error) {
	ciphertext, err := base64.URLEncoding.DecodeString(encryptedText)
	if err != nil {
		log.Println(err)
		return "", err
	}

	block, err := aes.NewCipher([]byte(ch.SecretKey))
	if err != nil {
		log.Println(err)
		return "", err
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext), nil
}
