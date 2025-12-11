package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {

	data := []byte("Hello")
	hash := sha256.Sum256(data)
	//fmt.Println(data)
	fmt.Println("Hash after sha256", hash)

}
