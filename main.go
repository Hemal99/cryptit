package main

import (
	"fmt"

	"github.com/Hemal99/cryptit.git/decrypt"
	"github.com/Hemal99/cryptit.git/encrypt"
)

func main() {
	encryptedString := encrypt.Nimbus("IMhemal")
	fmt.Println(encryptedString)
	fmt.Println(decrypt.Nimbus(encryptedString))
}
