package repl

import (
	"bufio"
	"fmt"
	"io"
	"learn-interpreter/lexer"
	"learn-interpreter/token"
)

const Prompt = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprintf(out, Prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for t := l.NextToken(); t.Type != token.EOF; t = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", t)
		}
	}
}
