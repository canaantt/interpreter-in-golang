package repl

import (
	"bufio"
	"fmt"
	"io"
	"github.com/canaantt/interpreter/v1/lexer"
	"github.com/canaantt/interpreter/v1/token"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.NewLexer(line)

		for tok := l.GetToken(); tok.Type != token.EOF; tok = l.GetToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}