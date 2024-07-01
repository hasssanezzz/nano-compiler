package lexer

import (
	"fmt"
	"unicode"
)

// ? where i left off
// the read_... functions read an extra byte to break
// therefore they discard the '(', ')'

type Lexer struct {
	Source []rune
}

func (l *Lexer) lexNumber(cursor int) (int, *Token) {
	curr := ""
	location := cursor

	for cursor < len(l.Source) {
		char := l.Source[cursor]
		if !(unicode.IsDigit(char) || char == '-' || char == '.') {
			break
		}
		curr += string(char)
		cursor++
	}

	if len(curr) == 0 {
		return cursor, nil
	}

	return cursor, &Token{
		Type:     Literal,
		Value:    curr,
		Location: uint32(location),
	}
}

func (l *Lexer) lexIdentifier(cursor int) (int, *Token) {
	curr := ""
	location := cursor

	for cursor < len(l.Source) {
		char := l.Source[cursor]
		if !unicode.IsLetter(char) {
			break
		}

		curr += string(char)
		cursor++
	}

	t := &Token{
		Value:    curr,
		Location: uint32(location),
	}

	if contains(curr, KeywordSet) != -1 {
		t.Type = Keyword
	} else {
		t.Type = Identifier
	}

	return cursor, t
}

func (l *Lexer) lexString(cursor int) (int, *Token) {
	curr := "\""
	location := cursor

	endFound := false

	for ; cursor < len(l.Source); cursor++ {
		char := l.Source[cursor]
		if char == '\\' {
			nextChar := l.Source[cursor+1]
			if nextChar == '"' {
				curr += string(nextChar)
				cursor++
				continue
			}
		}

		curr += string(char)
		if char == '"' {
			endFound = true
			cursor++
			break
		}
	}

	if !endFound {
		panic("lexer: could't lex string literal")
	}

	return cursor, &Token{
		Type:     Literal,
		Value:    curr,
		Location: uint32(location),
	}
}

func (l *Lexer) lexLiteral(cursor int) (int, *Token) {
	curr := ""
	location := cursor

	for cursor < len(l.Source) {
		char := l.Source[cursor]
		if unicode.IsSpace(char) || char == ')' {
			break
		}

		curr += string(l.Source[cursor])
		cursor++
	}

	if len(curr) == 0 {
		return cursor, nil
	}

	if contains(curr, OperatorSet) != -1 {
		return cursor, &Token{
			Type:     Operator,
			Value:    curr,
			Location: uint32(location),
		}
	}

	return cursor, &Token{
		Type:     Literal,
		Value:    curr,
		Location: uint32(location),
	}
}

func (l *Lexer) Lex() []Token {
	var tokens []Token

	cursor := 0
	for cursor < len(l.Source) {
		char := l.Source[cursor]

		switch true {
		case unicode.IsDigit(char):
			newCursor, token := l.lexNumber(cursor)
			tokens = append(tokens, *token)
			cursor = newCursor

		case unicode.IsLetter(char):
			newCursor, token := l.lexIdentifier(cursor)
			tokens = append(tokens, *token)
			cursor = newCursor

		case char == '"':
			newCursor, token := l.lexString(cursor + 1)
			tokens = append(tokens, *token)
			cursor = newCursor

		case char == '(':
			tokens = append(tokens, Token{Type: Separator, Value: "(", Location: uint32(cursor)})
			cursor++

		case char == ')':
			tokens = append(tokens, Token{Type: Separator, Value: ")", Location: uint32(cursor)})
			cursor++

		default:
			if unicode.IsSpace(char) {
				cursor++
				continue
			}

			newCursor, token := l.lexLiteral(cursor)
			if token != nil {
				cursor = newCursor
				tokens = append(tokens, *token)
			} else {
				fmt.Printf("what to do here?: %q\n", char)
				cursor++
			}
		}

	}

	return tokens
}

func Lex(source []rune) {
	l := Lexer{source}

	tokens := l.Lex()

	for _, token := range tokens {
		fmt.Printf("%v\n", token)
	}
}
