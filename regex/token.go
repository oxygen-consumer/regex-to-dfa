package regex

import "fmt"

type (
	Symbol   rune
	Operator rune
)

const (
	Union  Operator = '|'
	Concat Operator = '.'
	Kleene Operator = '*'
	LParen Operator = '('
	RParen Operator = ')'
)

type TokenKind int

const (
	SymbolToken TokenKind = iota
	OperatorToken
)

type Token struct {
	Kind  TokenKind
	Value rune
}

func NewSymbolToken(r rune) Token {
	return Token{Kind: SymbolToken, Value: r}
}

func NewOperatorToken(r rune) Token {
	return Token{Kind: OperatorToken, Value: r}
}

func (t Token) IsOperator() bool {
	return t.Kind == OperatorToken
}

func (t Token) IsSymbol() bool {
	return t.Kind == SymbolToken
}

func Tokenize(input []rune, alphabet map[Symbol]bool) ([]Token, error) {
	var tokens []Token
	for _, r := range input {
		switch {
		case IsOperator(r), r == '(' || r == ')':
			tokens = append(tokens, NewOperatorToken(r))
		case alphabet[Symbol(r)]:
			tokens = append(tokens, NewSymbolToken(r))
		default:
			return nil, fmt.Errorf("unknown token kind: %c", r)
		}
	}
	return tokens, nil
}

func IsOperator(r rune) bool {
	return r == '|' || r == '.' || r == '*'
}
