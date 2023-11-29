package cripto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

func Encrypt(c string) string {
	key := []byte("monosialotetraesosilgangliosideo")
	plaintext := []byte(c)
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	nonce := []byte("extrovertido")
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	return hex.EncodeToString(ciphertext)
}

func Decrypt(c string) string {
	key := []byte("monosialotetraesosilgangliosideo")
	ciphertext, _ := hex.DecodeString(c)
	nonce := []byte("extrovertido")
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}
	return string(plaintext)
}