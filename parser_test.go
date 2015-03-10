package main

import (
	"log"
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

	astm := ast.(map[string]interface{})
	if !reflect.DeepEqual(ast, expected) {
		t.Errorf("AST does not match: %+v, %+v", ast, expected)
		log.Printf("\n")
		log.Printf("%+v %+v", ast, expected)
		log.Printf("\nAST: ")
		printObj(ast.(map[string]interface{}))
		log.Printf("\nExpected: ")
		printObj(expected)
		log.Printf("\n")
		log.Printf("%t", reflect.DeepEqual(astm["a"], expected["a"]))
		log.Printf("%t", reflect.DeepEqual(astm["b"], expected["b"]))
		log.Printf("%t", reflect.DeepEqual(astm["cat"].(map[string]interface{})["dog"], expected["cat"].(map[string]interface{})["dog"]))
		log.Printf("%t %t", astm["cat"].(map[string]interface{})["dog"], expected["cat"].(map[string]interface{})["dog"])
	}
}
func Test_ParseJSFunction(t *testing.T) {
	if _, err := parse(NestedJSON); err != nil {
		t.Error(err)
	}
}

// convenience function for debugging
func printObj(o map[string]interface{}) {
	for key, val := range o {
		log.Printf("%+v (%s): %+v (%s)", reflect.ValueOf(key), reflect.TypeOf(key), reflect.ValueOf(val), reflect.TypeOf(val))
		if reflect.ValueOf(val).Kind() == reflect.Map {
			o2 := val.(map[string]interface{})
			for key2, val2 := range o2 {
				log.Printf("Nested: %+v (%s): %+v (%s)", reflect.ValueOf(key2), reflect.TypeOf(key2), reflect.ValueOf(val2), reflect.TypeOf(val2))
				if reflect.ValueOf(val2).Kind() == reflect.Slice {
					for _, item := range val2.([]interface{}) {
						log.Printf("Nested 2: %+v (%s)", reflect.ValueOf(item), reflect.TypeOf(item))
					}
				}
			}
		}
	}
}
