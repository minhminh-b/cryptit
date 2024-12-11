package main

import (
	"fmt"
	"github.com/minhminh-b/cryptit/encrypt"
	"github.com/minhminh-b/cryptit/decrypt"
)

func main() {
	encryptedStr := encrypt.Nimbus("KodeKloud")
	fmt.Println(encryptedStr)
	fmt.Println(decrypt.Nimbus(encryptedStr))
}
