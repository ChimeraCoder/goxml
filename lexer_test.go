package main

import (
	"log"
	"testing"
)

func Test_SimpleJSON(t *testing.T) {
	var items []item
	//_, itemsC := lex("testLex", `{"a":5, b : 'foo' }`)
	_, itemsC := lex("testLex", `{"a":5}`)
	for item := range itemsC {
		log.Print(item)
		if err := item.Err(); err != nil {
			t.Errorf("error: %s", err)
		}
		items = append(items, item)
	}
	expected := []item{
		item{itemLeftBrace, "{"},
		item{itemDoubleQuote, `"a"`},
		item{itemColon, ":"},
		item{itemNumber, "5"},
		item{itemRightBrace, "}"},
		item{itemEOF, string("")},
	}
	checkEqual(t, items, expected)
}

func checkEqual(t *testing.T, items, expected []item) {
	if len(items) != len(expected) {
		t.Errorf("Received %d tokens, expecting %d", len(items), len(expected))
		return
	}
	for i, item := range items {
		expectedItem := expected[i]
		if item.typ != expectedItem.typ || item.val != expectedItem.val {
			t.Errorf("Expected %+v and received %+v", expectedItem, item)
		}
	}
}
