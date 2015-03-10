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

	key string
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
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.y:95
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

const yyNprod = 22
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 45

var yyAct = []int{

	17, 19, 9, 18, 7, 6, 12, 4, 5, 23,
	28, 16, 15, 14, 22, 4, 5, 23, 25, 26,
	15, 14, 22, 15, 14, 13, 27, 12, 29, 31,
	30, 8, 24, 15, 14, 13, 4, 5, 21, 3,
	20, 2, 11, 10, 1,
}
var yyPact = []int{

	31, -1000, -11, -12, 23, 2, -1000, -1000, -1000, 24,
	4, 6, -1000, -1000, -1000, -1000, -1000, 17, -4, -1000,
	-1000, -1000, -1000, -1000, -1000, 13, 10, -1000, 10, -1000,
	-1000, -1000,
}
var yyPgo = []int{

	0, 44, 40, 38, 2, 43, 42, 3, 1, 0,
}
var yyR1 = []int{

	0, 1, 1, 1, 2, 2, 4, 4, 5, 6,
	6, 8, 8, 7, 7, 7, 7, 7, 3, 3,
	9, 9,
}
var yyR2 = []int{

	0, 0, 2, 2, 2, 3, 1, 3, 3, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 2, 3,
	1, 3,
}
var yyChk = []int{

	-1000, -1, -2, -3, 5, 6, 16, 16, 8, -4,
	-5, -6, -8, 12, 11, 10, 9, -9, -7, -8,
	-2, -3, 12, 7, 8, 14, 13, 9, 14, -4,
	-7, -9,
}
var yyDef = []int{

	1, -2, 0, 0, 0, 0, 2, 3, 4, 0,
	6, 0, 9, 10, 11, 12, 18, 0, 20, 13,
	14, 15, 16, 17, 5, 0, 0, 19, 0, 7,
	8, 21,
}
var yyTok1 = []int{

	1,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22,
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

	case 2:
		//line parser.y:57
		{
			parsedAST = yyVAL.val
		}
	case 5:
		//line parser.y:62
		{
			yyVAL.val = map[string]interface{}{yyS[yypt-1].key: yyS[yypt-1].val}
		}
	case 6:
		//line parser.y:65
		{
			yyVAL.key = yyS[yypt-0].key
			yyVAL.val = yyS[yypt-0].val
		}
	case 8:
		//line parser.y:69
		{
			yyVAL.val = fmt.Sprintf("%s : %v", yyS[yypt-2].val, yyS[yypt-0].val)
			yyVAL.key = yyS[yypt-2].val.(string)
			yyVAL.val = yyS[yypt-0].val
		}
	case 11:
		//line parser.y:76
		{
			yyVAL.val = strings.Trim(yyS[yypt-0].val.(string), "'")
		}
	case 12:
		//line parser.y:77
		{
			yyVAL.val = strings.Trim(yyS[yypt-0].val.(string), "\"")
		}
	case 17:
		//line parser.y:84
		{
			n, err := strconv.Atoi(yyS[yypt-0].val.(string))
			if err != nil {
				yylex.Error(err.Error())
			}
			yyVAL.val = n
		}
	case 19:
		//line parser.y:88
		{
			yyVAL.val = yyS[yypt-1].val.([]interface{})
		}
	case 20:
		//line parser.y:91
		{
			yyVAL.val = []interface{}{yyS[yypt-0].val}
		}
	case 21:
		//line parser.y:92
		{
			yyVAL.val = append(yyS[yypt-0].val.([]interface{}), yyS[yypt-2].val)
		}
	}
	goto yystack /* stack new state and value */
}
