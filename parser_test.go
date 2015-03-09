package main

import "testing"


func Test_ParseSimpleJSON(t *testing.T) {
    if err := parse(SimpleJSON); err != nil{
        t.Error(err)
    }
}
