package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/grzkv/m-interpreter/lexer"
	"github.com/grzkv/m-interpreter/token"
)

// RunREPL runs Monkey REPL
func RunREPL(r io.Reader, w io.Writer) {
	var PROMPT = "> "

	scnr := bufio.NewScanner(r)

	for {
		fmt.Fprint(w, PROMPT)

		if scnr.Scan() {
			line := scnr.Text()

			l := lexer.New(line)

			for {
				t := l.NextToken()
				fmt.Fprintf(w, "typ: %s # literal: %s\n", t.Typ, t.Literal)
				if t.Typ == token.EOF {
					break
				}
			}
		}

	}
}
