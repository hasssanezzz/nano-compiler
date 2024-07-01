package lexer

import "fmt"

type TokenType int

type Token struct {
	Type     TokenType
	Value    string
	Location uint32
}

type AST struct {
	Value    Token
	Children []AST
}

const (
	Identifier TokenType = iota
	Keyword
	Separator
	Operator
	Literal
	Whitespace
)

func (tt TokenType) String() string {
	switch tt {
	case Identifier:
		return "Identifier"
	case Keyword:
		return "Keyword"
	case Separator:
		return "Separator"
	case Operator:
		return "Operator"
	case Literal:
		return "Literal"
	case Whitespace:
		return "Whitespace"
	default:
		return fmt.Sprintf("%d", int(tt))
	}
}

var KeywordSet = []string{"true", "false", "nil"}
var OperatorSet = []string{"+", "-", "*", ">", "<", ">=", "<=", "!=", "==", "or", "and", "set", "macth", "while"}
