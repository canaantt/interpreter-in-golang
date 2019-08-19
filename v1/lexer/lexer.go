package lexer


import(
	"github.com/canaantt/interpreter/v1/token"
)

type Lexer struct {
	input string
	currPos int
	nextPos int
	ch byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.read()
	return l
}

func (l *Lexer) read() {
	if l.nextPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.nextPos]
	}
	// both pointers step over
	l.currPos = l.nextPos
	l.nextPos = l.nextPos + 1
}

func (l *Lexer) GetNextChar() byte {
	if l.nextPos >= len(l.input) {
		return 0
	} else {
		return l.input[l.nextPos]
	}
}

func (l *Lexer) GetToken() token.Token {
	var tok token.Token

	//l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.GetNextChar() == '=' {
			ch := l.ch
			l.read()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.GetNextChar() == '=' {
			ch := l.ch
			l.read()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.
				LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.read()
	return tok
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' ||
		   'A' <= ch && ch <= 'Z' ||
		   ch == '_'
}

func (l *Lexer) readIdentifier() string {
	position := l.currPos
	for isLetter(l.ch) {
		l.read()
	}
	return l.input[position:l.currPos]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.currPos
	for isDigit(l.ch) {
		l.read()
	}
	return l.input[position:l.currPos]
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
