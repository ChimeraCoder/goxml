%{
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
)

type itemType int


%}


// This is the bases of {yy}SymType

%union{
    val interface{}

    key string
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
        | OBJECT itemEOF {fmt.Printf("Parsed: %+v\n", $$.val)}
        | ARRAY itemEOF
        ;

OBJECT  : itemLeftBrace itemRightBrace /*{ $$.val = map[string]interface{}{}}*/
        | itemLeftBrace PAIRS itemRightBrace { $$.val = map[string]interface{}{$2.key : $2.val}}
        ;

PAIRS   : PAIR  {log.Printf("d %+v", $$); $$.key = $1.key; $$.val = $1.val}
        | PAIR itemComma PAIRS
        ;

PAIR    : KEY itemColon VALUE {$$.val = fmt.Sprintf("%s : %v", $1.val, $3.val); $$.key = $1.val.(string); $$.val = $3.val; log.Println($$.val)}
        ;

KEY     : STRING 
        | itemIdentifier
        ;

STRING  : itemSingleQuote
        | itemDoubleQuote
        ;

VALUE   : STRING
        | OBJECT
        | ARRAY
        | itemIdentifier  /* TODO specify keywords true/false/etc. */
        | itemNumber
        ;

ARRAY   : itemLeftSquareBracket itemRightSquareBracket
        | itemLeftSquareBracket ELEMENTS itemRightSquareBracket
        ;

ELEMENTS : VALUE
         | VALUE itemComma ELEMENTS
         ;

%%


type yyLex struct {
    err error
}

func (jl *yyLex) Lex(lval *yySymType) int {
    result := <- results
    item := result.item


    if item.typ == itemError || item.typ == itemEOF {
        done = true
    }

    lval.val = item.val
    typ := int(item.typ)
    log.Printf("Lexed\t%s\t%d", item.val, typ)
    return typ
}

func (jl *yyLex) Error(e string) {
    jl.err = fmt.Errorf("%s", e)
    log.Printf("Parsing error: %s", e)
}


// The actual lexer
var l *lexer
var results chan result

var done bool

func parse(input string) error {


    // Set up the lexer, which will run concurrently
    l, results = lex("testLex", input, nil)

    for {
        jl := &yyLex{}
        yyParse(jl)
        if jl.err != nil{
            return jl.err
        }
        if done {
            // reset for testing package
            done = false
            break
        }

    }
    return nil
}


func main(){
    bts, err := ioutil.ReadAll((os.NewFile(0, "stdin")))
    if err != nil{
        log.Printf("%s", err)
    }
    parse(string(bts))
}
