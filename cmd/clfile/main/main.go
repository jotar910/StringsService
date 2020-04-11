package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"joaor.dev.com/joao/stringservice/pkg/parser"
)

func main() {
	filename := flag.String("filename", "content.txt", "the filename you want to parse")
	flags := flag.String("flags", "", "the code sequence")
	flag.Parse()

	if !strings.HasSuffix(*filename, ".txt") {
		log.Fatalf("Syntax error!\n\t <command> -filename <filename> -flags <flag1[,flags...]>")
	}

	file, err := os.Open(*filename)
	if err != nil {
		fmt.Printf("Error trying to open file with name %q\n\t%+v\n", *filename, err)
		os.Exit(1)
	}

	fi, err := file.Stat()
	if err != nil {
		fmt.Printf("Error trying to read file information with name %q\n\t%+v\n", *filename, err)
		os.Exit(1)
	}

	textbytes := make([]byte, fi.Size())
	_, err = file.Read(textbytes)
	if err != nil {
		fmt.Printf("Error trying to read file with name %q\n\t%+v\n", *filename, err)
		os.Exit(1)
	}

	fmt.Println(parser.Parse(string(textbytes), *flags))
}
