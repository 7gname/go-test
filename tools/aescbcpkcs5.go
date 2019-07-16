package tools

import (
	"crypto/aes"
	"crypto/cipher"
	"bytes"
	"encoding/hex"
)

func AesEncrypt(content, k, iv string) (rs string, err error) {
	origData := []byte(content)
	key, err := hex.DecodeString(k)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return hex.EncodeToString(crypted), nil
}

func AesDecrypt(content, k, iv string) (rs string, err error) {
	crypted, err := hex.DecodeString(content)
	if err != nil {
		return "", err
	}
	key, err := hex.DecodeString(k)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return string(origData), nil
}

func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
