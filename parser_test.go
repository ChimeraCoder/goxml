package main

import "testing"
import "log"

func Test_ParseSimpleJSON(t *testing.T) {
	log.Printf("parse simple")
	if err := parse(SimpleJSON); err != nil {
		t.Error(err)
	}
}

func Test_ParseNestedJSON(t *testing.T) {
	log.Printf("parse nested")
	if err := parse(NestedJSON); err != nil {
		t.Error(err)
	}
}
