package usecases

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"os"
)

var secretKey = os.Getenv("ENCRYPT_KEY")

// EncryptString retorna a string encriptada.
func EncryptString(text string) string {
	plaintext := []byte(text)

	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		panic(err)
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return base64.URLEncoding.EncodeToString(ciphertext)
}

// DecryptString retorna a string decriptada.
func DecryptString(encryptedText string) string {
	ciphertext, _ := base64.URLEncoding.DecodeString(encryptedText)

	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		panic(err)
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}
