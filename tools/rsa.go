package tools

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

var PublicKey = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDMPRv0R8ONbIga1cDcXdrcVLo6
LpOB2Tod4/+ncTNvKpudI7MLtmLXPytWYoFl1s3/DmYZGyg55eHA0vZjYRajLjoj
7dKL7H/xL1a66AWb/FjjIs28+cRtt/mdx2g2UEdak46R2ycKVJM898heBqN738kW
ZZ9zcSX6gLmOCNLBFQIDAQAB
-----END PUBLIC KEY-----`

var PrivateKey = `-----BEGIN PRIVATE KEY-----
MIICdgIBADANBgkqhkiG9w0BAQEFAASCAmAwggJcAgEAAoGBAMw9G/RHw41siBrV
wNxd2txUujouk4HZOh3j/6dxM28qm50jswu2Ytc/K1ZigWXWzf8OZhkbKDnl4cDS
9mNhFqMuOiPt0ovsf/EvVrroBZv8WOMizbz5xG23+Z3HaDZQR1qTjpHbJwpUkzz3
yF4Go3vfyRZln3NxJfqAuY4I0sEVAgMBAAECgYAJfuHY1qlR3vTpAn0oAbkWO145
LEcxZ08baqlNOKciiQGZKbq+Val8xnQWXRgVCwqizCGVEz0oi/aWB3jrH+10bXyb
RRGr1JcwAAcExMDoIuG6MOprImGnE8UMXMhj1k3B+6e37WMavGcydntjPiAAw/jo
khSfX9mFV7MwN3k0AQJBAPCQX/oRx+xa4CtmBW1AYbHbz4ZW5uHOFoIPNxbO/sPL
lX6PdAIBJltjR7rLL+XfT7ikNH2+veNimOInF1n0LwECQQDZWAmLqSpPtRsp1JSQ
Y9x2MFzncLLoQitm+JbDBN/x15OPr8Frq82qjZBx2Sj1upqNrBqzxV2MV/3J4VzV
kOYVAkByQL+0rzk6ojaRphSxvMApjvJTJXbmi9DY2I0bghgxucE4qL06Ln2fLdnl
d5c6IANm+GYNyse49R0TW+mVSYoBAkAH0DZnouk2fFhBpLbCihR+2zY7y71ixB4z
UXR6Bk7Wrt1LKRJXAJIgM36h2SCz1MWBmlJLbCj0xqUFAOkJdHARAkEAh0Ekmdrc
W6viEMMNiab6OxbkcZ5X+Kq6Eg3c+ValEYlMcWn9OPb2v3OnRCcc7/LnChNycTg1
nlXrR47hWWUXSA==
-----END PRIVATE KEY-----`

//使用私钥加签，生成sign字符串
func Sha256WithRsa(content string, privateKey []byte) string {
	block, _ := pem.Decode(privateKey) //piravteKey为私钥文件的字节数组
	if block == nil {
		return ""
	}

	//priv即私钥对象,block.Bytes是私钥的字节流
	priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)

	if err != nil {
		return ""
	}

	h := crypto.Hash.New(crypto.SHA256)
	h.Write([]byte(content))
	hashed := h.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, priv.(*rsa.PrivateKey), crypto.SHA256, hashed) //签名

	return base64Encode(signature)
}

//公钥加密
func PublicEncrypt(data string, publicKey string) (string, error) {
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return "", errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	partLen := pub.N.BitLen()/8 - 11
	chunks := split([]byte(data), partLen)
	buffer := bytes.NewBufferString("")
	for _, chunk := range chunks {
		bytes, err := rsa.EncryptPKCS1v15(rand.Reader, pub, chunk)
		if err != nil {
			return "", err
		}
		buffer.Write(bytes)
	}

	return base64Encode(buffer.Bytes()), nil
}

// 私钥解密
func PrivateDecrypt(encrypted string, privateKey string) (string, error) {
	raw, err := base64.StdEncoding.DecodeString(encrypted)

	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return "", errors.New("private key error")
	}
	//解析PKCS8格式的私钥
	priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	privInterface := priv.(*rsa.PrivateKey)
	chunks := split([]byte(raw), privInterface.N.BitLen()/8)
	buffer := bytes.NewBufferString("")
	for _, chunk := range chunks {
		decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, privInterface, chunk)
		if err != nil {
			return "分段解密失败", err
		}
		buffer.Write(decrypted)
	}

	return buffer.String(), err
}

//验签
func verify(sign []byte, hashed []byte, publicKey string) bool {
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return false
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return false
	}
	pub := pubInterface.(*rsa.PublicKey) //pub:公钥对象
	err = rsa.VerifyPKCS1v15(pub, crypto.SHA256, hashed, []byte(base64Decode(string(sign))))

	if err != nil {
		return false
	} else {
		return true
	}
}

//base64加密
func base64Encode(src []byte) string {
	encodeString := base64.StdEncoding.EncodeToString(src)
	return encodeString
}

//base64解密
func base64Decode(src string) string {
	decodeBytes, _ := base64.StdEncoding.DecodeString(src)
	return string(decodeBytes)
}

//分段加解密私有方法
func split(buf []byte, lim int) [][]byte {
	var chunk []byte
	chunks := make([][]byte, 0, len(buf)/lim+1)
	for len(buf) >= lim {
		chunk, buf = buf[:lim], buf[lim:]
		chunks = append(chunks, chunk)
	}
	if len(buf) > 0 {
		chunks = append(chunks, buf[:len(buf)])
	}
	return chunks
}
