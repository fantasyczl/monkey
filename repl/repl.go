package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/fantasyczl/monkey/lexer"
	"github.com/fantasyczl/monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		if _, err := fmt.Fprint(out, PROMPT); err != nil {
			_, _ = out.Write([]byte(err.Error()))
			break
		}

		scanned := scanner.Scan()
		if !scanned {
			break
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			if _, err := fmt.Fprintf(out, "%+v\n", tok); err != nil {
				_, _ = out.Write([]byte(err.Error()))
				break
			}
		}
	}
}
