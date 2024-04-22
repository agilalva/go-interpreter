package lexer

import "monkey/token"

type Lexer struct {
	position     int
	readPosition int
	input        string
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	return l
}

func (l *Lexer) NextToken() token.Token {
	return token.Token{}
}
