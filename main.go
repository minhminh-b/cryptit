package main

import (
	"fmt"
	"github.com/minhminh2003/cryptit/encrypt"
	"github.com/minhminh2003/cryptit/decrypt"
)

func main() {
	encryptedStr := encrypt.Nimbus("KodeKloud")
	fmt.Println(encryptedStr)
	fmt.Println(decrypt.Nimbus(encryptedStr))
}
