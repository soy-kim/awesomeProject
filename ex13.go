package main

import (
	"crypto"
	"crypto/md5"
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func main() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err)
		return
	}
	publicKey := &privateKey.PublicKey

	message := "Hello, Golang"
	hash := md5.New()
	hash.Write([]byte(message))
	digest := hash.Sum(nil)

	var h1 crypto.Hash
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, h1, digest)

	var h2 crypto.Hash
	err = rsa.VerifyPKCS1v15(publicKey, h2, digest, signature)

	if err != nil {
		fmt.Println("fail")
	} else {
		fmt.Println("success")
	}

}
