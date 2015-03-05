%{
package main
import "fmt"
%}

%union {
  n int
  s string
}

%token NUM
%token DIGIT1to9
%token DIGIT
%token DIGITS
%token INT
%token FRAC
%token EXP
%token E
%token HEX_DIGIT 
%token NUMBER
%token UNESCAPEDCHAR
%token ESCAPEDCHAR
%token UNICODECHAR
%token CHAR 
%token CHARS
%token DBL_QUOTE


%%
input:    /* empty */
            | input line
;

line:     '\n'
           | exp '\n'      { fmt.Println($1.n); }
;

exp:     NUM           { $$.n = $1.n;        }
          | exp exp '+'   { $$.n = $1.n + $2.n; }
       | exp exp '-'   { $$.n = $1.n - $2.n; }
       | exp exp '*'   { $$.n = $1.n * $2.n; }
       | exp exp '/'   { $$.n = $1.n / $2.n; }
    /* Unary minus    */
       | exp 'n'       { $$.n = -$1.n;       }
;
%%
