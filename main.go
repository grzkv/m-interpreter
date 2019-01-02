package main

import (
	"os"

	"github.com/grzkv/m-interpreter/repl"
)

func main() {
	repl.RunREPL(os.Stdin, os.Stdout)
}
