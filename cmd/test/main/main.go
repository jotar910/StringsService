package main

import (
	"fmt"

	"joaor.dev.com/joao/stringservice/pkg/parser"
)

func main() {
	fmt.Printf("%v\n", parser.Parse("A  B  c     s", "l"))
	fmt.Printf("%v\n", parser.Parse("A  B  c     s", "la"))
	fmt.Printf("%v\n", parser.Parse("A  B  c     s", "laf"))
}
