package main

import (
	"log"
	"reflect"
	"testing"
)

const SimpleXML = `<first>John</first>`
const NestedXML = `<person id="13">
    <name>
        <first>John</first>
        <last>Doe</last>
    </name>
`

func Test_LexSimpleXML(t *testing.T) {
	var items []item
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

func Test_LexNestedXML(t *testing.T) {
	var items []item
	_, results := lex("testLex", NestedXML, nil)
	for result := range results {
		if err := result.item.Err(); err != nil {
			t.Errorf("error: %s", err)
		}
		items = append(items, result.item)
	}
	expected := []item{

		item{itemLeftAngleBracket, "<"},
		item{itemIdentifier, `person`},
		item{itemIdentifier, `id`},
		item{itemEqualSign, `=`},
		item{itemDoubleQuote, `"13"`},
		item{itemRightAngleBracket, ">"},

		item{itemLeftAngleBracket, "<"},
		item{itemIdentifier, `name`},
		item{itemRightAngleBracket, ">"},

		item{itemLeftAngleBracket, "<"},
		item{itemIdentifier, `first`},
		item{itemRightAngleBracket, ">"},
		item{itemIdentifier, "John"},
		item{itemLeftAngleBracket, "<"},
		item{itemForwardSlash, "/"},
		item{itemIdentifier, `first`},
		item{itemRightAngleBracket, ">"},

		item{itemLeftAngleBracket, "<"},
		item{itemIdentifier, `last`},
		item{itemRightAngleBracket, ">"},
		item{itemIdentifier, "Doe"},
		item{itemLeftAngleBracket, "<"},
		item{itemForwardSlash, "/"},
		item{itemIdentifier, `last`},
		item{itemRightAngleBracket, ">"},

		item{itemLeftAngleBracket, "<"},
		item{itemForwardSlash, "/"},
		item{itemIdentifier, `name`},
		item{itemRightAngleBracket, ">"},
		item{itemEOF, ""},
	}
	checkEqual(t, items, expected)
}

func checkEqual(t *testing.T, items, expected []item) {
	if len(items) != len(expected) {
		t.Fatalf("Received %d tokens, expecting %d:\n%+v\n%+v", len(items), len(expected), items, expected)
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
