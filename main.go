package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

func main() {
}

func EncryptAesToByte(plainText, encryptKey string) ([]byte, error) {
	cipherAdapter, err := aes.NewCipher([]byte(encryptKey))
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(cipherAdapter)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	result := gcm.Seal(nonce, nonce, []byte(plainText), nil)
	return result, nil
}

func DecryptAesFromByte(cipherText []byte, encryptKey string) (string, error) {
	cipherAdapter, err := aes.NewCipher([]byte(encryptKey))
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(cipherAdapter)
	if err != nil {
		return "", err
	}

	mainSize := gcm.NonceSize()
	if len(cipherText) < mainSize {
		return "", errors.New("too short")
	}

	nonce, result := cipherText[:mainSize], cipherText[mainSize:]
	mainResult, err := gcm.Open(nil, nonce, result, nil)
	return string(mainResult), err
}
