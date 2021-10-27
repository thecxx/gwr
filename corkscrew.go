package gwr

import (
	"errors"
)

// type express [2]string

// const (
// 	expressKey   = 0
// 	expressValue = 1
// )

// type WrParameter struct {
// 	expresses []express
// }

type TokenType uint8

type Token struct {
	Type  TokenType
	Value string
}

const (
	TokenUnknown TokenType = iota
	TokenField
	TokenNumber
)

type lexer struct {
	err    error
	syntax string
	tokens []string
}

func Scan(syntax string) (lex *lexer) {
	n := len(syntax)
	if n < 2 {
		return &lexer{err: errors.New("illegal expression")}
	}

	lex = &lexer{syntax: syntax}

	for i := 0; i < n; i++ {
		c := syntax[i]
		// Lex
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
			'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
			'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
			'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
		case '{':
		case '}':
		case '(':
		case ')':
		case '.':
		case ':':
		case '/':
		case '^':
		case '$':
		case '#':
		case '~', '`', '!', '@', '%', '&', '*', '_', '+', '-', '=', '[', ']', '|', '\\', ';', '"', '\'', '<', '>', ',', '?':
		case ' ', '\t', '\n', '\r':
		default:
			lex.err = errors.New("unrecognized character")
			return
		}
	}

	return lex
}

type parser struct {
	lex *lexer
	ast interface{}
}

func Parse(lex *lexer) (*parser, error) {
	return &parser{lex: lex}, nil
}

func (p *parser) GetAbstractSyntaxTree() interface{} {
	return p.ast
}

type corkscrew struct {
	opcodes interface{}
}

func (c *corkscrew) extract(vptr interface{}) {

}
