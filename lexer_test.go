package main

import "testing"

const SimpleJSON = `{"a":5}`
const NestedJSON = `{"a":5, b : 'bar', cat : { dog : true, elephant : ['hathi', 3]}}`

func Test_SimpleJSON(t *testing.T) {
	var items []item
	//_, itemsC := lex("testLex", `{"a":5, b : 'foo' }`)
	_, results := lex("testLex", SimpleJSON, nil)
	for result := range results {
		if err := result.item.Err(); err != nil {
			t.Errorf("error: %s", err)
		}
		items = append(items, result.item)
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

func Test_NestedJSON(t *testing.T) {
	var items []item
	_, results := lex("testLex", `{"a":5, b : 'bar', cat : { dog : true, elephant : ['hathi', 3]}`, nil)
	for result := range results {
		if err := result.item.Err(); err != nil {
			t.Errorf("error: %s", err)
		}
		items = append(items, result.item)
	}
	expected := []item{
		item{itemLeftBrace, "{"},
		item{itemDoubleQuote, `"a"`},
		item{itemColon, ":"},
		item{itemNumber, "5"},
		item{itemComma, ","},
		item{itemIdentifier, "b"},
		item{itemColon, ":"},
		item{itemSingleQuote, `'bar'`},
		item{itemComma, ","},
		item{itemIdentifier, "cat"},
		item{itemColon, ":"},
		item{itemLeftBrace, "{"},
		item{itemIdentifier, "dog"},
		item{itemColon, ":"},
		item{itemIdentifier, "true"},
		item{itemComma, ","},
		item{itemIdentifier, "elephant"},
		item{itemColon, ":"},
		item{itemLeftSquareBracket, "["},
		item{itemSingleQuote, `'hathi'`},
		item{itemComma, ","},
		item{itemNumber, "3"},
		item{itemRightSquareBracket, "]"},
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
