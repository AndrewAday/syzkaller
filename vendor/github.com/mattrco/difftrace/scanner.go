package parser

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"fmt"
)

type Token int

const (
	ILLEGAL Token = iota
	EOF
	NEWLINE
	WS          // Whitespace
	IDENT       // Identifier, e.g. function name.
	MEMADDR     // Hex address, e.g. 0xcafef00d
	OPEN_PAREN  // (
	CLOSE_PAREN // )
	OPEN_BRACE  // {
	CLOSE_BRACE // }
	OPEN_SQ     // [
	CLOSE_SQ    // ]
	PIPE        // |
	STRING      // Delimited by "
	SEP         // ,
	EQUALS      // =
	SIGNAL      // ---
	THROWAWAY   // This is meant to handle things like 0777
	POINTER     // 0x[addr]:=[data]
)

var terminator = map[rune]bool{
	',': true,
	']': true,
	'}': true,
	')': true,
}

var eof = rune(0)
var ErrEOF = errors.New("EOF")

// As it is useful to process newline differently to other
// whitespace don't include it here.
func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t'
}

func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || ch == '|'
}

func isDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}

// Scanner is a lexical scanner implemented with a buffered reader.
type Scanner struct {
	r *bufio.Reader
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{r: bufio.NewReader(r)}
}

// read reads the next rune from the buffered reader.
// Returns rune(0) if an error occurs.
func (s *Scanner) read() rune {
	r, _, err := s.r.ReadRune()
	if err != nil { // last line is format +++ exited with [] +++
		return eof
	}
	return r
}

// unreadRune places the previously read rune back on the reader.
func (s *Scanner) unreadRune() error { return s.r.UnreadRune() }

// unreadRunes allows a number of runes to be unread.
func (s *Scanner) unreadRunes(runes int) error {
	for runes > 0 {
		if err := s.r.UnreadRune(); err != nil {
			return err
		}
		runes--
	}
	return nil
}

// Scan returns the next token and literal string it represents.
func (s *Scanner) Scan() (tok Token, lit string) {
	// Read the next rune.
	r := s.read()

	if r == '0' {
		// 0 could be the start of a memory address. Read the next
		// rune and if it is 'x', pass to scanAddress.
		if r := s.read(); r == 'x' {
			return s.scanAddress()
		} else {
			s.unreadRunes(2)
			if isDigit(r) {
				return THROWAWAY, string(r)
			}
		}
	} else if r == '/'{
		var prev rune
		if r := s.read(); r == '*' {
			prev = r
			for {
				r = s.read()
				if prev == '*' && r =='/' {
					break
				}
				prev = r
			}
			return s.Scan()
		} else {
			s.unreadRunes(2)
		}
	} else if r=='&' {
		// then r is a pointer
		// read in the 0x
		s.read()
		s.read()
		tok, lit = s.scanAddress()
		lit = "&" + lit
		return tok, lit
	} else if r == '-' {
		// - could be a signal (lines start with ---).
		r = s.read()
		if r == '-' {
			r = s.read()
			if r == '-' {
				return SIGNAL, "---"
			}
		} else if isDigit(r) {
			s.unreadRune()
			return s.scanNegative()
		} else {
			return ILLEGAL, ""
		}
	} else if r == '"' {
		return s.scanString()
	} else if isWhitespace(r) {
		// Consume all contiguous whitespace.
		s.unreadRune()
		return s.scanWhitespace()
	} else if isLetter(r) || isDigit(r) {
		// If we see a letter then consume as an identifier.
		s.unreadRune()
		return s.scanIdent()
	}

	// Otherwise read the individual character.
	switch r {
	case eof:
		return EOF, ""
	case '(':
		return OPEN_PAREN, string(r)
	case ',':
		return SEP, string(r)
	case ')':
		return CLOSE_PAREN, string(r)
	case '{':
		return OPEN_BRACE, string(r)
	case '}':
		return CLOSE_BRACE, string(r)
	case '[':
		return OPEN_SQ, string(r)
	case ']':
		return CLOSE_SQ, string(r)
	case '|':
		return PIPE, string(r)
	case '=':
		return EQUALS, string(r)
	case '\n':
		return NEWLINE, string(r)
	case '~':
		if r := s.read(); r == '[' {
			return OPEN_SQ, "~["
		}
		s.unreadRune()
	}

	return ILLEGAL, string(r)
}

// scanAddress consumes a memory address from the scanner.
func (s *Scanner) scanAddress() (tok Token, lit string) {
	var buf bytes.Buffer
	// Prepend 0x as memory addresses are always hex.
	buf.WriteString("0x")

	// Read until there are no letters or digits.
	var r rune
	for {
		r = s.read()
		if r == eof {
			break
		} else if !(isDigit(r) || isLetter(r)) {
			s.unreadRune()
			break
		}
		buf.WriteRune(r)
	}

	if r == eof {
		return MEMADDR, buf.String()
	}

	// check to see if this is a pointer to data, in form 0x[addr]:=[data]
	r = s.read()
	if r != '=' {
		s.unreadRune()
		return MEMADDR, buf.String()
	}
	buf.WriteRune(r)
	if r = s.read(); r != '"' {
		fmt.Printf("unexpected pointer format, expected 0x[addr]=\"[data]\", got %v\n", buf.String())
		return MEMADDR, buf.String()
	}
	_, str := s.scanString()
	buf.WriteString(str)

	return POINTER, buf.String()
}

// scanString consumes a string from the scanner.
func (s *Scanner) scanString() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune('"')

	// Read up to the next inverted comma.
	for {
		r := s.read()
		if r == eof {
			break
		}
		if r == '"' {
			next := s.read()
			_,ok := terminator[next]
			//if next != ',' {
			//	s.unreadRune()
			//}
			s.unreadRune()
			if next == eof || ok {
				// Reached end of literal. Consume ellipsis if present.
				buf.WriteRune(r)
				if r = s.read(); r == '.' {
					s.read()
					s.read()
					buf.WriteString("...")
				} else {
					s.unreadRune()
				}
				break
			}
		}
		buf.WriteRune(r)
	}

	return STRING, buf.String()
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *Scanner) scanWhitespace() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent whitespace character into the buffer.
	for {
		if r := s.read(); r == eof {
			break
		} else if !isWhitespace(r) {
			s.unreadRune()
			break
		} else {
			buf.WriteRune(r)
		}
	}

	return WS, buf.String()
}

// scanIdent consumes the current rune and all contiguous ident runes.
func (s *Scanner) scanIdent() (tok Token, lit string) {
	// Create a buffer and read the current character into it.
	var buf bytes.Buffer
	buf.WriteRune(s.read())

	// Read every subsequent ident character into the buffer.
	// Non-ident characters and EOF will cause the loop to exit.
	for {
		if r := s.read(); r == eof {
			break
		} else if !isLetter(r) && !isDigit(r) && r != '_' {
			s.unreadRune()
			break
		} else {
			buf.WriteRune(r)
		}
	}

	// Otherwise return as a regular identifier.
	return IDENT, buf.String()
}

func (s *Scanner) scanNegative() (tok Token, lit string) {
	var buf bytes.Buffer
	buf.WriteRune('-')

	for {
		if r := s.read(); r == eof {
			break
		} else if !isDigit(r) {
			s.unreadRune()
			break
		} else {
			buf.WriteRune(r)
		}
	}

	// Return as a regular identifier.
	return IDENT, buf.String()
}
