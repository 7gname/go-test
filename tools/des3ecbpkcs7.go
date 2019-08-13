package tools

/*
	3des ecb pkcs7padding
*/

import (
	"bytes"
	"crypto/des"
	"encoding/base64"
	"errors"
)

func PKCS7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}

func UnPKCS7Padding(data []byte) []byte {
	length := len(data)
	unpadding := int(data[length-1])
	return data[:(length - unpadding)]
}

func DesECBEncrypt(s, k string) (string, error) {
	key, err := base64.RawStdEncoding.DecodeString(k)
	if err != nil {
		return "", err
	}
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return "", err
	}

	bs := block.BlockSize()
	data := PKCS7Padding([]byte(s), bs)
	if len(data)%bs != 0 {
		return "", errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	ss := base64.StdEncoding.EncodeToString(out)
	return ss, nil
}

func DesECBDecrypt(s, k string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	key, err := base64.RawStdEncoding.DecodeString(k)
	if err != nil {
		return "", err
	}
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	if len(data)%bs != 0 {
		return "", errors.New("crypto/cipher: input not full blocks")
	}
	out := make([]byte, len(data))
	dst := out
	for len(data) > 0 {
		block.Decrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}
	out = UnPKCS7Padding(out)
	return string(out), nil
}
