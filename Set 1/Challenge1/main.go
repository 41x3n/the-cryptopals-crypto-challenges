package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	input := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	output := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	fmt.Println("Input: ", input)
	fmt.Println("Output: ", output)

	byteStr, err := hex.DecodeString(input)
	if err != nil {
		panic(err)
	}

	base64Str := base64.StdEncoding.EncodeToString(byteStr)
	if output == base64Str {
		fmt.Println("The strings are equal")
	} else {
		fmt.Println("The strings are not equal")
	}
}
