package repl

import (
	"bufio"
	"fmt"
	"io"
	"github.com/KoMaTop10/Monkey/lexer"
	"github.com/KoMaTop10/Monkey/parser"
	"github.com/KoMaTop10/Monkey/evaluator"
	"github.com/KoMaTop10/Monkey/object"
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

		evaluated := evaluator.Eval(program,env)

		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out,"\n")
		} else {
			io.WriteString(out,MONKEY_FACE)
		}
		
	}

}

const MONKEY_FACE = `
            __,__
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

func printParserErrors(out io.Writer,errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t" + msg + "\n")
	}
}