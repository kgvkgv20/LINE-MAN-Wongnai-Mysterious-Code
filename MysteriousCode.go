package main

import (
	"encoding/base64"
	"fmt"
)

func reversedata(input []byte) []byte {
	if len(input) == 0 {
		return input
	}
	return append(reversedata(input[1:]), input[0])
}

func main() {
	var whatIsIt string
	secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
	sd, _ := base64.StdEncoding.DecodeString(secret)
	whatIsIt = string(reversedata(sd))
	fmt.Printf("Decode text is : " + whatIsIt)
}
