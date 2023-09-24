package repl

import (
	"arcane/evaluator"
	"arcane/lexer"
	"arcane/object"
	"arcane/parser"
	"bufio"
	"fmt"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Print(PROMPT)
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

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

const arcane_FACE = `
   __ _ _ __ ___ __ _ _ __   ___ 
  / _' | '__/ __/ _' | '_ \ / _ \
 | (_| | | | (_| (_| | | | |  __/
  \__'_|_|  \___\__'_|_| |_|\___|

`

func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, arcane_FACE)
	io.WriteString(out, "The parser isn't happy. Here's why:\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
