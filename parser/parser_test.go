package parser

import (
	"testing"
)

func TestParsing(t *testing.T) {
	input := `
  let five = 5;
  let f = 8 + 2;
  let function = func() {}
  let add = func(a, bcd) {
    return a + b;
  }
  `

	expectedResult := []Statement{
		&LetStatement{Identifier: "five", Value: &NumericExpression{LeftOperator: 5}},
		&LetStatement{Identifier: "f", Value: &NumericExpression{LeftOperator: 8, Operand: "+", RightOperator: &NumericExpression{LeftOperator: 2}}},
		&LetStatement{Identifier: "function", Value: &FunctionExpression{}},
		&LetStatement{Identifier: "add", Value: &FunctionExpression{Parameters: []string{"a", "bcd"}}},
	}

	parser := NewParser(input)
	statements := parser.ParseProgram()

	for i, st := range expectedResult {
		if !st.statementIsEqual(statements[i]) {
			t.Fatalf("execpted %+v got %+v\n", statements[i], st)
		}
	}
}
