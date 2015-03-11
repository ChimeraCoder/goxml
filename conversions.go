package main

import (
	"fmt"
	"log"
	"reflect"
)

// mergeKeys will merge the mapval fields, overwriting y.mapval
// it is safe to call even if y.mapval is nil
func (y *yySymType) mergeKeys(other map[string]interface{}) {
	map1 := y.mapval
	y.mapval = mergeKeys(map1, other)
}

// mergeKeys will produce the union of two sets
// The behavior for duplicate keys is undefined
func mergeKeys(a, b map[string]interface{}) map[string]interface{} {
	result := map[string]interface{}{}
	for _, m := range []map[string]interface{}{a, b} {
		for k, val := range m {
			result[k] = val
		}
	}
	return result
}

func parseIdentifier(y yySymType) interface{} {
	switch y.val {
	case "true":
		return true
	case "false":
		return false
	case "null":
		return nil
	default:
		return y.val
	}
	return y
}

type Symbol struct {
	name string
	val  yySymType
}

type symbolTable struct {
	// heap-allocated and therefore mutable
	symbols map[string]Symbol
}

func NewScope() symbolTable {
	return symbolTable{map[string]Symbol{}}
}

func (st *symbolTable) Lookup(name string) *Symbol {
	result, ok := st.symbols[name]
	if !ok {
		return nil
	}
	return &result
}

func (st symbolTable) Add(name string, value yySymType) {
	st.symbols[name] = Symbol{name, value}
}

// Nest is used for block scoping
func (st symbolTable) Nest() symbolTable {
	other := symbolTable{map[string]Symbol{}}
	for k, v := range st.symbols {
		other.symbols[k] = v
	}
	return other
}

func (st symbolTable) MergeInto(outerScope symbolTable) {
	// TODO account for shadowing properly
	for k, _ := range outerScope.symbols {
		outerScope.symbols[k] = st.symbols[k]
	}
}

type argument struct {
	name  string
	value Symbol
}

func makeFunc(args []argument, initialScope symbolTable) {
	localScope := initialScope.Nest()
	defer localScope.MergeInto(initialScope)
}

// postfixOperation runs a postfix operator on a number (float64)
// the only postfix operations are numeric
func postfixOperation(operator int, value yySymType, scope symbolTable, yylex yyLexer) (result float64) {
	log.Printf("Scope: %+v", scope)
	var f float64
	var ok bool

	val := value.val

	switch v := value.val.(type) {
	case string:
		s := scope.Lookup(v)
		if s == nil {
			yylex.Error(fmt.Sprintf("Variable %s not in scope", v))
			return
		}
		val = s.val
	case float64:
		val = v

	default:
		yylex.Error(fmt.Sprintf("Expected string and received %s %+v", reflect.TypeOf(v), v))
		return
	}

	f, ok = val.(float64)
	if !ok {
		yylex.Error(fmt.Sprintf("Expected float64 and received %s %+v", reflect.TypeOf(val), val))
		return
	}

	switch operator {
	case itemIncrement:
		result = f + 1
	case itemDecrement:
		result = f - 1
	default:
		yylex.Error("asfda")
		result = 0
	}
	log.Printf("Result %f", result)
	return
}

type postfixOperator struct {
	operation interface{} // This will have Kind() == Func
	args      []argument
}
