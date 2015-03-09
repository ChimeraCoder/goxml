package main

import (
	"reflect"
	"testing"
)

func Test_ParseSimpleJSON(t *testing.T) {
	ast, err := parse(SimpleJSON)
	if err != nil {
		t.Error(err)
	}
	expected := map[string]interface{}{"a": 5}
	if !reflect.DeepEqual(ast, expected) {
		t.Errorf("AST does not match: %+v, %+v", ast, expected)
	}
}

func Test_ParseNestedJSON(t *testing.T) {
	if _, err := parse(NestedJSON); err != nil {
		t.Error(err)
	}
}
