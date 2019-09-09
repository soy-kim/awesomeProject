package main

import (
	"crypto/aes"
	"fmt"
)

func main() {
	key := "Hello, Key 12345"
	s := "Hello, world! 12"

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		fmt.Println(err)
		return
	}

	ciphertext := make([]byte, len(s))
	block.Encrypt(ciphertext, []byte(s))
	fmt.Printf("%x\n", ciphertext)

	plantext := make([]byte, len(s))
	block.Decrypt(plantext, ciphertext)

	fmt.Println(string(plantext))

}
