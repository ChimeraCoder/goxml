package main

import (
	"reflect"
	"testing"
)

func Test_ParseSimpleJSON(t *testing.T) {
	expected := map[string]interface{}{"a": 5}

	ast, err := parse(SimpleJSON)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(ast, expected) {
		t.Errorf("AST does not match: %+v, %+v", ast, expected)
	}
}

func Test_ParseNestedJSON(t *testing.T) {
	expected := map[string]interface{}{"a": 4, "b": "bar", "cat": map[string]interface{}{"dog": true, "elephant": []interface{}{"hathi", 3}}}

	ast, err := parse(NestedJSON)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(ast, expected) {
		t.Errorf("AST does not match: %+v, %+v", ast, expected)
	}
}

func Test_ParseJSFunction(t *testing.T) {
	ast, err := parse(NestedJSON)
	if err != nil {
		t.Error(err)
	}

	if reflect.TypeOf(ast).Kind() != reflect.Func {
		t.Errorf("Expected Func and received Kind %s", reflect.TypeOf(ast).Kind())
	}
}
