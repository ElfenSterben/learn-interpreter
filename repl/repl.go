package repl

import (
	"bufio"
	"fmt"
	"io"
	"learn-interpreter/eval"
	"learn-interpreter/lexer"
	"learn-interpreter/object"
	"learn-interpreter/parser"
)

const Prompt = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	macroEnv := object.NewEnvironment()

	for {
		fmt.Fprintf(out, Prompt)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		eval.DefineMacros(program, macroEnv)
		expanded := eval.ExpandMacros(program, macroEnv)
		evaluated := eval.Eval(expanded, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, " parser errors:\n")
		io.WriteString(out, "\t"+msg+"\n")
	}
}
