package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

const (
	INT = "INT"

	LPAREN = "("
	RPAREN = ")"

	PLUS     = "+"
	MINUS    = "-"
	MULTIPLY = "*"
	DIVIDE   = "/"
)

type Token struct {
	flavour string
	literal string
	value   int
}

type Statement []*Token

func main() {
	// Read inputs into slice of int
	input := readInputs("input.txt")

	// Part 1 - parse, evaluate and sum statements with equal precedence
	precedence := map[string]int{
		LPAREN:   -1, // needed for parseInfix algorithm
		RPAREN:   -1, // needed for parseInfix algorithm
		PLUS:     1,
		MINUS:    1,
		MULTIPLY: 1,
		DIVIDE:   1,
	}
	total := processPart(input, precedence)
	fmt.Printf("Part 1 - sum of evaluated statements = %d\n", total)

	// Part 2 - parse, evaluate and sum statements with higher PLUS precedence
	precedence = map[string]int{
		LPAREN:   -1, // needed for parseInfix algorithm
		RPAREN:   -1, // needed for parseInfix algorithm
		PLUS:     5,
		MINUS:    1,
		MULTIPLY: 1,
		DIVIDE:   1,
	}
	total = processPart(input, precedence)
	fmt.Printf("Part 2 - sum of evaluated statements = %d\n", total)
}

// Stack is a simple implementation of a stack
type Stack []*Token

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(token *Token) {
	*s = append(*s, token)
}

func (s *Stack) Pop() (*Token, bool) {
	if s.isEmpty() {
		return nil, false
	}
	index := len(*s) - 1
	token := (*s)[index]
	*s = (*s)[:index]
	return token, true
}

func (s *Stack) asString() string {
	var value string
	for _, token := range *s {
		value += token.literal
	}
	value += "\n"
	return value
}

func processPart(input []string, precedence map[string]int) int {
	// Lex inputs into tokenised statements
	statements := lexInputs(input)

	// Parse tokenised statements and calculate total
	var total int
	for _, statement := range statements {
		postfix := statement.parseInfix(precedence)
		total += evaluatePostfix(&postfix)
	}
	return total
}

// Evaluate postfix form of statement
func evaluatePostfix(postfix *Stack) int {
	var stack Stack
	for _, token := range *postfix {
		if token.flavour == INT {
			stack.Push(token)
		} else {
			b, _ := stack.Pop()
			a, _ := stack.Pop()
			var value int
			switch {
			case token.flavour == PLUS:
				value = a.value + b.value
			case token.flavour == MINUS:
				value = a.value - b.value
			case token.flavour == MULTIPLY:
				value = a.value * b.value
			case token.flavour == DIVIDE:
				value = a.value / b.value
			}
			stack.Push(&Token{flavour: INT, value: value})
		}
	}
	// Should be left with one token containing the evaluated result
	if len(stack) != 1 {
		log.Fatalf("could not parse postfix")
	}
	value, _ := stack.Pop()
	return value.value
}

// Use shunting yard algorithm to parse infix form to postfix
func (s Statement) parseInfix(precedence map[string]int) Stack {
	var output Stack
	var operators Stack

	for _, token := range s {
		switch {
		case token.flavour == LPAREN:
			operators.Push(token)
		case token.flavour == RPAREN:
			for {
				if operators.isEmpty() {
					break
				}
				op, _ := operators.Pop()
				if op.flavour == LPAREN {
					break
				}
				output.Push(op)
			}
		case token.flavour == INT:
			output.Push(token)
		default: // Must be an operator defined by lexer
			if operators.isEmpty() {
				operators.Push(token)
			} else {
				op, _ := operators.Pop()
				if precedence[op.flavour] < precedence[token.flavour] {
					operators.Push(op)
					operators.Push(token)
				} else {
					output.Push(op)
					operators.Push(token)
				}
			}
		}
	}
	// Pop stack to output
	for !operators.isEmpty() {
		op, _ := operators.Pop()
		output.Push(op)
	}
	return output
}

// Lex input lines to tokens
func lexInputs(inputs []string) []Statement {
	var data []Statement

	for _, line := range inputs {
		if len(line) > 0 {
			// Build a tokenised statement from line
			statement := Statement{}
			for _, char := range []byte(line) {
				// Ignore whitespace
				if char == ' ' {
					continue
				}
				// Lex tokens
				token := &Token{literal: string(char)}
				switch {
				case char == '(':
					token.flavour = LPAREN
				case char == ')':
					token.flavour = RPAREN
				case char == '+':
					token.flavour = PLUS
				case char == '-':
					token.flavour = MINUS
				case char == '*':
					token.flavour = MULTIPLY
				case char == '/':
					token.flavour = DIVIDE
				// If not defined token above, then must be integer
				default:
					value, err := strconv.Atoi(string(char))
					if err != nil {
						log.Fatalf("unrecognised token %c", char)
					}
					token.flavour = INT
					token.value = value
				}
				statement = append(statement, token)
			}
			data = append(data, statement)
		}
	}
	return data
}

func readInputs(filename string) []string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("failed to open input.txt")
	}
	lines := string(b)

	var inputs []string
	for _, line := range strings.Split(lines, "\n") {
		inputs = append(inputs, line)
	}
	return inputs
}
