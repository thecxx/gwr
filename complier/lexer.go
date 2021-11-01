package complier

import (
	"bytes"
	"errors"
	"fmt"
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
	TokenInteger
	TokenIdentify
	TokenVariable
)

var (
	ErrEOF = errors.New("EOF")
)

type lexer struct {
	expression string
	pos        int
	tokens     []Token
}

func Scan(expression string) (lex *lexer, err error) {
	n := len(expression)
	if n < 2 {
		return nil, errors.New("illegal expression")
	}

	lex = &lexer{
		pos:        -1,
		expression: expression,
	}

	for {
		ch, err := lex.next()
		if err != nil {
			break
		}
		// Lex
		switch ch {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			lex.ParseInteger(string(ch))
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
			'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
			'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
			'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
			'_':
			lex.ParseIdentify(string(ch))
		case '{':
			lex.tokens = append(lex.tokens, Token{Value: "{"})
		case '}':
			lex.tokens = append(lex.tokens, Token{Value: "}"})
		case '(':
			lex.tokens = append(lex.tokens, Token{Value: "("})
		case ')':
			lex.tokens = append(lex.tokens, Token{Value: ")"})
		case '.':
			lex.tokens = append(lex.tokens, Token{Value: "."})
		case ':':
			lex.tokens = append(lex.tokens, Token{Value: ":"})
		case '/':
			lex.ParseRegular("/")
		case '^':
		case '$':
			lex.ParseVariable("$")
		case ',':
			lex.tokens = append(lex.tokens, Token{Value: ","})
		case '@':
			lex.tokens = append(lex.tokens, Token{Value: "@"})
		case '<':
			lex.tokens = append(lex.tokens, Token{Value: "<"})
		case '>':
			lex.tokens = append(lex.tokens, Token{Value: ">"})
		case '|':
			nch, err := lex.next()
			if err != nil {
				lex.tokens = append(lex.tokens, Token{Value: "|"})
			} else if nch == '|' {
				lex.tokens = append(lex.tokens, Token{Value: "||"})
			}
		case '#':
			lex.tokens = append(lex.tokens, Token{Value: "#"})
		case '"':
			lex.ParseString("\"")
		case ' ', '\t', '\n', '\r':
			lex.SkipWhiteSpaces(string(ch))
		default:
			fmt.Printf("%s\n", string(ch))
			return nil, errors.New("unrecognized character")
		}
	}

	return lex, nil
}

func (l *lexer) next() (byte, error) {
	n := len(l.expression)
	if n == 0 || n <= (l.pos+1) {
		return 0, ErrEOF
	}
	// Move to next
	l.pos++
	// Shift
	return l.expression[l.pos], nil
}

func (l *lexer) back() {
	if l.pos >= 0 {
		l.pos--
	}
}

func (l *lexer) SkipWhiteSpaces(firstK string) {
	// buf := bytes.NewBufferString(firstK)
	for {
		ch, err := l.next()
		if err != nil {
			return
		}
		switch ch {
		case ' ':
			fallthrough
		case '\r':
			fallthrough
		case '\n':
			fallthrough
		case '\t':
			// Ignore
		default:
			l.back()
			return
		}
	}
}

func (l *lexer) ParseInteger(firstK string) {
	buf := bytes.NewBufferString(firstK)
	for {
		ch, err := l.next()
		if err != nil {
			l.tokens = append(l.tokens, Token{Type: TokenInteger, Value: buf.String()})
			return
		}
		switch ch {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.':
			buf.WriteByte(ch)
		default:
			l.back()
			l.tokens = append(l.tokens, Token{Type: TokenInteger, Value: buf.String()})
			return
		}
	}
}

func (l *lexer) ParseIdentify(firstK string) {
	buf := bytes.NewBufferString(firstK)
	for {
		ch, err := l.next()
		if err != nil {
			l.tokens = append(l.tokens, Token{Type: TokenIdentify, Value: buf.String()})
			return
		}
		switch ch {
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
			'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
			'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
			'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
			'_':
			buf.WriteByte(ch)
		default:
			l.back()
			l.tokens = append(l.tokens, Token{Type: TokenIdentify, Value: buf.String()})
			return
		}
	}
}

func (l *lexer) ParseVariable(firstK string) {
	buf := bytes.NewBufferString(firstK)
	for {
		ch, err := l.next()
		if err != nil {
			l.tokens = append(l.tokens, Token{Type: TokenVariable, Value: buf.String()})
			return
		}
		switch ch {
		case '0':
			buf.WriteByte(ch)
		default:
			l.back()
			l.tokens = append(l.tokens, Token{Type: TokenVariable, Value: buf.String()})
			return
		}
	}
}

func (l *lexer) ParseRegular(firstK string) {
	buf := bytes.NewBufferString(firstK)
	for {
		ch, err := l.next()
		if err != nil {
			l.tokens = append(l.tokens, Token{Type: TokenIdentify, Value: buf.String()})
			return
		}
		switch ch {
		case '/':
			buf.WriteByte(ch)
			l.tokens = append(l.tokens, Token{Type: TokenIdentify, Value: buf.String()})
			return
		default:
			buf.WriteByte(ch)
		}
	}
}

func (l *lexer) ParseString(firstK string) {
	buf := bytes.NewBufferString("")
	for {
		ch, err := l.next()
		if err != nil {
			l.tokens = append(l.tokens, Token{Type: TokenIdentify, Value: buf.String()})
			return
		}
		switch ch {
		case '"':
			l.tokens = append(l.tokens, Token{Type: TokenIdentify, Value: buf.String()})
			return
		default:
			buf.WriteByte(ch)
		}
	}
}
