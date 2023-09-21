package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"io"
)

func pkcs_pad(ciphertext []byte, blockSize int) []byte {

	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	return append(ciphertext, padtext...)
}

func pkcs_trim(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

func getKeyAndIV(Password string, salt string) (string, string) {
	salted := ""
	dI := ""

	for len(salted) < 48 {
		md := md5.New()
		md.Write([]byte(dI + Password + salt))

		dM := md.Sum(nil)
		dI = string(dM[:16])

		salted = salted + dI
	}

	key := salted[0:32]
	iv := salted[32:48]

	return key, iv
}

func AesDecrypt(payload string, key string) string {

	ciphertext, _ := base64.StdEncoding.DecodeString(payload)

	if len(ciphertext) < 16 || string(ciphertext[:8]) != "Salted__" {
		return ""
	}

	salt := ciphertext[8:16]
	ct := ciphertext[16:]

	key, iv := getKeyAndIV(key, string(salt))

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	cbc := cipher.NewCBCDecrypter(block, []byte(iv))
	dst := make([]byte, len(ct))
	cbc.CryptBlocks(dst, ct)

	return string(pkcs_trim(dst))
}

func EncryptAes(text string, key string) string {
	salt := make([]byte, 8)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		panic(err)
	}

	key, iv := getKeyAndIV(key, string(salt))

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}

	pad := pkcs_pad([]byte(text), block.BlockSize())
	ecb := cipher.NewCBCEncrypter(block, []byte(iv))
	encrypted := make([]byte, len(pad))
	ecb.CryptBlocks(encrypted, pad)

	return base64.StdEncoding.EncodeToString([]byte("Salted__" + string(salt) + string(encrypted)))
}
