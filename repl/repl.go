package repl

import (
	"bufio"
	"fmt"
	"monkey/lexer"
	"monkey/token"
	"os"
)

func StartRepl() {
	sc := bufio.NewScanner(os.Stdin)
	fmt.Print(">>> ")
	for sc.Scan() {
		input := sc.Text()
		l := lexer.NewLexer(input)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n",tok)
		}

		fmt.Print(">>> ")
	}
}
