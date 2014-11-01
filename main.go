package main

import (
	"fmt"
	"github.com/kchunterdeluxe/simple-compiler/token"
)

func main() {
	t := token.New(token.MAIN, "MAIN")
	fmt.Printf("%v", t)
}
