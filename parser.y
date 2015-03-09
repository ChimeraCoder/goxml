%{
package main

import (
    "io/ioutil"
    "log"
    "os"
)

type itemType int


%}


// This is the bases of {yy}SymType

%union{
    val interface{}
}


// We don't need %token declarations
// as they are contained in the lexer
// as constants of type itemType

%token itemError 
%token itemLeftBrace
%token	itemLeftSquareBracket
%token	itemNumber
%token	itemRightBrace
%token	itemRightSquareBracket
%token	itemDoubleQuote
%token	itemSingleQuote
%token	itemIdentifier
%token	itemColon
%token	itemComma
%token	itemDot
%token	itemEOF



%%

json    : /* empty */
        | itemLeftBrace itemRightBrace
        | json itemEOF
        ;


%%


type yyLex struct {
}

func (jl *yyLex) Lex(lval *yySymType) int {
    result := <- results
    item := result.item


    if item.typ == itemError || item.typ == itemEOF {
        done = true
    }

    lval.val = item.val
    typ := int(item.typ)
    return typ
}

func (jl *yyLex) Error(e string) {
    log.Printf("Parsing error: %s", e)
}


// The actual lexer
var l *lexer
var results chan result

var done bool

func main() {


    // Set up the lexer, which will run concurrently
    bts, err := ioutil.ReadAll((os.NewFile(0, "stdin")))
    if err != nil{
        log.Fatal(err)
    }
    l, results = lex("testLex", string(bts), nil)

    for !done {
        yyParse(&yyLex{})
    }
}
