package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	s := "sha256 this string"
	fmt.Println(s)
	h := sha256.New()
	h.Write([]byte(s))
	fmt.Println(h)
	bs := h.Sum(nil)
	fmt.Println(bs)

	fmt.Printf("%x\n", bs)

	h1 := sha256.Sum256([]byte(s))
	fmt.Printf("%x\n", h1)

}
