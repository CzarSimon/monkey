package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/CzarSimon/monkey/evaluator"
	"github.com/CzarSimon/monkey/lexer"
	"github.com/CzarSimon/monkey/parser"
)

const PROMPT = ">> "

// Start Starts the repl
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		lex := lexer.New(line)
		p := parser.New(lex)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}
		evaluated := evaluator.Eval(program)
		io.WriteString(out, evaluated.Inspect())
		io.WriteString(out, "\n")
	}
}

func printParseErrors(out io.Writer, errors []error) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! Looks like we ran into some monkey bussiness here\n")
	io.WriteString(out, " parser errors:\n")
	for i, err := range errors {
		io.WriteString(out, fmt.Sprintf("\t%d. - %s\n", i, err.Error()))
	}
}

const MONKEY_FACE = `
  .--.  .-"     "-.  .--.
 / .. \/  .-. .-.  \/ .. \
| |  '|  /   Y   \  |'  | |
| \   \  \ 0 | 0 /  /   / |
 \ '- ,\.-"""""""-./, -' /
  ''-' /_   ^ ^   _\ '-''
      |  \._   _./  |
      \   \ '~' /   /
       '._ '-=-' _.'
          '-----'
`
