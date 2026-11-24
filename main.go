package main

import (
	"fmt"

	"github.com/Hemal99/cryptit/decrypt"
	"github.com/Hemal99/cryptit/encrypt"
)

func main() {
	encryptedString := encrypt.Nimbus("IMhemal")
	fmt.Println(encryptedString)
	fmt.Println(decrypt.Nimbus(encryptedString))
}
