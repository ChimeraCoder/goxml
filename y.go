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
const itemOperatorPlus = 57365
const itemVar = 57366
const itemAssignment = 57367

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
	"itemOperatorPlus",
	"itemVar",
	"itemAssignment",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.y:129
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

const yyNprod = 32
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 73

var yyAct = []int{

	3, 5, 4, 25, 32, 10, 15, 27, 26, 47,
	29, 28, 17, 7, 8, 20, 34, 7, 8, 33,
	14, 7, 8, 31, 40, 12, 23, 22, 30, 12,
	13, 37, 9, 38, 39, 29, 28, 43, 41, 50,
	29, 28, 29, 28, 46, 20, 36, 45, 48, 49,
	44, 7, 8, 31, 19, 24, 23, 22, 30, 16,
	18, 23, 22, 21, 23, 22, 21, 35, 2, 42,
	11, 6, 1,
}
var yyPact = []int{

	8, -1000, 14, 4, -1000, -1000, -14, 51, 46, -21,
	-1000, -1000, 1, -1000, -1000, 8, -1000, 38, 17, 20,
	-1000, -1000, -1000, -1000, -1000, 25, 10, -1000, -1000, -1000,
	-1000, -1000, 12, 16, -1000, -1000, -1000, 54, 16, -1000,
	16, -1000, -10, -1000, -1000, -1000, -1000, 43, 8, 31,
	-1000,
}
var yyPgo = []int{

	0, 72, 67, 0, 71, 5, 70, 69, 3, 2,
	1, 12, 60, 54, 8, 7,
}
var yyR1 = []int{

	0, 1, 1, 3, 3, 4, 4, 5, 5, 6,
	7, 7, 2, 2, 9, 9, 11, 11, 12, 13,
	13, 15, 15, 14, 14, 14, 14, 14, 10, 10,
	8, 8,
}
var yyR2 = []int{

	0, 2, 2, 0, 3, 3, 1, 1, 1, 7,
	0, 1, 1, 1, 2, 3, 1, 3, 3, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 2, 3,
	1, 3,
}
var yyChk = []int{

	-1000, -1, -2, -3, -9, -10, -4, 5, 6, 24,
	-5, -6, 17, 16, 16, 20, 8, -11, -12, -13,
	-15, 12, 11, 10, 9, -8, -14, -15, -9, -10,
	12, 7, 25, 18, -3, -2, 8, 14, 13, 9,
	14, -5, -7, -8, -11, -14, -8, 19, 5, -3,
	8,
}
var yyDef = []int{

	3, -2, 8, 0, 12, 13, 0, 0, 0, 0,
	6, 7, 0, 1, 2, 3, 14, 0, 16, 0,
	19, 20, 21, 22, 28, 0, 30, 23, 24, 25,
	26, 27, 0, 10, 4, 8, 15, 0, 0, 29,
	0, 5, 0, 11, 17, 18, 31, 0, 3, 0,
	9,
}
var yyTok1 = []int{

	1,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25,
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

	case 9:
		//line parser.y:77
		{
			log.Print("ahsasdf")
		}
	case 12:
		//line parser.y:91
		{
			parsedAST = yyVAL.val
		}
	case 15:
		//line parser.y:96
		{
			yyVAL.val = yyS[yypt-1].mapval
		}
	case 16:
		//line parser.y:99
		{
			yyVAL.mapval = yyS[yypt-0].mapval
		}
	case 17:
		//line parser.y:100
		{
			yyVAL.mapval = mergeKeys(yyS[yypt-2].mapval, yyS[yypt-0].mapval)
		}
	case 18:
		//line parser.y:103
		{
			yyVAL.mapval = map[string]interface{}{yyS[yypt-2].val.(string): yyS[yypt-0].val}
		}
	case 21:
		//line parser.y:110
		{
			yyVAL.val = strings.Trim(yyS[yypt-0].val.(string), "'")
		}
	case 22:
		//line parser.y:111
		{
			yyVAL.val = strings.Trim(yyS[yypt-0].val.(string), "\"")
		}
	case 26:
		//line parser.y:117
		{
			yyVAL.val = parseIdentifier(yyS[yypt-0])
		}
	case 27:
		//line parser.y:118
		{
			n, err := strconv.Atoi(yyS[yypt-0].val.(string))
			if err != nil {
				yylex.Error(err.Error())
			}
			yyVAL.val = n
		}
	case 29:
		//line parser.y:122
		{
			yyVAL.val = yyS[yypt-1].val.([]interface{})
		}
	case 30:
		//line parser.y:125
		{
			yyVAL.val = []interface{}{yyS[yypt-0].val}
		}
	case 31:
		//line parser.y:126
		{
			yyVAL.val = append([]interface{}{yyS[yypt-2].val}, yyS[yypt-0].val.([]interface{})...)
		}
	}
	goto yystack /* stack new state and value */
}
