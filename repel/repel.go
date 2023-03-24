package repel

import (
	"Nova/lexer"
	"Nova/token"
	"bufio"
	"fmt"
	"io"
	"os"
)

const PROMPT = "💫 "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if line == "quit" {
			os.Exit(1)
		}

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
