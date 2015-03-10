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
%token  itemOperatorPlus
%token  itemVar
%token  itemAssignment


%%

TOPLEVEL    : JSON itemEOF /* allow just to ensure raw JSON tests parse */
                           /* since JSON can be empty, 
                           /* this also allows the empty program */
            | STATEMENTS itemEOF
            ;

STATEMENTS  : /* empty */
            | STATEMENT itemSemicolon STATEMENTS
            ;

STATEMENT   : itemVar itemAssignment EXPRESSION
            ; 

EXPRESSION  : itemFunc itemLeftParen FUNCARGS itemRightParen itemLeftBrace
            | JSON
            ;


FUNCARGS    : /* empty */
            | ELEMENTS
            ;


/* 
   Datatypes required for parsing JSON
   Many of these will also be useful for parsing arbitrary expressions
*/

JSON    : OBJECT {parsedAST = $$.val}
        | ARRAY 
        ;

OBJECT  : itemLeftBrace itemRightBrace /*{ $$.val = map[string]interface{}{}}*/
        | itemLeftBrace PAIRS itemRightBrace { $$.val = $2.mapval}
        ;

PAIRS   : PAIR  {$$.mapval = $1.mapval}/*{$$.mergeKeys(map[string]interface{}{$1.key : $1.val}); log.Printf("is %+v", $$.mapval)}*/
        | PAIR itemComma PAIRS {$$.mapval = mergeKeys($1.mapval, $3.mapval)}
        ;

PAIR    : KEY itemColon VALUE {$$.mapval = map[string]interface{}{$1.val.(string) : $3.val}}
        ;

KEY     : STRING
       | itemIdentifier
        ;

STRING  : itemSingleQuote {$$.val = strings.Trim($1.val.(string), "'") }
        | itemDoubleQuote {$$.val = strings.Trim($1.val.(string), "\"")}
        ;

VALUE   : STRING
        | OBJECT
        | ARRAY
        | itemIdentifier  {$$.val = parseIdentifier($1)}
        | itemNumber  { n, err := strconv.Atoi($1.val.(string)); if err != nil { yylex.Error(err.Error()) }; $$.val = n}
        ;

ARRAY   : itemLeftSquareBracket itemRightSquareBracket
        | itemLeftSquareBracket ELEMENTS itemRightSquareBracket { $$.val = $2.val.([]interface{}) }
        ;

ELEMENTS : VALUE {$$.val = []interface{}{$1.val}}
         | VALUE itemComma ELEMENTS {$$.val = append([]interface{}{$1.val}, $3.val.([]interface{})...) }
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
    //log.Printf("Lexed\t%s\t%d", item.val, typ)
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
