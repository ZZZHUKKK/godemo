package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"os"
)

type Encrypt struct {
	Key []byte
}

func NewEncrypt() (*Encrypt, error) {
	hexKey := os.Getenv("KEY")
	if hexKey == "" {
		return nil, errors.New("переменная KEY не задана")
	}
	key, err := hex.DecodeString(hexKey)
	if err != nil {
		return nil, errors.New("ключ должен быть в hex-формате")
	}
	return &Encrypt{
		Key: key,
	}, nil
}

func (enc *Encrypt) Encrypt(str []byte) []byte {
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		panic(err.Error())
	}
	aesGSM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, aesGSM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		panic(err.Error())
	}
	return aesGSM.Seal(nonce, nonce, str, nil)
}

func (enc *Encrypt) Decrypt(str []byte) []byte {
	block, err := aes.NewCipher([]byte(enc.Key))
	if err != nil {
		panic(err.Error())
	}
	aesGSM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := aesGSM.NonceSize()
	nonce, clipherText := str[:nonceSize], str[nonceSize:]
	plainText, err := aesGSM.Open(nil, nonce, clipherText, nil)
	if err != nil {
		panic(err.Error())
	}
	return plainText
}
