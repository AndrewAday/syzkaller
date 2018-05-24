//line scanner/strace.y:2
package scanner

import __yyfmt__ "fmt"

//line scanner/strace.y:2
import (
	types "github.com/google/syzkaller/tools/moonshine/strace_types"
)

//line scanner/strace.y:11
type StraceSymType struct {
	yys               int
	data              string
	val_int           int64
	val_double        float64
	val_uint          uint64
	val_field         *types.Field
	val_call          *types.Call
	val_macro         *types.Macro
	val_int_type      *types.IntType
	val_identifiers   []*types.BufferType
	val_buf_type      *types.BufferType
	val_struct_type   *types.StructType
	val_dynamic_type  *types.DynamicType
	val_array_type    *types.ArrayType
	val_pointer_type  *types.PointerType
	val_flag_type     *types.FlagType
	val_expr_type     *types.Expression
	val_type          types.Type
	val_ip_type       *types.IpType
	val_types         []types.Type
	val_parenthetical *types.Parenthetical
	val_syscall       *types.Syscall
}

const STRING_LITERAL = 57346
const IPV4 = 57347
const IPV6 = 57348
const IDENTIFIER = 57349
const FLAG = 57350
const DATETIME = 57351
const SIGNAL_PLUS = 57352
const SIGNAL_MINUS = 57353
const MAC = 57354
const INT = 57355
const UINT = 57356
const DOUBLE = 57357
const QUESTION = 57358
const ARROW = 57359
const OR = 57360
const AND = 57361
const LOR = 57362
const TIMES = 57363
const LAND = 57364
const LEQUAL = 57365
const ONESCOMP = 57366
const LSHIFT = 57367
const RSHIFT = 57368
const NOT = 57369
const COMMA = 57370
const LBRACKET = 57371
const RBRACKET = 57372
const LBRACKET_SQUARE = 57373
const RBRACKET_SQUARE = 57374
const LPAREN = 57375
const RPAREN = 57376
const EQUALS = 57377
const UNFINISHED = 57378
const RESUMED = 57379
const NULL = 57380
const AT = 57381
const COLON = 57382
const NOTYPE = 57383
const NOFLAG = 57384

var StraceToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"STRING_LITERAL",
	"IPV4",
	"IPV6",
	"IDENTIFIER",
	"FLAG",
	"DATETIME",
	"SIGNAL_PLUS",
	"SIGNAL_MINUS",
	"MAC",
	"INT",
	"UINT",
	"DOUBLE",
	"QUESTION",
	"ARROW",
	"OR",
	"AND",
	"LOR",
	"TIMES",
	"LAND",
	"LEQUAL",
	"ONESCOMP",
	"LSHIFT",
	"RSHIFT",
	"NOT",
	"COMMA",
	"LBRACKET",
	"RBRACKET",
	"LBRACKET_SQUARE",
	"RBRACKET_SQUARE",
	"LPAREN",
	"RPAREN",
	"EQUALS",
	"UNFINISHED",
	"RESUMED",
	"NULL",
	"AT",
	"COLON",
	"NOTYPE",
	"NOFLAG",
}
var StraceStatenames = [...]string{}

const StraceEofCode = 1
const StraceErrCode = 2
const StraceInitialStackSize = 16

//line yacctab:1
var StraceExca = [...]int{
	-1, 1,
	1, -1,
	-2, 0,
}

const StracePrivate = 57344

const StraceLast = 561

var StraceAct = [...]int{

	22, 9, 12, 163, 164, 144, 129, 3, 128, 14,
	46, 122, 47, 56, 170, 176, 44, 48, 8, 38,
	35, 36, 82, 180, 81, 166, 167, 89, 85, 74,
	76, 84, 168, 169, 15, 165, 30, 4, 175, 44,
	42, 187, 44, 24, 79, 142, 90, 91, 137, 93,
	95, 179, 96, 97, 98, 99, 100, 101, 102, 159,
	92, 158, 105, 106, 107, 108, 109, 110, 111, 112,
	113, 114, 115, 116, 117, 118, 119, 156, 145, 143,
	19, 31, 32, 21, 34, 20, 126, 44, 33, 35,
	36, 154, 34, 43, 125, 28, 138, 35, 36, 120,
	27, 83, 157, 177, 46, 30, 47, 23, 27, 26,
	40, 48, 37, 39, 29, 75, 155, 26, 19, 31,
	32, 21, 34, 20, 140, 90, 33, 35, 36, 104,
	41, 162, 153, 28, 152, 133, 132, 44, 27, 80,
	5, 103, 139, 30, 123, 23, 124, 26, 146, 147,
	45, 148, 29, 94, 80, 131, 173, 149, 173, 134,
	135, 178, 136, 173, 173, 172, 181, 172, 183, 77,
	130, 2, 172, 172, 86, 87, 78, 88, 173, 173,
	161, 184, 173, 183, 160, 173, 183, 172, 172, 183,
	171, 172, 171, 151, 172, 150, 104, 171, 171, 174,
	121, 174, 145, 18, 13, 25, 174, 174, 17, 16,
	10, 11, 171, 171, 1, 0, 171, 0, 0, 171,
	0, 174, 174, 0, 0, 174, 0, 0, 174, 19,
	31, 32, 21, 34, 20, 0, 0, 33, 35, 36,
	0, 0, 0, 0, 28, 0, 0, 0, 0, 27,
	0, 0, 0, 0, 30, 0, 23, 0, 26, 7,
	0, 6, 0, 29, 19, 31, 32, 21, 34, 20,
	0, 0, 33, 35, 36, 0, 0, 0, 0, 28,
	0, 0, 0, 0, 27, 0, 0, 0, 0, 30,
	0, 23, 57, 26, 0, 0, 0, 0, 29, 19,
	31, 32, 21, 34, 20, 0, 0, 33, 35, 36,
	0, 0, 0, 0, 28, 0, 0, 0, 0, 27,
	0, 0, 0, 0, 30, 141, 23, 0, 26, 0,
	0, 0, 0, 29, 19, 31, 32, 21, 58, 20,
	0, 0, 33, 35, 36, 0, 0, 0, 0, 28,
	0, 0, 0, 0, 27, 0, 0, 0, 0, 30,
	0, 23, 57, 26, 0, 0, 0, 0, 29, 19,
	31, 32, 21, 34, 20, 0, 0, 33, 35, 36,
	0, 0, 0, 0, 28, 0, 0, 0, 0, 27,
	0, 0, 0, 0, 30, 0, 23, 0, 26, 0,
	0, 0, 0, 29, 19, 31, 32, 127, 34, 20,
	0, 0, 33, 35, 36, 0, 0, 0, 0, 28,
	0, 0, 0, 0, 27, 0, 0, 0, 0, 30,
	0, 23, 0, 26, 170, 176, 0, 0, 29, 0,
	35, 36, 0, 0, 0, 166, 167, 0, 0, 0,
	0, 0, 168, 169, 0, 165, 30, 0, 175, 170,
	176, 186, 0, 0, 0, 35, 36, 0, 0, 0,
	166, 167, 0, 0, 0, 0, 0, 168, 169, 0,
	165, 30, 0, 175, 170, 176, 185, 0, 0, 0,
	35, 36, 0, 0, 0, 166, 167, 0, 0, 0,
	0, 0, 168, 169, 0, 165, 30, 0, 175, 170,
	176, 182, 0, 0, 0, 35, 36, 0, 0, 0,
	166, 167, 0, 0, 0, 0, 0, 168, 169, 0,
	165, 30, 0, 175, 59, 63, 61, 66, 60, 62,
	0, 64, 65, 67, 71, 69, 0, 68, 70, 0,
	72, 73, 49, 53, 51, 0, 50, 52, 0, 54,
	55,
}
var StracePact = [...]int{

	158, -1000, 0, 107, 225, 76, 96, 5, 59, -1000,
	-1000, -1000, 133, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -23, 534, 330, 516, 525, 84, 84, 162, -1000,
	365, -1000, -1000, -1000, 106, -1000, -1000, -1000, -12, 67,
	-4, -7, 161, -8, 365, 365, 365, 114, 365, 84,
	84, 84, 84, 84, 84, 84, 109, -1000, 121, 84,
	84, 84, 84, 84, 84, 84, 84, 84, 84, 84,
	84, 84, 84, 84, 65, 192, -1000, -1000, -24, 116,
	400, -1000, -27, -29, 157, 139, 103, 102, -1000, 146,
	-1000, -1000, 14, -1000, 365, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 110, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, 188, 365, 295, -1000, 11, 45, 71, 135, 144,
	-1000, -1000, 187, 185, 101, 99, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 195, 83, 69, -1000, -1000,
	27, 25, 176, 172, 98, 502, 70, 502, -1000, -1000,
	17, -11, 502, 477, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, 260, -1000, 502, 452, -1000,
	-1000, 427, -1000, -1000, 7, -1000, -1000, -1000,
}
var StracePgo = [...]int{

	0, 214, 211, 5, 43, 210, 34, 209, 9, 0,
	2, 208, 4, 3, 205, 1, 204, 203, 13,
}
var StraceR1 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 13, 13, 12, 12, 12, 12, 12,
	12, 12, 12, 12, 12, 18, 18, 15, 15, 15,
	15, 15, 15, 15, 15, 15, 7, 11, 14, 14,
	16, 16, 16, 8, 8, 6, 6, 2, 2, 2,
	2, 5, 5, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 10, 10, 10, 10, 10,
	10, 4, 4, 9, 17, 17, 17, 3, 3,
}
var StraceR2 = [...]int{

	0, 4, 5, 6, 7, 5, 5, 5, 8, 8,
	6, 6, 6, 9, 9, 6, 7, 7, 7, 11,
	11, 10, 10, 1, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 3, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 3, 4, 4, 4,
	2, 4, 1, 3, 2, 3, 4, 2, 3, 3,
	4, 1, 1, 1, 4, 1, 1, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	2, 1, 1, 1, 1, 1, 1, 1, 2,
}
var StraceChk = [...]int{

	-1000, -1, 13, 7, 37, 33, 36, 34, -18, -15,
	-5, -2, -10, -16, -8, -6, -7, -11, -17, 4,
	9, 7, -9, 31, -4, -14, 33, 24, 19, 38,
	29, 5, 6, 12, 8, 13, 14, 36, -18, 37,
	34, 34, 35, 34, 28, 17, 33, 35, 40, 18,
	22, 20, 23, 19, 25, 26, -18, 32, 8, 18,
	22, 20, 23, 19, 25, 26, 21, 18, 22, 20,
	23, 19, 25, 26, -10, 31, -10, 7, 14, -18,
	33, 36, 34, 34, 35, 35, 13, 14, 16, 35,
	-15, -15, -18, -15, 39, -15, -10, -10, -10, -10,
	-10, -10, -10, 32, 8, -10, -10, -10, -10, -10,
	-10, -10, -10, -10, -10, -10, -10, -10, -10, -10,
	34, 8, 35, 28, 30, -18, -3, 7, 35, 35,
	13, 16, 33, 33, 13, 14, 16, 34, -15, 32,
	-15, 30, 34, 34, -3, 7, 13, 14, 16, 13,
	8, 8, 33, 33, 8, 33, 8, 33, 34, 34,
	8, 8, 33, -13, -12, 28, 18, 19, 25, 26,
	7, -6, -8, -9, -4, 31, 8, 33, -13, 34,
	34, -13, 34, -12, -13, 34, 34, 34,
}
var StraceDef = [...]int{

	0, -2, 0, 0, 0, 0, 0, 0, 0, 35,
	37, 38, 39, 40, 41, 42, 43, 44, 45, 61,
	62, 0, 63, 0, 65, 66, 0, 0, 0, 52,
	0, 94, 95, 96, 93, 91, 92, 1, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 57, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 54, 93, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 90, 50, 0, 0,
	0, 2, 0, 0, 0, 0, 5, 6, 7, 0,
	36, 46, 0, 58, 0, 59, 69, 72, 75, 78,
	81, 84, 87, 53, 0, 70, 73, 76, 79, 82,
	85, 88, 89, 68, 71, 74, 77, 80, 83, 86,
	67, 0, 0, 0, 55, 0, 0, 97, 0, 0,
	15, 3, 0, 0, 10, 11, 12, 47, 60, 64,
	51, 56, 48, 49, 98, 97, 16, 17, 18, 4,
	0, 0, 0, 0, 0, 0, 0, 0, 8, 9,
	0, 0, 0, 0, 23, 25, 26, 27, 28, 29,
	30, 31, 32, 33, 34, 0, 93, 0, 0, 13,
	14, 0, 21, 24, 0, 22, 19, 20,
}
var StraceTok1 = [...]int{

	1,
}
var StraceTok2 = [...]int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42,
}
var StraceTok3 = [...]int{
	0,
}

var StraceErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	StraceDebug        = 0
	StraceErrorVerbose = false
)

type StraceLexer interface {
	Lex(lval *StraceSymType) int
	Error(s string)
}

type StraceParser interface {
	Parse(StraceLexer) int
	Lookahead() int
}

type StraceParserImpl struct {
	lval  StraceSymType
	stack [StraceInitialStackSize]StraceSymType
	char  int
}

func (p *StraceParserImpl) Lookahead() int {
	return p.char
}

func StraceNewParser() StraceParser {
	return &StraceParserImpl{}
}

const StraceFlag = -1000

func StraceTokname(c int) string {
	if c >= 1 && c-1 < len(StraceToknames) {
		if StraceToknames[c-1] != "" {
			return StraceToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func StraceStatname(s int) string {
	if s >= 0 && s < len(StraceStatenames) {
		if StraceStatenames[s] != "" {
			return StraceStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func StraceErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !StraceErrorVerbose {
		return "syntax error"
	}

	for _, e := range StraceErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + StraceTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := StracePact[state]
	for tok := TOKSTART; tok-1 < len(StraceToknames); tok++ {
		if n := base + tok; n >= 0 && n < StraceLast && StraceChk[StraceAct[n]] == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if StraceDef[state] == -2 {
		i := 0
		for StraceExca[i] != -1 || StraceExca[i+1] != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; StraceExca[i] >= 0; i += 2 {
			tok := StraceExca[i]
			if tok < TOKSTART || StraceExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if StraceExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += StraceTokname(tok)
	}
	return res
}

func Stracelex1(lex StraceLexer, lval *StraceSymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = StraceTok1[0]
		goto out
	}
	if char < len(StraceTok1) {
		token = StraceTok1[char]
		goto out
	}
	if char >= StracePrivate {
		if char < StracePrivate+len(StraceTok2) {
			token = StraceTok2[char-StracePrivate]
			goto out
		}
	}
	for i := 0; i < len(StraceTok3); i += 2 {
		token = StraceTok3[i+0]
		if token == char {
			token = StraceTok3[i+1]
			goto out
		}
	}

out:
	if token == 0 {
		token = StraceTok2[1] /* unknown char */
	}
	if StraceDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", StraceTokname(token), uint(char))
	}
	return char, token
}

func StraceParse(Stracelex StraceLexer) int {
	return StraceNewParser().Parse(Stracelex)
}

func (Stracercvr *StraceParserImpl) Parse(Stracelex StraceLexer) int {
	var Stracen int
	var StraceVAL StraceSymType
	var StraceDollar []StraceSymType
	_ = StraceDollar // silence set and not used
	StraceS := Stracercvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	Stracestate := 0
	Stracercvr.char = -1
	Stracetoken := -1 // Stracercvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		Stracestate = -1
		Stracercvr.char = -1
		Stracetoken = -1
	}()
	Stracep := -1
	goto Stracestack

ret0:
	return 0

ret1:
	return 1

Stracestack:
	/* put a state and value onto the stack */
	if StraceDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", StraceTokname(Stracetoken), StraceStatname(Stracestate))
	}

	Stracep++
	if Stracep >= len(StraceS) {
		nyys := make([]StraceSymType, len(StraceS)*2)
		copy(nyys, StraceS)
		StraceS = nyys
	}
	StraceS[Stracep] = StraceVAL
	StraceS[Stracep].yys = Stracestate

Stracenewstate:
	Stracen = StracePact[Stracestate]
	if Stracen <= StraceFlag {
		goto Stracedefault /* simple state */
	}
	if Stracercvr.char < 0 {
		Stracercvr.char, Stracetoken = Stracelex1(Stracelex, &Stracercvr.lval)
	}
	Stracen += Stracetoken
	if Stracen < 0 || Stracen >= StraceLast {
		goto Stracedefault
	}
	Stracen = StraceAct[Stracen]
	if StraceChk[Stracen] == Stracetoken { /* valid shift */
		Stracercvr.char = -1
		Stracetoken = -1
		StraceVAL = Stracercvr.lval
		Stracestate = Stracen
		if Errflag > 0 {
			Errflag--
		}
		goto Stracestack
	}

Stracedefault:
	/* default state action */
	Stracen = StraceDef[Stracestate]
	if Stracen == -2 {
		if Stracercvr.char < 0 {
			Stracercvr.char, Stracetoken = Stracelex1(Stracelex, &Stracercvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if StraceExca[xi+0] == -1 && StraceExca[xi+1] == Stracestate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			Stracen = StraceExca[xi+0]
			if Stracen < 0 || Stracen == Stracetoken {
				break
			}
		}
		Stracen = StraceExca[xi+1]
		if Stracen < 0 {
			goto ret0
		}
	}
	if Stracen == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			Stracelex.Error(StraceErrorMessage(Stracestate, Stracetoken))
			Nerrs++
			if StraceDebug >= 1 {
				__yyfmt__.Printf("%s", StraceStatname(Stracestate))
				__yyfmt__.Printf(" saw %s\n", StraceTokname(Stracetoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for Stracep >= 0 {
				Stracen = StracePact[StraceS[Stracep].yys] + StraceErrCode
				if Stracen >= 0 && Stracen < StraceLast {
					Stracestate = StraceAct[Stracen] /* simulate a shift of "error" */
					if StraceChk[Stracestate] == StraceErrCode {
						goto Stracestack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if StraceDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", StraceS[Stracep].yys)
				}
				Stracep--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if StraceDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", StraceTokname(Stracetoken))
			}
			if Stracetoken == StraceEofCode {
				goto ret1
			}
			Stracercvr.char = -1
			Stracetoken = -1
			goto Stracenewstate /* try again in the same state */
		}
	}

	/* reduction by production Stracen */
	if StraceDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", Stracen, StraceStatname(Stracestate))
	}

	Stracent := Stracen
	Stracept := Stracep
	_ = Stracept // guard against "declared and not used"

	Stracep -= StraceR2[Stracen]
	// Stracep is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if Stracep+1 >= len(StraceS) {
		nyys := make([]StraceSymType, len(StraceS)*2)
		copy(nyys, StraceS)
		StraceS = nyys
	}
	StraceVAL = StraceS[Stracep+1]

	/* consult goto table to find next state */
	Stracen = StraceR1[Stracen]
	Straceg := StracePgo[Stracen]
	Stracej := Straceg + StraceS[Stracep].yys + 1

	if Stracej >= StraceLast {
		Stracestate = StraceAct[Straceg]
	} else {
		Stracestate = StraceAct[Stracej]
		if StraceChk[Stracestate] != -Stracen {
			Stracestate = StraceAct[Straceg]
		}
	}
	// dummy call; replaced with literal code
	switch Stracent {

	case 1:
		StraceDollar = StraceS[Stracept-4 : Stracept+1]
		//line scanner/strace.y:78
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, StraceDollar[2].data, nil, int64(-1), true, false)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 2:
		StraceDollar = StraceS[Stracept-5 : Stracept+1]
		//line scanner/strace.y:80
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, StraceDollar[2].data, StraceDollar[4].val_types, int64(-1), true, false)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 3:
		StraceDollar = StraceS[Stracept-6 : Stracept+1]
		//line scanner/strace.y:83
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, "tmp", nil, -1, true, true)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 4:
		StraceDollar = StraceS[Stracept-7 : Stracept+1]
		//line scanner/strace.y:88
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, StraceDollar[2].data, nil, int64(StraceDollar[7].val_int), false, false)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 5:
		StraceDollar = StraceS[Stracept-5 : Stracept+1]
		//line scanner/strace.y:92
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, "tmp", nil, int64(StraceDollar[5].val_int), false, true)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 6:
		StraceDollar = StraceS[Stracept-5 : Stracept+1]
		//line scanner/strace.y:94
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, "tmp", nil, int64(StraceDollar[5].val_uint), false, true)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 7:
		StraceDollar = StraceS[Stracept-5 : Stracept+1]
		//line scanner/strace.y:96
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, "tmp", nil, -1, false, true)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 8:
		StraceDollar = StraceS[Stracept-8 : Stracept+1]
		//line scanner/strace.y:98
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, "tmp", nil, int64(StraceDollar[5].val_int), false, true)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 9:
		StraceDollar = StraceS[Stracept-8 : Stracept+1]
		//line scanner/strace.y:100
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, "tmp", nil, int64(StraceDollar[5].val_uint), false, true)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 10:
		StraceDollar = StraceS[Stracept-6 : Stracept+1]
		//line scanner/strace.y:102
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, "tmp", StraceDollar[3].val_types, int64(StraceDollar[6].val_int), false, true)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 11:
		StraceDollar = StraceS[Stracept-6 : Stracept+1]
		//line scanner/strace.y:104
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, "tmp", StraceDollar[3].val_types, int64(StraceDollar[6].val_uint), false, true)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 12:
		StraceDollar = StraceS[Stracept-6 : Stracept+1]
		//line scanner/strace.y:106
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, "tmp", StraceDollar[3].val_types, -1, false, true)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 13:
		StraceDollar = StraceS[Stracept-9 : Stracept+1]
		//line scanner/strace.y:108
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, "tmp", StraceDollar[3].val_types, int64(StraceDollar[6].val_int), false, true)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 14:
		StraceDollar = StraceS[Stracept-9 : Stracept+1]
		//line scanner/strace.y:110
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, "tmp", StraceDollar[3].val_types, int64(StraceDollar[6].val_uint), false, true)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 15:
		StraceDollar = StraceS[Stracept-6 : Stracept+1]
		//line scanner/strace.y:112
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, StraceDollar[2].data, nil, StraceDollar[6].val_int, false, false)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 16:
		StraceDollar = StraceS[Stracept-7 : Stracept+1]
		//line scanner/strace.y:114
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, StraceDollar[2].data, StraceDollar[4].val_types, StraceDollar[7].val_int, false, false)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 17:
		StraceDollar = StraceS[Stracept-7 : Stracept+1]
		//line scanner/strace.y:117
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, StraceDollar[2].data, StraceDollar[4].val_types, int64(StraceDollar[7].val_uint), false, false)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 18:
		StraceDollar = StraceS[Stracept-7 : Stracept+1]
		//line scanner/strace.y:120
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, StraceDollar[2].data, StraceDollar[4].val_types, -1, false, false)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 19:
		StraceDollar = StraceS[Stracept-11 : Stracept+1]
		//line scanner/strace.y:123
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, StraceDollar[2].data, StraceDollar[4].val_types, StraceDollar[7].val_int, false, false)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 20:
		StraceDollar = StraceS[Stracept-11 : Stracept+1]
		//line scanner/strace.y:126
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, StraceDollar[2].data, StraceDollar[4].val_types, int64(StraceDollar[7].val_uint), false, false)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 21:
		StraceDollar = StraceS[Stracept-10 : Stracept+1]
		//line scanner/strace.y:129
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, StraceDollar[2].data, StraceDollar[4].val_types, StraceDollar[7].val_int, false, false)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 22:
		StraceDollar = StraceS[Stracept-10 : Stracept+1]
		//line scanner/strace.y:132
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, StraceDollar[2].data, StraceDollar[4].val_types, int64(StraceDollar[7].val_uint), false, false)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 23:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:137
		{
			StraceVAL.val_parenthetical = types.NewParenthetical()
		}
	case 24:
		StraceDollar = StraceS[Stracept-2 : Stracept+1]
		//line scanner/strace.y:138
		{
			StraceVAL.val_parenthetical = types.NewParenthetical()
		}
	case 25:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:141
		{
			StraceVAL.val_parenthetical = types.NewParenthetical()
		}
	case 26:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:142
		{
			StraceVAL.val_parenthetical = types.NewParenthetical()
		}
	case 27:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:143
		{
			StraceVAL.val_parenthetical = types.NewParenthetical()
		}
	case 28:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:144
		{
			StraceVAL.val_parenthetical = types.NewParenthetical()
		}
	case 29:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:145
		{
			StraceVAL.val_parenthetical = types.NewParenthetical()
		}
	case 30:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:146
		{
			StraceVAL.val_parenthetical = types.NewParenthetical()
		}
	case 31:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:147
		{
			StraceVAL.val_parenthetical = types.NewParenthetical()
		}
	case 32:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:148
		{
			StraceVAL.val_parenthetical = types.NewParenthetical()
		}
	case 33:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:149
		{
			StraceVAL.val_parenthetical = types.NewParenthetical()
		}
	case 34:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:150
		{
			StraceVAL.val_parenthetical = types.NewParenthetical()
		}
	case 35:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:154
		{
			types := make([]types.Type, 0)
			types = append(types, StraceDollar[1].val_type)
			StraceVAL.val_types = types
		}
	case 36:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:155
		{
			StraceDollar[1].val_types = append(StraceDollar[1].val_types, StraceDollar[3].val_type)
			StraceVAL.val_types = StraceDollar[1].val_types
		}
	case 37:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:160
		{
			StraceVAL.val_type = StraceDollar[1].val_buf_type
		}
	case 38:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:161
		{
			StraceVAL.val_type = StraceDollar[1].val_field
		}
	case 39:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:162
		{
			StraceVAL.val_type = StraceDollar[1].val_expr_type
		}
	case 40:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:163
		{
			StraceVAL.val_type = StraceDollar[1].val_pointer_type
		}
	case 41:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:164
		{
			StraceVAL.val_type = StraceDollar[1].val_array_type
		}
	case 42:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:165
		{
			StraceVAL.val_type = StraceDollar[1].val_struct_type
		}
	case 43:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:166
		{
			StraceVAL.val_type = StraceDollar[1].val_dynamic_type
		}
	case 44:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:167
		{
			StraceVAL.val_type = StraceDollar[1].val_call
		}
	case 45:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:168
		{
			StraceVAL.val_type = StraceDollar[1].val_ip_type
		}
	case 46:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:171
		{
			StraceVAL.val_dynamic_type = types.NewDynamicType(StraceDollar[1].val_expr_type, StraceDollar[3].val_type)
		}
	case 47:
		StraceDollar = StraceS[Stracept-4 : Stracept+1]
		//line scanner/strace.y:174
		{
			StraceVAL.val_call = types.NewCallType(StraceDollar[1].data, StraceDollar[3].val_types)
		}
	case 48:
		StraceDollar = StraceS[Stracept-4 : Stracept+1]
		//line scanner/strace.y:177
		{
			StraceVAL.val_macro = types.NewMacroType(StraceDollar[1].data, StraceDollar[3].val_types)
		}
	case 49:
		StraceDollar = StraceS[Stracept-4 : Stracept+1]
		//line scanner/strace.y:178
		{
			StraceVAL.val_macro = types.NewMacroType(StraceDollar[1].data, nil)
		}
	case 50:
		StraceDollar = StraceS[Stracept-2 : Stracept+1]
		//line scanner/strace.y:181
		{
			StraceVAL.val_pointer_type = types.NullPointer()
		}
	case 51:
		StraceDollar = StraceS[Stracept-4 : Stracept+1]
		//line scanner/strace.y:182
		{
			StraceVAL.val_pointer_type = types.NewPointerType(StraceDollar[2].val_uint, StraceDollar[4].val_type)
		}
	case 52:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:183
		{
			StraceVAL.val_pointer_type = types.NullPointer()
		}
	case 53:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:186
		{
			arr := types.NewArrayType(StraceDollar[2].val_types)
			StraceVAL.val_array_type = arr
		}
	case 54:
		StraceDollar = StraceS[Stracept-2 : Stracept+1]
		//line scanner/strace.y:187
		{
			arr := types.NewArrayType(nil)
			StraceVAL.val_array_type = arr
		}
	case 55:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:190
		{
			StraceVAL.val_struct_type = types.NewStructType(StraceDollar[2].val_types)
		}
	case 56:
		StraceDollar = StraceS[Stracept-4 : Stracept+1]
		//line scanner/strace.y:191
		{
			StraceVAL.val_struct_type = types.NewStructType(StraceDollar[2].val_types)
		}
	case 57:
		StraceDollar = StraceS[Stracept-2 : Stracept+1]
		//line scanner/strace.y:194
		{
			StraceVAL.val_field = types.NewField(StraceDollar[1].data, nil)
		}
	case 58:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:195
		{
			StraceVAL.val_field = types.NewField(StraceDollar[1].data, StraceDollar[3].val_type)
		}
	case 59:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:196
		{
			StraceVAL.val_field = types.NewField(StraceDollar[1].data, StraceDollar[3].val_type)
		}
	case 60:
		StraceDollar = StraceS[Stracept-4 : Stracept+1]
		//line scanner/strace.y:197
		{
			StraceVAL.val_field = types.NewField(StraceDollar[1].data, StraceDollar[4].val_type)
		}
	case 61:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:200
		{
			StraceVAL.val_buf_type = types.NewBufferType(StraceDollar[1].data)
		}
	case 62:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:201
		{
			StraceVAL.val_buf_type = types.NewBufferType(StraceDollar[1].data)
		}
	case 63:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:204
		{
			StraceVAL.val_expr_type = types.NewExpression(StraceDollar[1].val_flag_type)
		}
	case 64:
		StraceDollar = StraceS[Stracept-4 : Stracept+1]
		//line scanner/strace.y:205
		{
			expr1 := types.NewExpression(types.NewFlagType(StraceDollar[2].data))
			expr2 := types.NewExpression(types.NewFlagType(StraceDollar[3].data))
			bs := types.NewBinarySet(expr1, expr2)
			StraceVAL.val_expr_type = types.NewExpression(bs)
		}
	case 65:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:210
		{
			StraceVAL.val_expr_type = types.NewExpression(StraceDollar[1].val_int_type)
		}
	case 66:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:211
		{
			StraceVAL.val_expr_type = types.NewExpression(StraceDollar[1].val_macro)
		}
	case 67:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:212
		{
			StraceVAL.val_expr_type = StraceDollar[2].val_expr_type
		}
	case 68:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:213
		{
			StraceVAL.val_expr_type = types.NewExpression(types.NewBinop(types.NewExpression(StraceDollar[1].val_macro), types.OR, StraceDollar[3].val_expr_type))
		}
	case 69:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:214
		{
			StraceVAL.val_expr_type = types.NewExpression(types.NewBinop(types.NewExpression(StraceDollar[1].val_flag_type), types.OR, StraceDollar[3].val_expr_type))
		}
	case 70:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:215
		{
			StraceVAL.val_expr_type = types.NewExpression(types.NewBinop(types.NewExpression(StraceDollar[1].val_int_type), types.OR, StraceDollar[3].val_expr_type))
		}
	case 71:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:216
		{
			StraceVAL.val_expr_type = types.NewExpression(types.NewBinop(types.NewExpression(StraceDollar[1].val_macro), types.OR, StraceDollar[3].val_expr_type))
		}
	case 72:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:217
		{
			StraceVAL.val_expr_type = types.NewExpression(types.NewBinop(types.NewExpression(StraceDollar[1].val_flag_type), types.LAND, StraceDollar[3].val_expr_type))
		}
	case 73:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:218
		{
			StraceVAL.val_expr_type = types.NewExpression(types.NewBinop(types.NewExpression(StraceDollar[1].val_int_type), types.LAND, StraceDollar[3].val_expr_type))
		}
	case 74:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:219
		{
			StraceVAL.val_expr_type = types.NewExpression(types.NewBinop(types.NewExpression(StraceDollar[1].val_macro), types.LOR, StraceDollar[3].val_expr_type))
		}
	case 75:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:220
		{
			StraceVAL.val_expr_type = types.NewExpression(types.NewBinop(types.NewExpression(StraceDollar[1].val_flag_type), types.LOR, StraceDollar[3].val_expr_type))
		}
	case 76:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:221
		{
			StraceVAL.val_expr_type = types.NewExpression(types.NewBinop(types.NewExpression(StraceDollar[1].val_int_type), types.LOR, StraceDollar[3].val_expr_type))
		}
	case 77:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:222
		{
			StraceVAL.val_expr_type = types.NewExpression(types.NewBinop(types.NewExpression(StraceDollar[1].val_macro), types.LEQUAL, StraceDollar[3].val_expr_type))
		}
	case 78:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:223
		{
			StraceVAL.val_expr_type = types.NewExpression(types.NewBinop(types.NewExpression(StraceDollar[1].val_flag_type), types.LEQUAL, StraceDollar[3].val_expr_type))
		}
	case 79:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:224
		{
			StraceVAL.val_expr_type = types.NewExpression(types.NewBinop(types.NewExpression(StraceDollar[1].val_int_type), types.LEQUAL, StraceDollar[3].val_expr_type))
		}
	case 80:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:225
		{
			StraceVAL.val_expr_type = types.NewExpression(types.NewBinop(types.NewExpression(StraceDollar[1].val_macro), types.AND, StraceDollar[3].val_expr_type))
		}
	case 81:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:226
		{
			StraceVAL.val_expr_type = types.NewExpression(types.NewBinop(types.NewExpression(StraceDollar[1].val_flag_type), types.AND, StraceDollar[3].val_expr_type))
		}
	case 82:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:227
		{
			StraceVAL.val_expr_type = types.NewExpression(types.NewBinop(types.NewExpression(StraceDollar[1].val_int_type), types.AND, StraceDollar[3].val_expr_type))
		}
	case 83:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:228
		{
			StraceVAL.val_expr_type = types.NewExpression(types.NewBinop(types.NewExpression(StraceDollar[1].val_macro), types.LSHIFT, StraceDollar[3].val_expr_type))
		}
	case 84:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:229
		{
			StraceVAL.val_expr_type = types.NewExpression(types.NewBinop(types.NewExpression(StraceDollar[1].val_flag_type), types.LSHIFT, StraceDollar[3].val_expr_type))
		}
	case 85:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:230
		{
			StraceVAL.val_expr_type = types.NewExpression(types.NewBinop(types.NewExpression(StraceDollar[1].val_int_type), types.LSHIFT, StraceDollar[3].val_expr_type))
		}
	case 86:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:231
		{
			StraceVAL.val_expr_type = types.NewExpression(types.NewBinop(types.NewExpression(StraceDollar[1].val_macro), types.RSHIFT, StraceDollar[3].val_expr_type))
		}
	case 87:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:232
		{
			StraceVAL.val_expr_type = types.NewExpression(types.NewBinop(types.NewExpression(StraceDollar[1].val_flag_type), types.RSHIFT, StraceDollar[3].val_expr_type))
		}
	case 88:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:233
		{
			StraceVAL.val_expr_type = types.NewExpression(types.NewBinop(types.NewExpression(StraceDollar[1].val_int_type), types.RSHIFT, StraceDollar[3].val_expr_type))
		}
	case 89:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:234
		{
			StraceVAL.val_expr_type = types.NewExpression(types.NewBinop(types.NewExpression(StraceDollar[1].val_int_type), types.TIMES, StraceDollar[3].val_expr_type))
		}
	case 90:
		StraceDollar = StraceS[Stracept-2 : Stracept+1]
		//line scanner/strace.y:235
		{
			StraceVAL.val_expr_type = types.NewExpression(types.NewUnop(StraceDollar[2].val_expr_type, types.ONESCOMP))
		}
	case 91:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:240
		{
			StraceVAL.val_int_type = types.NewIntType(StraceDollar[1].val_int)
		}
	case 92:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:241
		{
			StraceVAL.val_int_type = types.NewIntType(int64(StraceDollar[1].val_uint))
		}
	case 93:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:244
		{
			StraceVAL.val_flag_type = types.NewFlagType(StraceDollar[1].data)
		}
	case 94:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:247
		{
			StraceVAL.val_ip_type = types.NewIpType(StraceDollar[1].data)
		}
	case 95:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:248
		{
			StraceVAL.val_ip_type = types.NewIpType(StraceDollar[1].data)
		}
	case 96:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:249
		{
			StraceVAL.val_ip_type = types.NewIpType(StraceDollar[1].data)
		}
	case 97:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:252
		{
			ids := make([]*types.BufferType, 0)
			ids = append(ids, types.NewBufferType(StraceDollar[1].data))
			StraceVAL.val_identifiers = ids
		}
	case 98:
		StraceDollar = StraceS[Stracept-2 : Stracept+1]
		//line scanner/strace.y:253
		{
			StraceDollar[2].val_identifiers = append(StraceDollar[2].val_identifiers, types.NewBufferType(StraceDollar[1].data))
			StraceVAL.val_identifiers = StraceDollar[2].val_identifiers
		}
	}
	goto Stracestack /* stack new state and value */
}
