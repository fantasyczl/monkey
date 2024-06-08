package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/fantasyczl/monkey/evaluator"
	"github.com/fantasyczl/monkey/lexer"
	"github.com/fantasyczl/monkey/object"
	"github.com/fantasyczl/monkey/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	macroEnv := object.NewEnvironment()

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
		p := parser.New(l)
		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluator.DefineMacros(program, macroEnv)
		expanded := evaluator.ExpandMacros(program, macroEnv)

		evaluated := evaluator.Eval(expanded, env)
		if evaluated != nil {
			writeString(out, evaluated.Inspect())
			writeString(out, "\n")
		}
	}
}

const MONKEY_FACE = `
            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ ..  \
 | |  '|  /   Y   \  |'  |  |
 | \   \  \ 0 | 0 /  /   /  |
  \ '- ,\.-"""""""-./, -'  /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

func printParserErrors(out io.Writer, errors []string) {
	writeString(out, MONKEY_FACE)
	writeString(out, "Woops! We ran into some monkey business here!\n")
	writeString(out, " parser errors:\n")
	for _, msg := range errors {
		_, _ = out.Write([]byte("\t" + msg + "\n"))
	}
}

func writeString(out io.Writer, s string) {
	if _, err := io.WriteString(out, s); err != nil {
		panic(err)
	}
}
