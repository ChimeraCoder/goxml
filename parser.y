%{
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "strconv"
    "strings"
)

type itemType int


%}


// This is the bases of {yy}SymType

%union{
    val interface{}

    key string
    mapval map[string]interface{}

}


// We don't need %token declarations
// as they are contained in the lexer
// as constants of type itemType

%token  itemError 
%token	itemLeftSquareBracket
%token	itemNumber
%token	itemDoubleQuote
%token	itemSingleQuote
%token	itemIdentifier
%token	itemEOF
%token  itemLeftAngleBracket
%token  itemRightAngleBracket
%token  itemForwardSlash
%token  itemEqualSign


%%

/* 
   Datatypes required for parsing XML
   Many of these will also be useful for parsing arbitrary expressions
*/

XML : OPEN ATTRIBUTES CONTENT CLOSE
    ;


OPEN : itemLeftAngleBracket
     ;

CLOSE : itemRightAngleBracket
      ;


CONTENT : CONTENT XML
        | /* empty */
        ;


ATTRIBUTES : ATTRIBUTES ATTRIBUTE
           | /* empty */
           ;


ATTRIBUTE : NAME
          | NAME itemEqualSign VALUE


NAME : itemIdentifier
     ;

VALUE : itemIdentifier
      ;
%%

var parsedAST interface{}

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

func parse(input string) (interface{}, error) {
    parsedAST = nil

    // Set up the lexer, which will run concurrently
    l, results = lex("testLex", input, nil)

    for {
        jl := &yyLex{}
        yyParse(jl)
        if jl.err != nil{
            return parsedAST, jl.err
        }
        if done {
            // reset for testing package
            done = false
            break
        }

    }
    ast := parsedAST
    parsedAST = nil
    return ast, nil
}


func main(){
    bts, err := ioutil.ReadAll((os.NewFile(0, "stdin")))
    if err != nil{
        log.Printf("%s", err)
    }
    parse(string(bts))
}
