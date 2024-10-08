package ast

import (
	"reflect"
	"testing"
)

func TestModify(t *testing.T) {
	one := func() Expression {
		return &IntegerLiteral{Value: 1}
	}
	two := func() Expression {
		return &IntegerLiteral{Value: 2}
	}

	turnOneIntoTwo := func(node Node) Node {
		integer, ok := node.(*IntegerLiteral)
		if !ok {
			return node
		}
		if integer.Value != 1 {
			return node
		}

		integer.Value = 2
		return integer
	}

	tests := []struct {
		input    Node
		expected Node
	}{
		{
			input:    one(),
			expected: two(),
		},
		{
			input: &Program{
				Statements: []Statement{
					&ExpressionStatement{Expression: one()},
				},
			},
			expected: &Program{
				Statements: []Statement{
					&ExpressionStatement{Expression: two()},
				},
			},
		},
		{
			&InfixExpression{Left: one(), Operator: "+", Right: two()},
			&InfixExpression{Left: two(), Operator: "+", Right: two()},
		},
		{
			&InfixExpression{Left: two(), Operator: "+", Right: one()},
			&InfixExpression{Left: two(), Operator: "+", Right: two()},
		},
		{
			&PrefixExpression{Operator: "-", Right: one()},
			&PrefixExpression{Operator: "-", Right: two()},
		},
		{
			&IndexExpression{Left: one(), Index: one()},
			&IndexExpression{Left: two(), Index: two()},
		},
		{
			&IfExpression{
				Condition: one(),
				Consequence: &BlockStatement{
					Statements: []Statement{
						&ExpressionStatement{Expression: one()},
					},
				},
				Alternative: &BlockStatement{
					Statements: []Statement{
						&ExpressionStatement{Expression: one()},
					},
				},
			},
			&IfExpression{
				Condition: two(),
				Consequence: &BlockStatement{
					Statements: []Statement{
						&ExpressionStatement{Expression: two()},
					},
				},
				Alternative: &BlockStatement{
					Statements: []Statement{
						&ExpressionStatement{Expression: two()},
					},
				},
			},
		},
		{
			&ReturnStatement{ReturnValue: one()},
			&ReturnStatement{ReturnValue: two()},
		},
		{
			&LetStatement{Value: one()},
			&LetStatement{Value: two()},
		},
		{
			&FunctionLiteral{
				Parameters: []*Identifier{},
				Body: &BlockStatement{
					Statements: []Statement{
						&ExpressionStatement{Expression: one()},
					},
				},
			},
			&FunctionLiteral{
				Parameters: []*Identifier{},
				Body: &BlockStatement{
					Statements: []Statement{
						&ExpressionStatement{Expression: two()},
					},
				},
			},
		},
		{
			input:    &ArrayLiteral{Elements: []Expression{one(), one()}},
			expected: &ArrayLiteral{Elements: []Expression{two(), two()}},
		},
	}

	for _, tt := range tests {
		modified := Modify(tt.input, turnOneIntoTwo)
		equal := reflect.DeepEqual(modified, tt.expected)
		if !equal {
			t.Errorf("not equal. got=%#v\nwant=%#v", modified, tt.expected)
		}
	}

	testModifyHashLiteral(t, one, two, turnOneIntoTwo)
}

func testModifyHashLiteral(t *testing.T, one, two func() Expression, turnOneIntoTwo func(node Node) Node) {
	hashLiteral := &HashLiteral{
		Pairs: map[Expression]Expression{
			one(): one(),
			one(): one(),
		},
	}

	Modify(hashLiteral, turnOneIntoTwo)
	for key, val := range hashLiteral.Pairs {
		key, _ := key.(*IntegerLiteral)
		if key.Value != 2 {
			t.Errorf("key is not 2. got=%d", key.Value)
		}
		val, _ := val.(*IntegerLiteral)
		if val.Value != 2 {
			t.Errorf("val is not 2. got=%d", val.Value)
		}
	}
}
