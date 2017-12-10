package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/CzarSimon/monkey/lexer"
	"github.com/CzarSimon/monkey/token"
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
		for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
