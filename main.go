package main

import (
	"fmt"
	"os"

	"github.com/hasssanezzz/nano-compiler/lexer"
)

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	stat, _ := file.Stat()
	buff := make([]byte, stat.Size())
	n, err := file.Read(buff)
	if err != nil {
		panic(err)
	}
	if n != int(stat.Size()) {
		panic(fmt.Errorf("number of read bytes doesn't match file size %d != %d", n, stat.Size()))
	}

	tokens := (&lexer.Lexer{Source: []rune(string(buff))}).Lex()
	for _, token := range tokens {
		fmt.Printf("%v\n", token)
	}
	(&lexer.Parser{Tokens: tokens}).Start()
}
