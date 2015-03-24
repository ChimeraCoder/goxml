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
%token  itemLeftBrace
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
%token  itemFunc
%token  itemLeftParen
%token  itemRightParen
%token  itemSemicolon
%token  itemReturn
%token  itemIncrement
%token  itemDecrement
%token  itemOperatorPlus
%token  itemVar
%token  itemAssignment


%%

/* 
   Datatypes required for parsing XML
   Many of these will also be useful for parsing arbitrary expressions
*/

XML : TAGS


TAGS : TAG
     | TAG TAGS
     ;

TAG :  itemLeftAngleBracket itemIdentifier ATTRIBUTES itemRightAngleBracket XML itemLeftAngleBracket itemForwardSlash itemIdentifier itemRightAngleBracket


ATTRIBUTES : /* empty */
           | ATTRIBUTE ATTRIBUTES
           ;

ATTRIBUTE  : itemIdentifier itemEqualSign itemDoubleQuote 
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
