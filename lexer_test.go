package main

import (
	"log"
	"reflect"
	"testing"
)

const SimpleXML = `<first>John</first>`

func Test_SimpleXML(t *testing.T) {
	var items []item
	//_, itemsC := lex("testLex", `{"a":5, b : 'foo' }`)
	_, results := lex("testLex", SimpleXML, nil)
	for result := range results {
		if err := result.item.Err(); err != nil {
			t.Errorf("error: %s", err)
		}
		items = append(items, result.item)
	}
	expected := []item{
		item{itemLeftAngleBracket, "<"},
		item{itemIdentifier, `first`},
		item{itemRightAngleBracket, ">"},
		item{itemIdentifier, "John"},
		item{itemLeftAngleBracket, "<"},
		item{itemForwardSlash, "/"},
		item{itemIdentifier, `first`},
		item{itemRightAngleBracket, ">"},
		item{itemEOF, ""},
	}
	checkEqual(t, items, expected)
}

func checkEqual(t *testing.T, items, expected []item) {
	if len(items) != len(expected) {
		t.Errorf("Received %d tokens, expecting %d: %+v", len(items), len(expected), items)
		return
	}
	for i, item := range items {
		expectedItem := expected[i]
		if item.typ != expectedItem.typ || item.val != expectedItem.val {
			t.Errorf("Expected %+v and received %+v", expectedItem, item)
			log.Print(expectedItem.typ)
			log.Print(item.typ)
			log.Print(reflect.TypeOf(expectedItem))
			log.Print(reflect.TypeOf(expectedItem))
		}
	}
}
