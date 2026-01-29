package main

import (
	"fmt"

	// TODO: for now goos package must be loaded to cache any nostd module
	// imported by GOOSPKG.
	_ "github.com/abarisani/golang-proposal-73608/goos"
)

func main() {
	fmt.Println("hello world")
}
