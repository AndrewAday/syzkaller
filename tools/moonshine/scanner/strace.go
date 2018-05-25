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
	val_array_type    *types.ArrayType
	val_pointer_type  *types.PointerType
	val_flag_type     *types.FlagType
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

const StraceLast = 613

var StraceAct = [...]int{

	33, 9, 116, 57, 102, 58, 8, 44, 44, 101,
	59, 93, 38, 114, 71, 117, 70, 145, 13, 46,
	47, 50, 53, 51, 52, 3, 48, 49, 78, 74,
	66, 62, 64, 17, 36, 97, 14, 54, 44, 34,
	35, 57, 73, 58, 110, 44, 79, 80, 59, 42,
	56, 43, 55, 150, 30, 4, 149, 24, 141, 32,
	90, 92, 146, 36, 89, 140, 68, 138, 34, 35,
	115, 136, 99, 72, 41, 147, 98, 144, 135, 56,
	81, 82, 83, 84, 85, 86, 87, 88, 32, 134,
	55, 106, 139, 111, 105, 112, 137, 79, 44, 69,
	5, 95, 94, 96, 53, 53, 131, 131, 48, 49,
	128, 67, 118, 119, 104, 120, 34, 35, 121, 60,
	103, 124, 125, 2, 130, 130, 61, 67, 126, 127,
	117, 123, 25, 16, 24, 131, 131, 157, 131, 12,
	131, 29, 129, 129, 31, 131, 131, 15, 131, 131,
	10, 11, 131, 130, 130, 131, 130, 148, 130, 1,
	132, 132, 151, 130, 130, 154, 130, 130, 122, 133,
	130, 129, 129, 130, 129, 53, 129, 52, 0, 48,
	49, 129, 129, 0, 129, 129, 0, 0, 129, 132,
	132, 129, 132, 0, 132, 0, 0, 142, 143, 132,
	132, 0, 132, 132, 107, 108, 132, 109, 153, 132,
	0, 153, 75, 76, 153, 77, 0, 153, 19, 26,
	27, 21, 36, 20, 0, 0, 28, 34, 35, 46,
	47, 0, 53, 22, 52, 0, 48, 49, 18, 0,
	0, 0, 0, 25, 0, 24, 0, 32, 40, 0,
	37, 39, 23, 19, 26, 27, 21, 36, 20, 0,
	0, 28, 34, 35, 0, 47, 0, 53, 22, 52,
	0, 48, 49, 18, 0, 0, 0, 0, 25, 0,
	24, 0, 32, 0, 0, 0, 0, 23, 91, 19,
	26, 27, 21, 36, 20, 0, 0, 28, 34, 35,
	0, 0, 0, 0, 22, 0, 0, 0, 0, 18,
	0, 0, 0, 0, 25, 0, 24, 0, 32, 7,
	0, 6, 0, 23, 19, 26, 27, 21, 36, 20,
	0, 0, 28, 34, 35, 0, 0, 0, 0, 22,
	0, 0, 0, 0, 18, 0, 0, 0, 0, 25,
	113, 24, 0, 32, 0, 0, 0, 0, 23, 19,
	26, 27, 21, 36, 20, 0, 0, 28, 34, 35,
	0, 0, 0, 0, 22, 0, 0, 0, 0, 18,
	0, 0, 0, 0, 25, 65, 24, 0, 32, 0,
	0, 0, 0, 23, 19, 26, 27, 21, 36, 20,
	0, 0, 28, 34, 35, 0, 0, 0, 0, 22,
	0, 0, 0, 0, 18, 0, 0, 0, 0, 25,
	0, 24, 63, 32, 0, 0, 0, 0, 23, 19,
	26, 27, 21, 36, 20, 0, 0, 28, 34, 35,
	0, 0, 0, 0, 22, 0, 0, 0, 0, 18,
	0, 0, 0, 0, 25, 0, 24, 0, 32, 0,
	0, 0, 0, 23, 19, 26, 27, 100, 36, 20,
	0, 0, 28, 34, 35, 0, 0, 0, 0, 22,
	0, 0, 0, 0, 18, 0, 0, 0, 0, 25,
	0, 24, 0, 32, 128, 67, 0, 0, 23, 0,
	34, 35, 0, 0, 0, 124, 125, 0, 0, 0,
	0, 0, 126, 127, 0, 123, 25, 0, 24, 128,
	67, 156, 0, 0, 0, 34, 35, 0, 0, 0,
	124, 125, 0, 0, 0, 0, 0, 126, 127, 0,
	123, 25, 0, 24, 128, 67, 155, 0, 0, 0,
	34, 35, 0, 0, 0, 124, 125, 0, 0, 0,
	0, 0, 126, 127, 0, 123, 25, 0, 24, 128,
	67, 152, 0, 0, 0, 34, 35, 0, 0, 0,
	124, 125, 0, 0, 0, 0, 0, 126, 127, 0,
	123, 25, 0, 24, 45, 46, 47, 50, 53, 51,
	52, 0, 48, 49, 46, 47, 0, 53, 51, 52,
	0, 48, 49,
}
var StracePact = [...]int{

	110, -1000, 18, 67, 285, 214, 40, 14, 17, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 577, 26, -1000,
	-1000, -30, 112, -1000, 390, 355, -1000, -1000, -1000, 119,
	-1000, -1000, 55, -1000, -1000, -1000, 66, -1000, -20, 39,
	7, -6, 199, -7, 425, 425, 55, 55, 55, 55,
	55, 55, 55, 55, -1000, -1000, 55, 425, 249, 425,
	-1000, -24, 70, -1000, 73, -1000, -1000, -1000, 1, 460,
	-1000, -26, -31, 107, 98, 61, 58, -1000, 191, -1000,
	-1000, 246, 154, 84, 84, 586, 211, 83, -1000, 10,
	-1000, 425, -1000, 425, -1000, 320, -1000, -1000, -21, 36,
	8, 99, 105, -1000, -1000, 562, 562, 56, 45, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, 123, 63, 59,
	-1000, -1000, 31, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, 24, 562, 562, 44, 562, 42, 562,
	-1000, -1000, 22, 19, 562, 537, -1000, 562, 512, -1000,
	-1000, 487, -1000, -1000, 103, -1000, -1000, -1000,
}
var StracePgo = [...]int{

	0, 159, 151, 2, 54, 150, 36, 18, 0, 147,
	62, 17, 144, 1, 33, 141, 139, 133, 6,
}
var StraceR1 = [...]int{

	0, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 11, 11, 10, 10, 10, 10, 10,
	10, 10, 10, 10, 10, 18, 18, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 14, 14, 14,
	14, 14, 14, 14, 14, 14, 14, 14, 14, 14,
	15, 15, 9, 12, 12, 16, 16, 16, 7, 7,
	6, 6, 6, 2, 2, 2, 2, 5, 5, 4,
	4, 8, 17, 17, 17, 3, 3,
}
var StraceR2 = [...]int{

	0, 4, 5, 6, 7, 5, 5, 5, 8, 8,
	6, 6, 6, 9, 9, 6, 7, 7, 7, 11,
	11, 10, 10, 1, 2, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 3, 1, 1, 1,
	1, 1, 1, 1, 1, 3, 2, 1, 1, 1,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 2,
	1, 2, 4, 4, 4, 2, 4, 1, 3, 2,
	3, 4, 2, 2, 3, 3, 4, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 2,
}
var StraceChk = [...]int{

	-1000, -1, 13, 7, 37, 33, 36, 34, -18, -13,
	-5, -2, -16, -7, -6, -9, -17, -14, 24, 4,
	9, 7, 19, 38, 31, 29, 5, 6, 12, -15,
	-4, -12, 33, -8, 13, 14, 8, 36, -18, 37,
	34, 34, 35, 34, 28, 17, 18, 19, 25, 26,
	20, 22, 23, 21, -7, -14, 24, 33, 35, 40,
	7, 14, -18, 32, -18, 30, -8, 8, -14, 33,
	36, 34, 34, 35, 35, 13, 14, 16, 35, -13,
	-13, -14, -14, -14, -14, -14, -14, -14, -14, -18,
	-13, 39, -13, 35, 32, 28, 30, 34, -18, -3,
	7, 35, 35, 13, 16, 33, 33, 13, 14, 16,
	34, -13, -13, 30, 34, 34, -3, 7, 13, 14,
	16, 13, -10, 28, 18, 19, 25, 26, 7, -6,
	-7, -8, -4, -10, 33, 33, 8, 33, 8, 33,
	34, 34, -10, -10, 33, -11, -10, 33, -11, 34,
	34, -11, 34, -10, -11, 34, 34, 34,
}
var StraceDef = [...]int{

	0, -2, 0, 0, 0, 0, 0, 0, 0, 35,
	37, 38, 39, 40, 41, 42, 43, 44, 0, 77,
	78, 0, 0, 67, 0, 0, 82, 83, 84, 47,
	48, 49, 0, 60, 79, 80, 81, 1, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 46, 59, 0, 0, 73, 0,
	65, 0, 0, 69, 0, 72, 61, 81, 0, 0,
	2, 0, 0, 0, 0, 5, 6, 7, 0, 36,
	45, 50, 51, 52, 53, 54, 55, 56, 58, 0,
	74, 0, 75, 0, 68, 0, 70, 57, 0, 0,
	85, 0, 0, 15, 3, 0, 0, 10, 11, 12,
	62, 76, 66, 71, 63, 64, 86, 85, 16, 17,
	18, 4, 0, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 0, 0, 0, 0, 0, 0, 0,
	8, 9, 0, 0, 0, 0, 23, 0, 0, 13,
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
		//line scanner/strace.y:74
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, StraceDollar[2].data, nil, int64(-1), true, false)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 2:
		StraceDollar = StraceS[Stracept-5 : Stracept+1]
		//line scanner/strace.y:76
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, StraceDollar[2].data, StraceDollar[4].val_types, int64(-1), true, false)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 3:
		StraceDollar = StraceS[Stracept-6 : Stracept+1]
		//line scanner/strace.y:79
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, "tmp", nil, -1, true, true)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 4:
		StraceDollar = StraceS[Stracept-7 : Stracept+1]
		//line scanner/strace.y:84
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, StraceDollar[2].data, nil, int64(StraceDollar[7].val_int), false, false)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 5:
		StraceDollar = StraceS[Stracept-5 : Stracept+1]
		//line scanner/strace.y:88
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, "tmp", nil, int64(StraceDollar[5].val_int), false, true)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 6:
		StraceDollar = StraceS[Stracept-5 : Stracept+1]
		//line scanner/strace.y:90
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, "tmp", nil, int64(StraceDollar[5].val_uint), false, true)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 7:
		StraceDollar = StraceS[Stracept-5 : Stracept+1]
		//line scanner/strace.y:92
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, "tmp", nil, -1, false, true)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 8:
		StraceDollar = StraceS[Stracept-8 : Stracept+1]
		//line scanner/strace.y:94
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, "tmp", nil, int64(StraceDollar[5].val_int), false, true)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 9:
		StraceDollar = StraceS[Stracept-8 : Stracept+1]
		//line scanner/strace.y:96
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, "tmp", nil, int64(StraceDollar[5].val_uint), false, true)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 10:
		StraceDollar = StraceS[Stracept-6 : Stracept+1]
		//line scanner/strace.y:98
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, "tmp", StraceDollar[3].val_types, int64(StraceDollar[6].val_int), false, true)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 11:
		StraceDollar = StraceS[Stracept-6 : Stracept+1]
		//line scanner/strace.y:100
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, "tmp", StraceDollar[3].val_types, int64(StraceDollar[6].val_uint), false, true)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 12:
		StraceDollar = StraceS[Stracept-6 : Stracept+1]
		//line scanner/strace.y:102
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, "tmp", StraceDollar[3].val_types, -1, false, true)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 13:
		StraceDollar = StraceS[Stracept-9 : Stracept+1]
		//line scanner/strace.y:104
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, "tmp", StraceDollar[3].val_types, int64(StraceDollar[6].val_int), false, true)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 14:
		StraceDollar = StraceS[Stracept-9 : Stracept+1]
		//line scanner/strace.y:106
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, "tmp", StraceDollar[3].val_types, int64(StraceDollar[6].val_uint), false, true)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 15:
		StraceDollar = StraceS[Stracept-6 : Stracept+1]
		//line scanner/strace.y:108
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, StraceDollar[2].data, nil, StraceDollar[6].val_int, false, false)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 16:
		StraceDollar = StraceS[Stracept-7 : Stracept+1]
		//line scanner/strace.y:110
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, StraceDollar[2].data, StraceDollar[4].val_types, StraceDollar[7].val_int, false, false)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 17:
		StraceDollar = StraceS[Stracept-7 : Stracept+1]
		//line scanner/strace.y:113
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, StraceDollar[2].data, StraceDollar[4].val_types, int64(StraceDollar[7].val_uint), false, false)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 18:
		StraceDollar = StraceS[Stracept-7 : Stracept+1]
		//line scanner/strace.y:116
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, StraceDollar[2].data, StraceDollar[4].val_types, -1, false, false)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 19:
		StraceDollar = StraceS[Stracept-11 : Stracept+1]
		//line scanner/strace.y:119
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, StraceDollar[2].data, StraceDollar[4].val_types, StraceDollar[7].val_int, false, false)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 20:
		StraceDollar = StraceS[Stracept-11 : Stracept+1]
		//line scanner/strace.y:122
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, StraceDollar[2].data, StraceDollar[4].val_types, int64(StraceDollar[7].val_uint), false, false)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 21:
		StraceDollar = StraceS[Stracept-10 : Stracept+1]
		//line scanner/strace.y:125
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, StraceDollar[2].data, StraceDollar[4].val_types, StraceDollar[7].val_int, false, false)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 22:
		StraceDollar = StraceS[Stracept-10 : Stracept+1]
		//line scanner/strace.y:128
		{
			StraceVAL.val_syscall = types.NewSyscall(StraceDollar[1].val_int, StraceDollar[2].data, StraceDollar[4].val_types, int64(StraceDollar[7].val_uint), false, false)
			Stracelex.(*lexer).result = StraceVAL.val_syscall
		}
	case 23:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:133
		{
			StraceVAL.val_parenthetical = types.NewParenthetical()
		}
	case 24:
		StraceDollar = StraceS[Stracept-2 : Stracept+1]
		//line scanner/strace.y:134
		{
			StraceVAL.val_parenthetical = types.NewParenthetical()
		}
	case 25:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:137
		{
			StraceVAL.val_parenthetical = types.NewParenthetical()
		}
	case 26:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:138
		{
			StraceVAL.val_parenthetical = types.NewParenthetical()
		}
	case 27:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:139
		{
			StraceVAL.val_parenthetical = types.NewParenthetical()
		}
	case 28:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:140
		{
			StraceVAL.val_parenthetical = types.NewParenthetical()
		}
	case 29:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:141
		{
			StraceVAL.val_parenthetical = types.NewParenthetical()
		}
	case 30:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:142
		{
			StraceVAL.val_parenthetical = types.NewParenthetical()
		}
	case 31:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:143
		{
			StraceVAL.val_parenthetical = types.NewParenthetical()
		}
	case 32:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:144
		{
			StraceVAL.val_parenthetical = types.NewParenthetical()
		}
	case 33:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:145
		{
			StraceVAL.val_parenthetical = types.NewParenthetical()
		}
	case 34:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:146
		{
			StraceVAL.val_parenthetical = types.NewParenthetical()
		}
	case 35:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:150
		{
			types := make([]types.Type, 0)
			types = append(types, StraceDollar[1].val_type)
			StraceVAL.val_types = types
		}
	case 36:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:151
		{
			StraceDollar[1].val_types = append(StraceDollar[1].val_types, StraceDollar[3].val_type)
			StraceVAL.val_types = StraceDollar[1].val_types
		}
	case 37:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:155
		{
			StraceVAL.val_type = StraceDollar[1].val_buf_type
		}
	case 38:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:156
		{
			StraceVAL.val_type = StraceDollar[1].val_field
		}
	case 39:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:157
		{
			StraceVAL.val_type = StraceDollar[1].val_pointer_type
		}
	case 40:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:158
		{
			StraceVAL.val_type = StraceDollar[1].val_array_type
		}
	case 41:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:159
		{
			StraceVAL.val_type = StraceDollar[1].val_struct_type
		}
	case 42:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:160
		{
			StraceVAL.val_type = StraceDollar[1].val_call
		}
	case 43:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:161
		{
			StraceVAL.val_type = StraceDollar[1].val_ip_type
		}
	case 44:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:162
		{
			StraceVAL.val_type = StraceDollar[1].val_type
		}
	case 45:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:163
		{
			StraceVAL.val_type = types.NewDynamicType(StraceDollar[1].val_type, StraceDollar[3].val_type)
		}
	case 46:
		StraceDollar = StraceS[Stracept-2 : Stracept+1]
		//line scanner/strace.y:164
		{
			StraceVAL.val_type = StraceDollar[2].val_array_type
		}
	case 47:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:168
		{
			StraceVAL.val_type = types.NewExpression(StraceDollar[1].val_type)
		}
	case 48:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:169
		{
			StraceVAL.val_type = types.NewExpression(StraceDollar[1].val_int_type)
		}
	case 49:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:170
		{
			StraceVAL.val_type = types.NewExpression(StraceDollar[1].val_macro)
		}
	case 50:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:171
		{
			StraceVAL.val_type = types.NewExpression(types.NewBinop(StraceDollar[1].val_type, types.OR, StraceDollar[3].val_type))
		}
	case 51:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:172
		{
			StraceVAL.val_type = types.NewExpression(types.NewBinop(StraceDollar[1].val_type, types.AND, StraceDollar[3].val_type))
		}
	case 52:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:173
		{
			StraceVAL.val_type = types.NewExpression(types.NewBinop(StraceDollar[1].val_type, types.LSHIFT, StraceDollar[3].val_type))
		}
	case 53:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:174
		{
			StraceVAL.val_type = types.NewExpression(types.NewBinop(StraceDollar[1].val_type, types.RSHIFT, StraceDollar[3].val_type))
		}
	case 54:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:175
		{
			StraceVAL.val_type = types.NewExpression(types.NewBinop(StraceDollar[1].val_type, types.LOR, StraceDollar[3].val_type))
		}
	case 55:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:176
		{
			StraceVAL.val_type = types.NewExpression(types.NewBinop(StraceDollar[1].val_type, types.LAND, StraceDollar[3].val_type))
		}
	case 56:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:177
		{
			StraceVAL.val_type = types.NewExpression(types.NewBinop(StraceDollar[1].val_type, types.LEQUAL, StraceDollar[3].val_type))
		}
	case 57:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:178
		{
			StraceVAL.val_type = StraceDollar[2].val_type
		}
	case 58:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:179
		{
			StraceVAL.val_type = types.NewExpression(types.NewBinop(StraceDollar[1].val_type, types.TIMES, StraceDollar[3].val_type))
		}
	case 59:
		StraceDollar = StraceS[Stracept-2 : Stracept+1]
		//line scanner/strace.y:180
		{
			StraceVAL.val_type = types.NewExpression(types.NewUnop(StraceDollar[2].val_type, types.ONESCOMP))
		}
	case 60:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:183
		{
			f := make(types.Flags, 1)
			f[0] = StraceDollar[1].val_flag_type
			StraceVAL.val_type = f
		}
	case 61:
		StraceDollar = StraceS[Stracept-2 : Stracept+1]
		//line scanner/strace.y:184
		{
			StraceVAL.val_type = append(StraceDollar[1].val_type.(types.Flags), StraceDollar[2].val_flag_type)
		}
	case 62:
		StraceDollar = StraceS[Stracept-4 : Stracept+1]
		//line scanner/strace.y:187
		{
			StraceVAL.val_call = types.NewCallType(StraceDollar[1].data, StraceDollar[3].val_types)
		}
	case 63:
		StraceDollar = StraceS[Stracept-4 : Stracept+1]
		//line scanner/strace.y:190
		{
			StraceVAL.val_macro = types.NewMacroType(StraceDollar[1].data, StraceDollar[3].val_types)
		}
	case 64:
		StraceDollar = StraceS[Stracept-4 : Stracept+1]
		//line scanner/strace.y:191
		{
			StraceVAL.val_macro = types.NewMacroType(StraceDollar[1].data, nil)
		}
	case 65:
		StraceDollar = StraceS[Stracept-2 : Stracept+1]
		//line scanner/strace.y:194
		{
			StraceVAL.val_pointer_type = types.NullPointer()
		}
	case 66:
		StraceDollar = StraceS[Stracept-4 : Stracept+1]
		//line scanner/strace.y:195
		{
			StraceVAL.val_pointer_type = types.NewPointerType(StraceDollar[2].val_uint, StraceDollar[4].val_type)
		}
	case 67:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:196
		{
			StraceVAL.val_pointer_type = types.NullPointer()
		}
	case 68:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:199
		{
			arr := types.NewArrayType(StraceDollar[2].val_types)
			StraceVAL.val_array_type = arr
		}
	case 69:
		StraceDollar = StraceS[Stracept-2 : Stracept+1]
		//line scanner/strace.y:200
		{
			arr := types.NewArrayType(nil)
			StraceVAL.val_array_type = arr
		}
	case 70:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:203
		{
			StraceVAL.val_struct_type = types.NewStructType(StraceDollar[2].val_types)
		}
	case 71:
		StraceDollar = StraceS[Stracept-4 : Stracept+1]
		//line scanner/strace.y:204
		{
			StraceVAL.val_struct_type = types.NewStructType(StraceDollar[2].val_types)
		}
	case 72:
		StraceDollar = StraceS[Stracept-2 : Stracept+1]
		//line scanner/strace.y:205
		{
			StraceVAL.val_struct_type = types.NewStructType(nil)
		}
	case 73:
		StraceDollar = StraceS[Stracept-2 : Stracept+1]
		//line scanner/strace.y:208
		{
			StraceVAL.val_field = types.NewField(StraceDollar[1].data, nil)
		}
	case 74:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:209
		{
			StraceVAL.val_field = types.NewField(StraceDollar[1].data, StraceDollar[3].val_type)
		}
	case 75:
		StraceDollar = StraceS[Stracept-3 : Stracept+1]
		//line scanner/strace.y:210
		{
			StraceVAL.val_field = types.NewField(StraceDollar[1].data, StraceDollar[3].val_type)
		}
	case 76:
		StraceDollar = StraceS[Stracept-4 : Stracept+1]
		//line scanner/strace.y:211
		{
			StraceVAL.val_field = types.NewField(StraceDollar[1].data, StraceDollar[4].val_type)
		}
	case 77:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:214
		{
			StraceVAL.val_buf_type = types.NewBufferType(StraceDollar[1].data)
		}
	case 78:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:215
		{
			StraceVAL.val_buf_type = types.NewBufferType(StraceDollar[1].data)
		}
	case 79:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:219
		{
			StraceVAL.val_int_type = types.NewIntType(StraceDollar[1].val_int)
		}
	case 80:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:220
		{
			StraceVAL.val_int_type = types.NewIntType(int64(StraceDollar[1].val_uint))
		}
	case 81:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:223
		{
			StraceVAL.val_flag_type = types.NewFlagType(StraceDollar[1].data)
		}
	case 82:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:226
		{
			StraceVAL.val_ip_type = types.NewIpType(StraceDollar[1].data)
		}
	case 83:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:227
		{
			StraceVAL.val_ip_type = types.NewIpType(StraceDollar[1].data)
		}
	case 84:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:228
		{
			StraceVAL.val_ip_type = types.NewIpType(StraceDollar[1].data)
		}
	case 85:
		StraceDollar = StraceS[Stracept-1 : Stracept+1]
		//line scanner/strace.y:231
		{
			ids := make([]*types.BufferType, 0)
			ids = append(ids, types.NewBufferType(StraceDollar[1].data))
			StraceVAL.val_identifiers = ids
		}
	case 86:
		StraceDollar = StraceS[Stracept-2 : Stracept+1]
		//line scanner/strace.y:232
		{
			StraceDollar[2].val_identifiers = append(StraceDollar[2].val_identifiers, types.NewBufferType(StraceDollar[1].data))
			StraceVAL.val_identifiers = StraceDollar[2].val_identifiers
		}
	}
	goto Stracestack /* stack new state and value */
}
