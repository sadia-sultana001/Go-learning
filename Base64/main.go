package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	var s string
	s = "hello world"
	byteArr := []byte(s)
	enc := base64.URLEncoding
	enc = enc.WithPadding(base64.NoPadding)
	b64str := enc.Strict().EncodeToString(byteArr)

	fmt.Println(s)
	fmt.Println(byteArr)
	fmt.Println(b64str)

	decodedStr, err := enc.DecodeString(b64str)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(decodedStr)
}
