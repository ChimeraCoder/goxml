//line parser.y:2
package main

import __yyfmt__ "fmt"

//line parser.y:2
import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type itemType int

//line parser.y:21
type yySymType struct {
	yys int
	val interface{}

	key    string
	mapval map[string]interface{}

	scope symbolTable
}

const itemError = 57346
const itemLeftBrace = 57347
const itemLeftSquareBracket = 57348
const itemNumber = 57349
const itemRightBrace = 57350
const itemRightSquareBracket = 57351
const itemDoubleQuote = 57352
const itemSingleQuote = 57353
const itemIdentifier = 57354
const itemColon = 57355
const itemComma = 57356
const itemDot = 57357
const itemEOF = 57358
const itemFunc = 57359
const itemLeftParen = 57360
const itemRightParen = 57361
const itemSemicolon = 57362
const itemReturn = 57363
const itemIncrement = 57364
const itemDecrement = 57365
const itemOperatorPlus = 57366
const itemVar = 57367
const itemAssignment = 57368

var yyToknames = []string{
	"itemError",
	"itemLeftBrace",
	"itemLeftSquareBracket",
	"itemNumber",
	"itemRightBrace",
	"itemRightSquareBracket",
	"itemDoubleQuote",
	"itemSingleQuote",
	"itemIdentifier",
	"itemColon",
	"itemComma",
	"itemDot",
	"itemEOF",
	"itemFunc",
	"itemLeftParen",
	"itemRightParen",
	"itemSemicolon",
	"itemReturn",
	"itemIncrement",
	"itemDecrement",
	"itemOperatorPlus",
	"itemVar",
	"itemAssignment",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.y:139

var parsedAST interface{}

type yyLex struct {
	err error
}

func (jl *yyLex) Lex(lval *yySymType) int {
	result := <-results
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
		if jl.err != nil {
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

func main() {
	bts, err := ioutil.ReadAll((os.NewFile(0, "stdin")))
	if err != nil {
		log.Printf("%s", err)
	}
	parse(string(bts))
}

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 37
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 89

var yyAct = []int{

	3, 10, 9, 34, 44, 50, 15, 36, 20, 35,
	22, 26, 16, 17, 21, 2, 38, 37, 16, 17,
	41, 29, 13, 14, 38, 37, 43, 24, 18, 6,
	49, 47, 13, 14, 8, 46, 23, 19, 12, 6,
	48, 57, 11, 25, 8, 32, 31, 30, 45, 38,
	37, 38, 37, 54, 29, 55, 56, 53, 52, 51,
	13, 14, 28, 27, 13, 14, 40, 6, 33, 32,
	31, 39, 8, 32, 31, 30, 12, 13, 14, 40,
	11, 42, 32, 31, 39, 5, 4, 7, 1,
}
var yyPact = []int{

	55, -1000, -10, 12, -1000, -1000, -1000, 17, -8, -1000,
	-1000, 24, 27, 35, 59, -1000, -1000, -1000, -1000, 55,
	-1000, -4, 72, -22, -4, -1000, 40, 21, 18, -1000,
	-1000, -1000, -1000, -1000, 31, 16, -1000, -1000, -1000, -1000,
	-1000, -1000, -14, -1000, 27, -1000, 63, 72, -1000, 72,
	50, -4, -1000, -1000, -1000, 55, 33, -1000,
}
var yyPgo = []int{

	0, 88, 14, 0, 87, 86, 85, 81, 3, 2,
	1, 11, 63, 62, 9, 7,
}
var yyR1 = []int{

	0, 1, 1, 3, 3, 3, 4, 4, 4, 2,
	2, 2, 2, 2, 5, 7, 7, 6, 6, 9,
	9, 11, 11, 12, 13, 13, 15, 15, 14, 14,
	14, 14, 14, 10, 10, 8, 8,
}
var yyR2 = []int{

	0, 2, 2, 0, 3, 2, 4, 1, 2, 1,
	1, 1, 2, 2, 7, 0, 1, 1, 1, 2,
	3, 1, 3, 3, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 2, 3, 1, 3,
}
var yyChk = []int{

	-1000, -1, -2, -3, -5, -6, 12, -4, 17, -9,
	-10, 25, 21, 5, 6, 16, 22, 23, 16, 20,
	-3, -2, 18, 12, -2, 8, -11, -12, -13, -15,
	12, 11, 10, 9, -8, -14, -15, -9, -10, 12,
	7, -3, -7, -8, 26, 8, 14, 13, 9, 14,
	19, -2, -11, -14, -8, 5, -3, 8,
}
var yyDef = []int{

	3, -2, 7, 0, 9, 10, 11, 3, 0, 17,
	18, 0, 0, 0, 0, 1, 12, 13, 2, 3,
	5, 7, 15, 0, 8, 19, 0, 21, 0, 24,
	25, 26, 27, 33, 0, 35, 28, 29, 30, 31,
	32, 4, 0, 16, 0, 20, 0, 0, 34, 0,
	0, 6, 22, 23, 36, 3, 0, 14,
}
var yyTok1 = []int{

	1,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line parser.y:62
		{
			parsedAST = yyS[yypt-1].val.(func(symbolTable) map[string]interface{})(NewScope())
		}
	case 2:
		//line parser.y:66
		{
			parsedAST = yyVAL.val
		}
	case 8:
		//line parser.y:76
		{
			scope := NewScope()
			scope.Add("i", yySymType{val: 5.0})
			yyVAL.val = yyS[yypt-0].val.(func(symbolTable) float64)(scope)
		}
	case 10:
		//line parser.y:81
		{
			r1 := yyS[yypt-0].val
			yyVAL.val = func(st symbolTable) map[string]interface{} { return r1.(map[string]interface{}) }
		}
	case 12:
		//line parser.y:83
		{
			ident := yyS[yypt-1].val
			yyVAL.val = func(st symbolTable) float64 {
				return postfixOperation(itemIncrement, st.Lookup(ident.(string)).val, st, yylex)
			}
		}
	case 17:
		//line parser.y:101
		{
			yyVAL.val = yyS[yypt-0].val
		}
	case 20:
		//line parser.y:106
		{
			yyVAL.val = yyS[yypt-1].mapval
		}
	case 21:
		//line parser.y:109
		{
			yyVAL.mapval = yyS[yypt-0].mapval
		}
	case 22:
		//line parser.y:110
		{
			yyVAL.mapval = mergeKeys(yyS[yypt-2].mapval, yyS[yypt-0].mapval)
		}
	case 23:
		//line parser.y:113
		{
			yyVAL.mapval = map[string]interface{}{yyS[yypt-2].val.(string): yyS[yypt-0].val}
		}
	case 26:
		//line parser.y:120
		{
			yyVAL.val = strings.Trim(yyS[yypt-0].val.(string), "'")
		}
	case 27:
		//line parser.y:121
		{
			yyVAL.val = strings.Trim(yyS[yypt-0].val.(string), "\"")
		}
	case 31:
		//line parser.y:127
		{
			yyVAL.val = parseIdentifier(yyS[yypt-0])
		}
	case 32:
		//line parser.y:128
		{
			n, err := strconv.Atoi(yyS[yypt-0].val.(string))
			if err != nil {
				yylex.Error(err.Error())
			}
			yyVAL.val = n
		}
	case 34:
		//line parser.y:132
		{
			yyVAL.val = yyS[yypt-1].val.([]interface{})
		}
	case 35:
		//line parser.y:135
		{
			yyVAL.val = []interface{}{yyS[yypt-0].val}
		}
	case 36:
		//line parser.y:136
		{
			yyVAL.val = append([]interface{}{yyS[yypt-2].val}, yyS[yypt-0].val.([]interface{})...)
		}
	}
	goto yystack /* stack new state and value */
}
