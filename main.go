package main

import (
	"fmt"
	"os"

	"github.com/wu-xian/ast-dumper/syntax"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:")
		fmt.Println("\tast-dumper [filename]")
		os.Exit(1)
	}
	fname := os.Args[1]
	errh := syntax.ErrorHandler(
		func(err error) {
			fmt.Println(err)
		})

	mode := 0 | syntax.AllowGenerics

	ast, _ := syntax.ParseFile(
		fname,
		errh,
		nil,
		mode)
	f, _ := os.Create(fname + ".ast")
	defer f.Close()
	syntax.Fdump(f, ast) //<--this prints out your AST nicely

}
