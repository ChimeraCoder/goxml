package main

import (
	"fmt"
    "log"
	"regexp"
	"strings"
	"unicode/utf8"
)

// item represents a token returned from the scanner.
type item struct {
	typ itemType // Type, such as itemNumber.
	val string   // Value, such as "23.2".
}

type itemType int

const (
	itemError itemType = iota //error occurred

	itemLeftBrace
	itemNumber
	itemRightBrace
	itemQuote
	itemIdentifier

	itemDot
	itemEOF
)

const leftMeta = '{'
const rightMeta = '}'

const EOF = 0

func (i item) String() string {
	switch i.typ {
	case itemEOF:
		return "EOF"
	case itemError:
		return i.val
	}
	return fmt.Sprintf("%q", i.val)
}

// lexer holds the state of the scanner
type lexer struct {
	name  string    // for error reporting
	input string    // string being scanned
	start int       // start position of item
	pos   int       // current position in input
	width int       // width of last rune read from input
	items chan item // channel of scanned items
}

type stateFn func(*lexer) stateFn

func lex(name string, input string) (*lexer, chan item) {
	l := &lexer{
		name:  name,
		input: input,
		items: make(chan item, 2), // two items is sufficient
	}

	go l.run()
	return l, l.items
}

func (l *lexer) run() {
	for state := startState; state != nil; {
		state = state(l)
	}
	close(l.items) // no more tokens will be delivered
}

func (l *lexer) emit(t itemType) {
    i := item{t, l.input[l.start:l.pos]}
	l.items <- i
	l.start = l.pos
}

func lexText(l *lexer) stateFn {
	for {
		// check if the string starts with {
		if strings.HasPrefix(l.input[l.pos:], string(leftMeta)) {
			if l.pos > l.start {
				l.emit(itemLeftBrace)
			}
			// return the next state
			return lexInsideAction
		}
		if l.next() == EOF {
			break
		}
	}
	// Correctly reached EOF
	if l.pos > l.start {
		l.emit(itemEOF)
	}

	l.emit(itemEOF) // useful to make EOF a token
	return nil      // this will stop the run loop
}

// lexInsideAction means we are inside a new block
// having just read {
func lexInsideAction(l *lexer) stateFn {
	// Ether a quoted string or a closing brace
	// TODO or a non-quoted string
	// spaces separate and are ignored

	for {
		switch r := l.next(); {
		case r == EOF:
			return l.errorf("unclosed action")
        case r == rightMeta:
            l.emit(itemRightBrace)
            // TODO fix this
            return startState
		case isSpace(r):
			l.ignore()
		case r == '"':
			l.emit(itemQuote)
		case r == '+' || r == '-' || '0' <= r && r <= '9':
			l.backup()
			return lexNumber
		case isAlphaNumeric(r):
			l.backup()
			return lexIdentifier
		}
	}
}

// lexIdentifier means we are reading an identifier
func lexIdentifier(l *lexer) stateFn {
	const alphanumeric = "ABCDEFGHIJKLMNOPQRSTUVQXYZabcdefghijklmnopqrstuvwxyz0123456789"
	l.acceptRun(alphanumeric)
	l.emit(itemIdentifier)
	return startState
}

// next returns the next rune in the input
func (l *lexer) next() (r rune) {
	if l.pos >= len(l.input) {
		l.width = 0
		return EOF
	}
	r, l.width = utf8.DecodeRuneInString(l.input[l.pos:])
	l.pos += l.width
	return r
}

// ignore skips over the pending input before this point
func (l *lexer) ignore() {
	l.start = l.pos
}

// backup steps back one rune
// can only be called ONCE per call to next()
func (l *lexer) backup() {
	l.pos -= l.width
}

// peek returns but does not consume
// the next rune in the input.
func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

// accept consumes the next rune
// if it's from the valid set
// and is a no-op if the rune is not in the valid set
func (l *lexer) accept(valid string) bool {
	if strings.IndexRune(valid, l.next()) >= 0 {
		return true
	}
	l.backup()
	return false
}

// acceptRun consumes a run of runes from the valid set
func (l *lexer) acceptRun(valid string) {
	for strings.IndexRune(valid, l.next()) >= 0 {
		// accept
	}
	l.backup()
}

func lexNumber(l *lexer) stateFn {
	// optional leading sign
	l.accept("+-")

	// is it hex?
	digits := "0123456789"
	if l.accept("0") && l.accept("xX") {
		digits = "0123456789abcdefABCDEF"
	}
	l.acceptRun(digits)
	if l.accept(".") {
		l.acceptRun(digits)
	}

	l.emit(itemNumber)
	// TODO scientific notation
	return startState
}

func lexRightMeta(l *lexer) stateFn {
    if l.accept(string(rightMeta)) {
        l.emit(itemRightBrace)
    } else {
        return l.errorf("expected } but received %s", l.peek())
    }

    switch r := l.next(); {
    case r == EOF:
        l.emit(itemEOF)
        return nil
    default:
        return l.errorf("expected EOF but received %s", string(r))
    }
}

func startState(l *lexer) stateFn {
	// accept leading whitespace
	for {
		switch r := l.next(); {
		case isSpace(r):
			l.ignore()
		case r == leftMeta:
            l.emit(itemLeftBrace)
			return lexInsideAction
        case r == EOF:
            l.emit(itemEOF)
            return nil
		default:
            return l.errorf("unexpected token: %s", string(r))
		}
	}
}

// error returns an error token and terminates the scan
// by passing back a nil pointer that will be the next state
// thereby terminating l.run
func (l *lexer) errorf(format string, args ...interface{}) stateFn {
	l.items <- item{
		itemError,
		fmt.Sprintf(format, args...),
	}
	return nil
}

func (l *lexer) nextItem() item {
	return <-l.items
	panic("unreachable")
}

func isSpace(r rune) bool {
	return len(strings.TrimSpace(string(r))) == 0

}

func isAlphaNumeric(r rune) bool {
	reg := regexp.MustCompile("[A-Za-z]")
	s := string(r)
	return reg.Match([]byte(s))
}

func main() {
    _, items := lex("testLex", "{}")
    for item := range items {
        log.Printf("Received %+v", item)
    }
}