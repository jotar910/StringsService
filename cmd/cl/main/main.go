package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"joaor.dev.com/joao/stringservice/pkg/parser"
)

func main() {
	tStr := flag.String("t", "", "the text you want to parse")
	fStr := flag.String("f", "", "the code sequence")
	flag.Parse()
	if len(*tStr) == 0 {
		log.Fatalf("Syntax error!\n\t <command> -t <text> -f <flag1[,flags...]>")
	}

	fs := strings.Split(*fStr, ",")

	fmt.Println(parser.Parse(*tStr, fs))
}
