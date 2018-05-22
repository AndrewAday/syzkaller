
// line 1 "scanner/lex.rl"
package scanner

import (
    "fmt"
    "encoding/hex"
    "strconv"
    "strings"
    "github.com/google/syzkaller/tools/moonshine/strace_types"
)


// line 15 "scanner/lex.go"
const strace_start int = 111
const strace_first_final int = 111
const strace_error int = 0

const strace_en_comment int = 139
const strace_en_main int = 111


// line 17 "scanner/lex.rl"


type lexer struct {
    result *strace_types.Syscall
    data []byte
    p, pe, cs int
    ts, te, act int
}

func newLexer (data []byte) *lexer {
    lex := &lexer {
        data: data,
        pe: len(data),
    }

    
// line 41 "scanner/lex.go"
	{
	 lex.cs = strace_start
	 lex.ts = 0
	 lex.te = 0
	 lex.act = 0
	}

// line 33 "scanner/lex.rl"
    return lex
}

func (lex *lexer) Lex(out *StraceSymType) int {
    eof := lex.pe
    tok := 0
    
// line 57 "scanner/lex.go"
	{
	if ( lex.p) == ( lex.pe) {
		goto _test_eof
	}
	switch  lex.cs {
	case 111:
		goto st_case_111
	case 0:
		goto st_case_0
	case 1:
		goto st_case_1
	case 112:
		goto st_case_112
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	case 23:
		goto st_case_23
	case 24:
		goto st_case_24
	case 113:
		goto st_case_113
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 114:
		goto st_case_114
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
	case 115:
		goto st_case_115
	case 116:
		goto st_case_116
	case 117:
		goto st_case_117
	case 118:
		goto st_case_118
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 32:
		goto st_case_32
	case 33:
		goto st_case_33
	case 34:
		goto st_case_34
	case 35:
		goto st_case_35
	case 36:
		goto st_case_36
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 39:
		goto st_case_39
	case 40:
		goto st_case_40
	case 41:
		goto st_case_41
	case 42:
		goto st_case_42
	case 43:
		goto st_case_43
	case 44:
		goto st_case_44
	case 45:
		goto st_case_45
	case 46:
		goto st_case_46
	case 47:
		goto st_case_47
	case 119:
		goto st_case_119
	case 120:
		goto st_case_120
	case 121:
		goto st_case_121
	case 122:
		goto st_case_122
	case 48:
		goto st_case_48
	case 49:
		goto st_case_49
	case 50:
		goto st_case_50
	case 51:
		goto st_case_51
	case 52:
		goto st_case_52
	case 53:
		goto st_case_53
	case 54:
		goto st_case_54
	case 55:
		goto st_case_55
	case 56:
		goto st_case_56
	case 57:
		goto st_case_57
	case 58:
		goto st_case_58
	case 59:
		goto st_case_59
	case 60:
		goto st_case_60
	case 61:
		goto st_case_61
	case 62:
		goto st_case_62
	case 63:
		goto st_case_63
	case 64:
		goto st_case_64
	case 65:
		goto st_case_65
	case 66:
		goto st_case_66
	case 123:
		goto st_case_123
	case 67:
		goto st_case_67
	case 124:
		goto st_case_124
	case 125:
		goto st_case_125
	case 68:
		goto st_case_68
	case 69:
		goto st_case_69
	case 70:
		goto st_case_70
	case 71:
		goto st_case_71
	case 126:
		goto st_case_126
	case 127:
		goto st_case_127
	case 128:
		goto st_case_128
	case 129:
		goto st_case_129
	case 130:
		goto st_case_130
	case 72:
		goto st_case_72
	case 73:
		goto st_case_73
	case 74:
		goto st_case_74
	case 75:
		goto st_case_75
	case 76:
		goto st_case_76
	case 77:
		goto st_case_77
	case 78:
		goto st_case_78
	case 79:
		goto st_case_79
	case 80:
		goto st_case_80
	case 81:
		goto st_case_81
	case 82:
		goto st_case_82
	case 83:
		goto st_case_83
	case 84:
		goto st_case_84
	case 85:
		goto st_case_85
	case 131:
		goto st_case_131
	case 86:
		goto st_case_86
	case 87:
		goto st_case_87
	case 88:
		goto st_case_88
	case 89:
		goto st_case_89
	case 90:
		goto st_case_90
	case 91:
		goto st_case_91
	case 92:
		goto st_case_92
	case 93:
		goto st_case_93
	case 94:
		goto st_case_94
	case 95:
		goto st_case_95
	case 96:
		goto st_case_96
	case 97:
		goto st_case_97
	case 98:
		goto st_case_98
	case 99:
		goto st_case_99
	case 100:
		goto st_case_100
	case 101:
		goto st_case_101
	case 102:
		goto st_case_102
	case 103:
		goto st_case_103
	case 104:
		goto st_case_104
	case 105:
		goto st_case_105
	case 106:
		goto st_case_106
	case 107:
		goto st_case_107
	case 108:
		goto st_case_108
	case 109:
		goto st_case_109
	case 132:
		goto st_case_132
	case 110:
		goto st_case_110
	case 133:
		goto st_case_133
	case 134:
		goto st_case_134
	case 135:
		goto st_case_135
	case 136:
		goto st_case_136
	case 137:
		goto st_case_137
	case 138:
		goto st_case_138
	case 139:
		goto st_case_139
	case 140:
		goto st_case_140
	}
	goto st_out
tr14:
// line 62 "scanner/lex.rl"

 lex.te = ( lex.p)+1
{out.data = string(lex.data[lex.ts+1:lex.te-1]); tok=IPV4; {( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr25:
// line 63 "scanner/lex.rl"

 lex.te = ( lex.p)+1
{out.data = string(lex.data[lex.ts+1:lex.te-1]); tok=IPV6; {( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr33:
// line 90 "scanner/lex.rl"

( lex.p) = ( lex.te) - 1
{tok = COMMA;{( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr37:
// line 1 "NONE"

	switch  lex.act {
	case 0:
	{{goto st0 }}
	case 3:
	{( lex.p) = ( lex.te) - 1
out.val_int, _ = strconv.ParseInt(string(lex.data[lex.ts : lex.te]), 10, 64); tok = INT;{( lex.p)++;  lex.cs = 111; goto _out }}
	case 5:
	{( lex.p) = ( lex.te) - 1
out.val_int, _ = strconv.ParseInt(string(lex.data[lex.ts : lex.te]), 8, 64); tok = INT; {( lex.p)++;  lex.cs = 111; goto _out }}
	case 10:
	{( lex.p) = ( lex.te) - 1
tok = NULL; {( lex.p)++;  lex.cs = 111; goto _out }}
	case 11:
	{( lex.p) = ( lex.te) - 1
out.data = string(lex.data[lex.ts:lex.te]); tok = FLAG;{( lex.p)++;  lex.cs = 111; goto _out }}
	case 12:
	{( lex.p) = ( lex.te) - 1
out.data = string(lex.data[lex.ts:lex.te]); tok = IDENTIFIER;{( lex.p)++;  lex.cs = 111; goto _out }}
	case 35:
	{( lex.p) = ( lex.te) - 1
tok = COMMA;{( lex.p)++;  lex.cs = 111; goto _out }}
	}
	
	goto st111
tr51:
// line 68 "scanner/lex.rl"

 lex.te = ( lex.p)+1
{tok = UNFINISHED; {( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr52:
// line 87 "scanner/lex.rl"

 lex.te = ( lex.p)+1
{tok = ARROW; {( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr53:
// line 92 "scanner/lex.rl"

 lex.te = ( lex.p)+1
{{goto st139 }}
	goto st111
tr73:
// line 91 "scanner/lex.rl"

( lex.p) = ( lex.te) - 1
{out.data = string(lex.data[lex.ts:lex.te]); tok = DATETIME; {( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr75:
// line 60 "scanner/lex.rl"

( lex.p) = ( lex.te) - 1
{out.val_int, _ = strconv.ParseInt(string(lex.data[lex.ts : lex.te]), 8, 64); tok = INT; {( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr81:
// line 85 "scanner/lex.rl"

 lex.te = ( lex.p)+1
{tok = LSHIFT; {( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr96:
// line 69 "scanner/lex.rl"

( lex.p) = ( lex.te) - 1
{tok = RESUMED; {( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr97:
// line 69 "scanner/lex.rl"

 lex.te = ( lex.p)+1
{tok = RESUMED; {( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr120:
// line 86 "scanner/lex.rl"

 lex.te = ( lex.p)+1
{tok = RSHIFT; {( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr121:
// line 94 "scanner/lex.rl"

 lex.te = ( lex.p)+1

	goto st111
tr122:
// line 83 "scanner/lex.rl"

 lex.te = ( lex.p)+1
{tok = NOT;{( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr125:
// line 72 "scanner/lex.rl"

 lex.te = ( lex.p)+1
{tok = LPAREN;{( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr126:
// line 74 "scanner/lex.rl"

 lex.te = ( lex.p)+1
{tok = RPAREN;{( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr127:
// line 77 "scanner/lex.rl"

 lex.te = ( lex.p)+1
{tok = TIMES; {( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr134:
// line 81 "scanner/lex.rl"

 lex.te = ( lex.p)+1
{tok = COLON; {( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr138:
// line 93 "scanner/lex.rl"

 lex.te = ( lex.p)+1
{tok = QUESTION; {( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr139:
// line 73 "scanner/lex.rl"

 lex.te = ( lex.p)+1
{tok = AT; {( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr142:
// line 75 "scanner/lex.rl"

 lex.te = ( lex.p)+1
{tok = LBRACKET_SQUARE;{( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr143:
// line 76 "scanner/lex.rl"

 lex.te = ( lex.p)+1
{tok = RBRACKET_SQUARE;{( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr145:
// line 78 "scanner/lex.rl"

 lex.te = ( lex.p)+1
{tok = LBRACKET;{( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr147:
// line 79 "scanner/lex.rl"

 lex.te = ( lex.p)+1
{tok = RBRACKET;{( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr148:
// line 84 "scanner/lex.rl"

 lex.te = ( lex.p)+1
{tok = ONESCOMP; {( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr149:
// line 64 "scanner/lex.rl"

 lex.te = ( lex.p)
( lex.p)--
{out.data = ParseString(string(lex.data[lex.ts+1:lex.te-1])); tok = STRING_LITERAL;{( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr150:
// line 82 "scanner/lex.rl"

 lex.te = ( lex.p)
( lex.p)--
{tok = AND;{( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr151:
// line 89 "scanner/lex.rl"

 lex.te = ( lex.p)+1
{tok = LAND;{( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr152:
// line 59 "scanner/lex.rl"

 lex.te = ( lex.p)
( lex.p)--
{out.val_double, _ = strconv.ParseFloat(string(lex.data[lex.ts : lex.te]), 64); tok= DOUBLE; {( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr153:
// line 58 "scanner/lex.rl"

 lex.te = ( lex.p)
( lex.p)--
{out.val_int, _ = strconv.ParseInt(string(lex.data[lex.ts : lex.te]), 10, 64); tok = INT;{( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr155:
// line 90 "scanner/lex.rl"

 lex.te = ( lex.p)
( lex.p)--
{tok = COMMA;{( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr157:
// line 60 "scanner/lex.rl"

 lex.te = ( lex.p)
( lex.p)--
{out.val_int, _ = strconv.ParseInt(string(lex.data[lex.ts : lex.te]), 8, 64); tok = INT; {( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr164:
// line 91 "scanner/lex.rl"

 lex.te = ( lex.p)
( lex.p)--
{out.data = string(lex.data[lex.ts:lex.te]); tok = DATETIME; {( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr166:
// line 61 "scanner/lex.rl"

 lex.te = ( lex.p)
( lex.p)--
{out.val_uint, _ = strconv.ParseUint(string(lex.data[lex.ts:lex.te]), 0, 64); tok = UINT;{( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr170:
// line 69 "scanner/lex.rl"

 lex.te = ( lex.p)
( lex.p)--
{tok = RESUMED; {( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr172:
// line 70 "scanner/lex.rl"

 lex.te = ( lex.p)
( lex.p)--
{tok = EQUALS;{( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr173:
// line 71 "scanner/lex.rl"

 lex.te = ( lex.p)+1
{tok = LEQUAL; {( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr175:
// line 67 "scanner/lex.rl"

 lex.te = ( lex.p)
( lex.p)--
{out.data = string(lex.data[lex.ts:lex.te]); tok = IDENTIFIER;{( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr177:
// line 66 "scanner/lex.rl"

 lex.te = ( lex.p)
( lex.p)--
{out.data = string(lex.data[lex.ts:lex.te]); tok = FLAG;{( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr180:
// line 80 "scanner/lex.rl"

 lex.te = ( lex.p)
( lex.p)--
{tok = OR;{( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
tr181:
// line 88 "scanner/lex.rl"

 lex.te = ( lex.p)+1
{tok = LOR;{( lex.p)++;  lex.cs = 111; goto _out }}
	goto st111
	st111:
// line 1 "NONE"

 lex.ts = 0

// line 1 "NONE"

 lex.act = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof111
		}
	st_case_111:
// line 1 "NONE"

 lex.ts = ( lex.p)

// line 650 "scanner/lex.go"
		switch  lex.data[( lex.p)] {
		case 32:
			goto tr121
		case 33:
			goto tr122
		case 34:
			goto st1
		case 38:
			goto st113
		case 39:
			goto st25
		case 40:
			goto tr125
		case 41:
			goto tr126
		case 42:
			goto tr127
		case 43:
			goto st27
		case 44:
			goto tr129
		case 45:
			goto st46
		case 47:
			goto st47
		case 48:
			goto tr132
		case 58:
			goto tr134
		case 60:
			goto st72
		case 61:
			goto st132
		case 62:
			goto st110
		case 63:
			goto tr138
		case 64:
			goto tr139
		case 78:
			goto st135
		case 91:
			goto tr142
		case 93:
			goto tr143
		case 95:
			goto st25
		case 123:
			goto tr145
		case 124:
			goto st138
		case 125:
			goto tr147
		case 126:
			goto tr148
		}
		switch {
		case  lex.data[( lex.p)] < 49:
			if 9 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 13 {
				goto tr121
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto st134
				}
			case  lex.data[( lex.p)] >= 65:
				goto tr140
			}
		default:
			goto st127
		}
		goto st0
st_case_0:
	st0:
		 lex.cs = 0
		goto _out
	st1:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof1
		}
	st_case_1:
		switch  lex.data[( lex.p)] {
		case 34:
			goto st112
		case 42:
			goto st2
		case 47:
			goto st2
		case 58:
			goto st22
		case 92:
			goto st2
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto st3
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st2
			}
		default:
			goto st2
		}
		goto st0
	st112:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof112
		}
	st_case_112:
		switch  lex.data[( lex.p)] {
		case 39:
			goto st112
		case 46:
			goto st112
		}
		goto tr149
	st2:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof2
		}
	st_case_2:
		switch  lex.data[( lex.p)] {
		case 34:
			goto st112
		case 42:
			goto st2
		case 92:
			goto st2
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto st2
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st2
			}
		default:
			goto st2
		}
		goto st0
	st3:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof3
		}
	st_case_3:
		switch  lex.data[( lex.p)] {
		case 34:
			goto st112
		case 42:
			goto st2
		case 46:
			goto st4
		case 47:
			goto st2
		case 92:
			goto st2
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto st19
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st2
			}
		default:
			goto st2
		}
		goto st0
	st4:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof4
		}
	st_case_4:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st5
		}
		goto st0
	st5:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof5
		}
	st_case_5:
		if  lex.data[( lex.p)] == 46 {
			goto st6
		}
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st16
		}
		goto st0
	st6:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof6
		}
	st_case_6:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st7
		}
		goto st0
	st7:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof7
		}
	st_case_7:
		if  lex.data[( lex.p)] == 46 {
			goto st8
		}
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st13
		}
		goto st0
	st8:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof8
		}
	st_case_8:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st9
		}
		goto st0
	st9:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof9
		}
	st_case_9:
		if  lex.data[( lex.p)] == 34 {
			goto tr14
		}
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st10
		}
		goto st0
	st10:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof10
		}
	st_case_10:
		if  lex.data[( lex.p)] == 34 {
			goto tr14
		}
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st11
		}
		goto st0
	st11:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof11
		}
	st_case_11:
		if  lex.data[( lex.p)] == 34 {
			goto tr14
		}
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st12
		}
		goto st0
	st12:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof12
		}
	st_case_12:
		if  lex.data[( lex.p)] == 34 {
			goto tr14
		}
		goto st0
	st13:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof13
		}
	st_case_13:
		if  lex.data[( lex.p)] == 46 {
			goto st8
		}
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st14
		}
		goto st0
	st14:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof14
		}
	st_case_14:
		if  lex.data[( lex.p)] == 46 {
			goto st8
		}
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st15
		}
		goto st0
	st15:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof15
		}
	st_case_15:
		if  lex.data[( lex.p)] == 46 {
			goto st8
		}
		goto st0
	st16:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof16
		}
	st_case_16:
		if  lex.data[( lex.p)] == 46 {
			goto st6
		}
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st17
		}
		goto st0
	st17:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof17
		}
	st_case_17:
		if  lex.data[( lex.p)] == 46 {
			goto st6
		}
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st18
		}
		goto st0
	st18:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof18
		}
	st_case_18:
		if  lex.data[( lex.p)] == 46 {
			goto st6
		}
		goto st0
	st19:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof19
		}
	st_case_19:
		switch  lex.data[( lex.p)] {
		case 34:
			goto st112
		case 42:
			goto st2
		case 46:
			goto st4
		case 47:
			goto st2
		case 92:
			goto st2
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto st20
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st2
			}
		default:
			goto st2
		}
		goto st0
	st20:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof20
		}
	st_case_20:
		switch  lex.data[( lex.p)] {
		case 34:
			goto st112
		case 42:
			goto st2
		case 46:
			goto st4
		case 47:
			goto st2
		case 92:
			goto st2
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto st21
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st2
			}
		default:
			goto st2
		}
		goto st0
	st21:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof21
		}
	st_case_21:
		switch  lex.data[( lex.p)] {
		case 34:
			goto st112
		case 42:
			goto st2
		case 46:
			goto st4
		case 92:
			goto st2
		}
		switch {
		case  lex.data[( lex.p)] < 65:
			if 47 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto st2
			}
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st2
			}
		default:
			goto st2
		}
		goto st0
	st22:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof22
		}
	st_case_22:
		if  lex.data[( lex.p)] == 58 {
			goto st23
		}
		goto st0
	st23:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof23
		}
	st_case_23:
		if  lex.data[( lex.p)] == 34 {
			goto tr25
		}
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st24
		}
		goto st0
	st24:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof24
		}
	st_case_24:
		if  lex.data[( lex.p)] == 34 {
			goto tr25
		}
		goto st0
	st113:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof113
		}
	st_case_113:
		if  lex.data[( lex.p)] == 38 {
			goto tr151
		}
		goto tr150
	st25:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof25
		}
	st_case_25:
		switch  lex.data[( lex.p)] {
		case 39:
			goto st25
		case 95:
			goto st25
		}
		if 65 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 90 {
			goto st26
		}
		goto st0
	st26:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof26
		}
	st_case_26:
		switch  lex.data[( lex.p)] {
		case 39:
			goto tr29
		case 95:
			goto tr29
		}
		switch {
		case  lex.data[( lex.p)] > 57:
			if 65 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 90 {
				goto tr29
			}
		case  lex.data[( lex.p)] >= 48:
			goto tr29
		}
		goto st0
tr29:
// line 1 "NONE"

 lex.te = ( lex.p)+1

// line 66 "scanner/lex.rl"

 lex.act = 11;
	goto st114
tr179:
// line 1 "NONE"

 lex.te = ( lex.p)+1

// line 65 "scanner/lex.rl"

 lex.act = 10;
	goto st114
	st114:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof114
		}
	st_case_114:
// line 1174 "scanner/lex.go"
		switch  lex.data[( lex.p)] {
		case 39:
			goto tr29
		case 95:
			goto tr29
		}
		switch {
		case  lex.data[( lex.p)] > 57:
			if 65 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 90 {
				goto tr29
			}
		case  lex.data[( lex.p)] >= 48:
			goto tr29
		}
		goto tr37
	st27:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof27
		}
	st_case_27:
		if  lex.data[( lex.p)] == 48 {
			goto st28
		}
		if 49 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st116
		}
		goto st0
	st28:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof28
		}
	st_case_28:
		if  lex.data[( lex.p)] == 46 {
			goto st115
		}
		goto st0
	st115:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof115
		}
	st_case_115:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st115
		}
		goto tr152
	st116:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof116
		}
	st_case_116:
		if  lex.data[( lex.p)] == 46 {
			goto st115
		}
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st117
		}
		goto tr153
	st117:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof117
		}
	st_case_117:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st117
		}
		goto tr153
tr129:
// line 1 "NONE"

 lex.te = ( lex.p)+1

// line 90 "scanner/lex.rl"

 lex.act = 35;
	goto st118
	st118:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof118
		}
	st_case_118:
// line 1255 "scanner/lex.go"
		if  lex.data[( lex.p)] == 32 {
			goto st29
		}
		goto tr155
	st29:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof29
		}
	st_case_29:
		if  lex.data[( lex.p)] == 32 {
			goto st30
		}
		goto tr33
	st30:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof30
		}
	st_case_30:
		if  lex.data[( lex.p)] == 60 {
			goto st31
		}
		goto tr33
	st31:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof31
		}
	st_case_31:
		if  lex.data[( lex.p)] == 117 {
			goto st32
		}
		goto tr33
	st32:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof32
		}
	st_case_32:
		if  lex.data[( lex.p)] == 110 {
			goto st33
		}
		goto tr37
	st33:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof33
		}
	st_case_33:
		if  lex.data[( lex.p)] == 102 {
			goto st34
		}
		goto tr37
	st34:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof34
		}
	st_case_34:
		if  lex.data[( lex.p)] == 105 {
			goto st35
		}
		goto tr37
	st35:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof35
		}
	st_case_35:
		if  lex.data[( lex.p)] == 110 {
			goto st36
		}
		goto tr37
	st36:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof36
		}
	st_case_36:
		if  lex.data[( lex.p)] == 105 {
			goto st37
		}
		goto tr37
	st37:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof37
		}
	st_case_37:
		if  lex.data[( lex.p)] == 115 {
			goto st38
		}
		goto tr37
	st38:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof38
		}
	st_case_38:
		if  lex.data[( lex.p)] == 104 {
			goto st39
		}
		goto tr37
	st39:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof39
		}
	st_case_39:
		if  lex.data[( lex.p)] == 101 {
			goto st40
		}
		goto tr37
	st40:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof40
		}
	st_case_40:
		if  lex.data[( lex.p)] == 100 {
			goto st41
		}
		goto tr37
	st41:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof41
		}
	st_case_41:
		if  lex.data[( lex.p)] == 32 {
			goto st42
		}
		goto tr37
	st42:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof42
		}
	st_case_42:
		if  lex.data[( lex.p)] == 46 {
			goto st43
		}
		goto tr37
	st43:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof43
		}
	st_case_43:
		if  lex.data[( lex.p)] == 46 {
			goto st44
		}
		goto tr37
	st44:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof44
		}
	st_case_44:
		if  lex.data[( lex.p)] == 46 {
			goto st45
		}
		goto tr37
	st45:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof45
		}
	st_case_45:
		if  lex.data[( lex.p)] == 62 {
			goto tr51
		}
		goto tr37
	st46:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof46
		}
	st_case_46:
		switch  lex.data[( lex.p)] {
		case 48:
			goto st28
		case 62:
			goto tr52
		}
		if 49 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st116
		}
		goto st0
	st47:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof47
		}
	st_case_47:
		if  lex.data[( lex.p)] == 42 {
			goto tr53
		}
		goto st0
tr132:
// line 1 "NONE"

 lex.te = ( lex.p)+1

// line 60 "scanner/lex.rl"

 lex.act = 5;
	goto st119
	st119:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof119
		}
	st_case_119:
// line 1451 "scanner/lex.go"
		switch  lex.data[( lex.p)] {
		case 46:
			goto st115
		case 120:
			goto st71
		}
		switch {
		case  lex.data[( lex.p)] > 55:
			if 56 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto st70
			}
		case  lex.data[( lex.p)] >= 48:
			goto tr158
		}
		goto tr157
tr158:
// line 1 "NONE"

 lex.te = ( lex.p)+1

// line 60 "scanner/lex.rl"

 lex.act = 5;
	goto st120
	st120:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof120
		}
	st_case_120:
// line 1481 "scanner/lex.go"
		switch {
		case  lex.data[( lex.p)] > 55:
			if 56 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto st69
			}
		case  lex.data[( lex.p)] >= 48:
			goto tr161
		}
		goto tr157
tr161:
// line 1 "NONE"

 lex.te = ( lex.p)+1

// line 60 "scanner/lex.rl"

 lex.act = 5;
	goto st121
	st121:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof121
		}
	st_case_121:
// line 1505 "scanner/lex.go"
		switch {
		case  lex.data[( lex.p)] > 55:
			if 56 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto st68
			}
		case  lex.data[( lex.p)] >= 48:
			goto tr162
		}
		goto tr157
tr162:
// line 1 "NONE"

 lex.te = ( lex.p)+1

// line 60 "scanner/lex.rl"

 lex.act = 5;
	goto st122
	st122:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof122
		}
	st_case_122:
// line 1529 "scanner/lex.go"
		if  lex.data[( lex.p)] == 45 {
			goto st48
		}
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 55 {
			goto st125
		}
		goto tr157
	st48:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof48
		}
	st_case_48:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st49
		}
		goto tr37
	st49:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof49
		}
	st_case_49:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st50
		}
		goto tr37
	st50:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof50
		}
	st_case_50:
		if  lex.data[( lex.p)] == 45 {
			goto st51
		}
		goto tr37
	st51:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof51
		}
	st_case_51:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st52
		}
		goto tr37
	st52:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof52
		}
	st_case_52:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st53
		}
		goto tr37
	st53:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof53
		}
	st_case_53:
		if  lex.data[( lex.p)] == 84 {
			goto st54
		}
		goto tr37
	st54:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof54
		}
	st_case_54:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st55
		}
		goto tr37
	st55:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof55
		}
	st_case_55:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st56
		}
		goto tr37
	st56:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof56
		}
	st_case_56:
		if  lex.data[( lex.p)] == 58 {
			goto st57
		}
		goto tr37
	st57:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof57
		}
	st_case_57:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st58
		}
		goto tr37
	st58:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof58
		}
	st_case_58:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st59
		}
		goto tr37
	st59:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof59
		}
	st_case_59:
		if  lex.data[( lex.p)] == 58 {
			goto st60
		}
		goto tr37
	st60:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof60
		}
	st_case_60:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st61
		}
		goto tr37
	st61:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof61
		}
	st_case_61:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st62
		}
		goto tr37
	st62:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof62
		}
	st_case_62:
		if  lex.data[( lex.p)] == 43 {
			goto st63
		}
		goto tr37
	st63:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof63
		}
	st_case_63:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st64
		}
		goto tr37
	st64:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof64
		}
	st_case_64:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st65
		}
		goto tr37
	st65:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof65
		}
	st_case_65:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st66
		}
		goto tr37
	st66:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof66
		}
	st_case_66:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto tr72
		}
		goto tr37
tr72:
// line 1 "NONE"

 lex.te = ( lex.p)+1

	goto st123
	st123:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof123
		}
	st_case_123:
// line 1719 "scanner/lex.go"
		if  lex.data[( lex.p)] == 46 {
			goto st67
		}
		goto tr164
	st67:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof67
		}
	st_case_67:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st124
		}
		goto tr73
	st124:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof124
		}
	st_case_124:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st124
		}
		goto tr164
	st125:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof125
		}
	st_case_125:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 55 {
			goto st125
		}
		goto tr157
	st68:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof68
		}
	st_case_68:
		if  lex.data[( lex.p)] == 45 {
			goto st48
		}
		goto tr75
	st69:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof69
		}
	st_case_69:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st68
		}
		goto tr75
	st70:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof70
		}
	st_case_70:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st69
		}
		goto tr75
	st71:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof71
		}
	st_case_71:
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto st126
			}
		case  lex.data[( lex.p)] > 70:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 102 {
				goto st126
			}
		default:
			goto st126
		}
		goto tr75
	st126:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof126
		}
	st_case_126:
		switch {
		case  lex.data[( lex.p)] < 65:
			if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
				goto st126
			}
		case  lex.data[( lex.p)] > 70:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 102 {
				goto st126
			}
		default:
			goto st126
		}
		goto tr166
	st127:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof127
		}
	st_case_127:
		if  lex.data[( lex.p)] == 46 {
			goto st115
		}
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st128
		}
		goto tr153
	st128:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof128
		}
	st_case_128:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st129
		}
		goto tr153
	st129:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof129
		}
	st_case_129:
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto tr169
		}
		goto tr153
tr169:
// line 1 "NONE"

 lex.te = ( lex.p)+1

// line 58 "scanner/lex.rl"

 lex.act = 3;
	goto st130
	st130:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof130
		}
	st_case_130:
// line 1858 "scanner/lex.go"
		if  lex.data[( lex.p)] == 45 {
			goto st48
		}
		if 48 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 57 {
			goto st117
		}
		goto tr153
	st72:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof72
		}
	st_case_72:
		switch  lex.data[( lex.p)] {
		case 46:
			goto st73
		case 60:
			goto tr81
		case 117:
			goto st32
		}
		goto st0
	st73:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof73
		}
	st_case_73:
		if  lex.data[( lex.p)] == 46 {
			goto st74
		}
		goto st0
	st74:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof74
		}
	st_case_74:
		if  lex.data[( lex.p)] == 46 {
			goto st75
		}
		goto st0
	st75:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof75
		}
	st_case_75:
		if  lex.data[( lex.p)] == 32 {
			goto st76
		}
		goto st0
	st76:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof76
		}
	st_case_76:
		if  lex.data[( lex.p)] == 114 {
			goto st87
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st77
			}
		case  lex.data[( lex.p)] >= 65:
			goto st77
		}
		goto st0
	st77:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof77
		}
	st_case_77:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st78
		case 39:
			goto st77
		case 42:
			goto st77
		case 95:
			goto st77
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 45 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 46 {
				goto st77
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto st77
				}
			case  lex.data[( lex.p)] >= 65:
				goto st77
			}
		default:
			goto st77
		}
		goto st0
	st78:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof78
		}
	st_case_78:
		if  lex.data[( lex.p)] == 114 {
			goto st79
		}
		goto st0
	st79:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof79
		}
	st_case_79:
		if  lex.data[( lex.p)] == 101 {
			goto st80
		}
		goto st0
	st80:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof80
		}
	st_case_80:
		if  lex.data[( lex.p)] == 115 {
			goto st81
		}
		goto st0
	st81:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof81
		}
	st_case_81:
		if  lex.data[( lex.p)] == 117 {
			goto st82
		}
		goto st0
	st82:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof82
		}
	st_case_82:
		if  lex.data[( lex.p)] == 109 {
			goto st83
		}
		goto st0
	st83:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof83
		}
	st_case_83:
		if  lex.data[( lex.p)] == 101 {
			goto st84
		}
		goto st0
	st84:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof84
		}
	st_case_84:
		if  lex.data[( lex.p)] == 100 {
			goto st85
		}
		goto st0
	st85:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof85
		}
	st_case_85:
		if  lex.data[( lex.p)] == 62 {
			goto tr95
		}
		goto st0
tr95:
// line 1 "NONE"

 lex.te = ( lex.p)+1

	goto st131
	st131:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof131
		}
	st_case_131:
// line 2040 "scanner/lex.go"
		if  lex.data[( lex.p)] == 32 {
			goto st86
		}
		goto tr170
	st86:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof86
		}
	st_case_86:
		if  lex.data[( lex.p)] == 44 {
			goto tr97
		}
		goto tr96
	st87:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof87
		}
	st_case_87:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st78
		case 39:
			goto st77
		case 42:
			goto st77
		case 95:
			goto st77
		case 101:
			goto st88
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 45 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 46 {
				goto st77
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto st77
				}
			case  lex.data[( lex.p)] >= 65:
				goto st77
			}
		default:
			goto st77
		}
		goto st0
	st88:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof88
		}
	st_case_88:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st78
		case 39:
			goto st77
		case 42:
			goto st77
		case 95:
			goto st77
		case 115:
			goto st89
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 45 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 46 {
				goto st77
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto st77
				}
			case  lex.data[( lex.p)] >= 65:
				goto st77
			}
		default:
			goto st77
		}
		goto st0
	st89:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof89
		}
	st_case_89:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st78
		case 39:
			goto st77
		case 42:
			goto st77
		case 95:
			goto st77
		case 117:
			goto st90
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 45 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 46 {
				goto st77
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto st77
				}
			case  lex.data[( lex.p)] >= 65:
				goto st77
			}
		default:
			goto st77
		}
		goto st0
	st90:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof90
		}
	st_case_90:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st78
		case 39:
			goto st77
		case 42:
			goto st77
		case 95:
			goto st77
		case 109:
			goto st91
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 45 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 46 {
				goto st77
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto st77
				}
			case  lex.data[( lex.p)] >= 65:
				goto st77
			}
		default:
			goto st77
		}
		goto st0
	st91:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof91
		}
	st_case_91:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st78
		case 39:
			goto st77
		case 42:
			goto st77
		case 95:
			goto st77
		case 105:
			goto st92
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 45 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 46 {
				goto st77
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto st77
				}
			case  lex.data[( lex.p)] >= 65:
				goto st77
			}
		default:
			goto st77
		}
		goto st0
	st92:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof92
		}
	st_case_92:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st78
		case 39:
			goto st77
		case 42:
			goto st77
		case 95:
			goto st77
		case 110:
			goto st93
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 45 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 46 {
				goto st77
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto st77
				}
			case  lex.data[( lex.p)] >= 65:
				goto st77
			}
		default:
			goto st77
		}
		goto st0
	st93:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof93
		}
	st_case_93:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st78
		case 39:
			goto st77
		case 42:
			goto st77
		case 95:
			goto st77
		case 103:
			goto st94
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 45 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 46 {
				goto st77
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto st77
				}
			case  lex.data[( lex.p)] >= 65:
				goto st77
			}
		default:
			goto st77
		}
		goto st0
	st94:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof94
		}
	st_case_94:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st95
		case 39:
			goto st77
		case 42:
			goto st77
		case 95:
			goto st77
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 45 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 46 {
				goto st77
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto st77
				}
			case  lex.data[( lex.p)] >= 65:
				goto st77
			}
		default:
			goto st77
		}
		goto st0
	st95:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof95
		}
	st_case_95:
		if  lex.data[( lex.p)] == 114 {
			goto st103
		}
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st96
			}
		case  lex.data[( lex.p)] >= 65:
			goto st96
		}
		goto st0
	st96:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof96
		}
	st_case_96:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st97
		case 39:
			goto st96
		case 42:
			goto st96
		case 95:
			goto st96
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 45 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 46 {
				goto st96
			}
		case  lex.data[( lex.p)] > 57:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st96
			}
		default:
			goto st96
		}
		goto st0
	st97:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof97
		}
	st_case_97:
		switch {
		case  lex.data[( lex.p)] > 90:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st98
			}
		case  lex.data[( lex.p)] >= 65:
			goto st98
		}
		goto st0
	st98:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof98
		}
	st_case_98:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st99
		case 39:
			goto st98
		case 42:
			goto st98
		case 95:
			goto st98
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 45 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 46 {
				goto st98
			}
		case  lex.data[( lex.p)] > 57:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st98
			}
		default:
			goto st98
		}
		goto st0
	st99:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof99
		}
	st_case_99:
		if  lex.data[( lex.p)] == 46 {
			goto st100
		}
		goto st0
	st100:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof100
		}
	st_case_100:
		if  lex.data[( lex.p)] == 46 {
			goto st101
		}
		goto st0
	st101:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof101
		}
	st_case_101:
		if  lex.data[( lex.p)] == 46 {
			goto st102
		}
		goto st0
	st102:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof102
		}
	st_case_102:
		if  lex.data[( lex.p)] == 62 {
			goto tr97
		}
		goto st0
	st103:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof103
		}
	st_case_103:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st97
		case 39:
			goto st96
		case 42:
			goto st96
		case 95:
			goto st96
		case 101:
			goto st104
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 45 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 46 {
				goto st96
			}
		case  lex.data[( lex.p)] > 57:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st96
			}
		default:
			goto st96
		}
		goto st0
	st104:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof104
		}
	st_case_104:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st97
		case 39:
			goto st96
		case 42:
			goto st96
		case 95:
			goto st96
		case 115:
			goto st105
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 45 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 46 {
				goto st96
			}
		case  lex.data[( lex.p)] > 57:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st96
			}
		default:
			goto st96
		}
		goto st0
	st105:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof105
		}
	st_case_105:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st97
		case 39:
			goto st96
		case 42:
			goto st96
		case 95:
			goto st96
		case 117:
			goto st106
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 45 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 46 {
				goto st96
			}
		case  lex.data[( lex.p)] > 57:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st96
			}
		default:
			goto st96
		}
		goto st0
	st106:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof106
		}
	st_case_106:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st97
		case 39:
			goto st96
		case 42:
			goto st96
		case 95:
			goto st96
		case 109:
			goto st107
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 45 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 46 {
				goto st96
			}
		case  lex.data[( lex.p)] > 57:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st96
			}
		default:
			goto st96
		}
		goto st0
	st107:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof107
		}
	st_case_107:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st97
		case 39:
			goto st96
		case 42:
			goto st96
		case 95:
			goto st96
		case 101:
			goto st108
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 45 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 46 {
				goto st96
			}
		case  lex.data[( lex.p)] > 57:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st96
			}
		default:
			goto st96
		}
		goto st0
	st108:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof108
		}
	st_case_108:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st97
		case 39:
			goto st96
		case 42:
			goto st96
		case 95:
			goto st96
		case 100:
			goto st109
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 45 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 46 {
				goto st96
			}
		case  lex.data[( lex.p)] > 57:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st96
			}
		default:
			goto st96
		}
		goto st0
	st109:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof109
		}
	st_case_109:
		switch  lex.data[( lex.p)] {
		case 32:
			goto st97
		case 39:
			goto st96
		case 42:
			goto st96
		case 62:
			goto tr95
		case 95:
			goto st96
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 45 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 46 {
				goto st96
			}
		case  lex.data[( lex.p)] > 57:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st96
			}
		default:
			goto st96
		}
		goto st0
	st132:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof132
		}
	st_case_132:
		if  lex.data[( lex.p)] == 61 {
			goto tr173
		}
		goto tr172
	st110:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof110
		}
	st_case_110:
		if  lex.data[( lex.p)] == 62 {
			goto tr120
		}
		goto st0
tr174:
// line 1 "NONE"

 lex.te = ( lex.p)+1

// line 66 "scanner/lex.rl"

 lex.act = 11;
	goto st133
tr140:
// line 1 "NONE"

 lex.te = ( lex.p)+1

// line 67 "scanner/lex.rl"

 lex.act = 12;
	goto st133
	st133:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof133
		}
	st_case_133:
// line 2706 "scanner/lex.go"
		switch  lex.data[( lex.p)] {
		case 39:
			goto tr174
		case 42:
			goto st134
		case 95:
			goto tr174
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 45 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 46 {
				goto st134
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto st134
				}
			case  lex.data[( lex.p)] >= 65:
				goto tr29
			}
		default:
			goto tr174
		}
		goto tr37
	st134:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof134
		}
	st_case_134:
		switch  lex.data[( lex.p)] {
		case 39:
			goto st134
		case 42:
			goto st134
		case 95:
			goto st134
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 45 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 46 {
				goto st134
			}
		case  lex.data[( lex.p)] > 57:
			if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
				goto st134
			}
		default:
			goto st134
		}
		goto tr175
	st135:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof135
		}
	st_case_135:
		switch  lex.data[( lex.p)] {
		case 39:
			goto tr174
		case 42:
			goto st134
		case 85:
			goto st136
		case 95:
			goto tr174
		}
		switch {
		case  lex.data[( lex.p)] < 48:
			if 45 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 46 {
				goto st134
			}
		case  lex.data[( lex.p)] > 57:
			switch {
			case  lex.data[( lex.p)] > 90:
				if 97 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 122 {
					goto st134
				}
			case  lex.data[( lex.p)] >= 65:
				goto tr29
			}
		default:
			goto tr174
		}
		goto tr175
	st136:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof136
		}
	st_case_136:
		switch  lex.data[( lex.p)] {
		case 39:
			goto tr29
		case 76:
			goto st137
		case 95:
			goto tr29
		}
		switch {
		case  lex.data[( lex.p)] > 57:
			if 65 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 90 {
				goto tr29
			}
		case  lex.data[( lex.p)] >= 48:
			goto tr29
		}
		goto tr177
	st137:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof137
		}
	st_case_137:
		switch  lex.data[( lex.p)] {
		case 39:
			goto tr29
		case 76:
			goto tr179
		case 95:
			goto tr29
		}
		switch {
		case  lex.data[( lex.p)] > 57:
			if 65 <=  lex.data[( lex.p)] &&  lex.data[( lex.p)] <= 90 {
				goto tr29
			}
		case  lex.data[( lex.p)] >= 48:
			goto tr29
		}
		goto tr177
	st138:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof138
		}
	st_case_138:
		if  lex.data[( lex.p)] == 124 {
			goto tr181
		}
		goto tr180
tr182:
// line 53 "scanner/lex.rl"

 lex.te = ( lex.p)+1

	goto st139
tr184:
// line 53 "scanner/lex.rl"

 lex.te = ( lex.p)
( lex.p)--

	goto st139
tr185:
// line 54 "scanner/lex.rl"

 lex.te = ( lex.p)+1
{{goto st111 }}
	goto st139
	st139:
// line 1 "NONE"

 lex.ts = 0

		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof139
		}
	st_case_139:
// line 1 "NONE"

 lex.ts = ( lex.p)

// line 2877 "scanner/lex.go"
		if  lex.data[( lex.p)] == 42 {
			goto st140
		}
		goto tr182
	st140:
		if ( lex.p)++; ( lex.p) == ( lex.pe) {
			goto _test_eof140
		}
	st_case_140:
		if  lex.data[( lex.p)] == 47 {
			goto tr185
		}
		goto tr184
	st_out:
	_test_eof111:  lex.cs = 111; goto _test_eof
	_test_eof1:  lex.cs = 1; goto _test_eof
	_test_eof112:  lex.cs = 112; goto _test_eof
	_test_eof2:  lex.cs = 2; goto _test_eof
	_test_eof3:  lex.cs = 3; goto _test_eof
	_test_eof4:  lex.cs = 4; goto _test_eof
	_test_eof5:  lex.cs = 5; goto _test_eof
	_test_eof6:  lex.cs = 6; goto _test_eof
	_test_eof7:  lex.cs = 7; goto _test_eof
	_test_eof8:  lex.cs = 8; goto _test_eof
	_test_eof9:  lex.cs = 9; goto _test_eof
	_test_eof10:  lex.cs = 10; goto _test_eof
	_test_eof11:  lex.cs = 11; goto _test_eof
	_test_eof12:  lex.cs = 12; goto _test_eof
	_test_eof13:  lex.cs = 13; goto _test_eof
	_test_eof14:  lex.cs = 14; goto _test_eof
	_test_eof15:  lex.cs = 15; goto _test_eof
	_test_eof16:  lex.cs = 16; goto _test_eof
	_test_eof17:  lex.cs = 17; goto _test_eof
	_test_eof18:  lex.cs = 18; goto _test_eof
	_test_eof19:  lex.cs = 19; goto _test_eof
	_test_eof20:  lex.cs = 20; goto _test_eof
	_test_eof21:  lex.cs = 21; goto _test_eof
	_test_eof22:  lex.cs = 22; goto _test_eof
	_test_eof23:  lex.cs = 23; goto _test_eof
	_test_eof24:  lex.cs = 24; goto _test_eof
	_test_eof113:  lex.cs = 113; goto _test_eof
	_test_eof25:  lex.cs = 25; goto _test_eof
	_test_eof26:  lex.cs = 26; goto _test_eof
	_test_eof114:  lex.cs = 114; goto _test_eof
	_test_eof27:  lex.cs = 27; goto _test_eof
	_test_eof28:  lex.cs = 28; goto _test_eof
	_test_eof115:  lex.cs = 115; goto _test_eof
	_test_eof116:  lex.cs = 116; goto _test_eof
	_test_eof117:  lex.cs = 117; goto _test_eof
	_test_eof118:  lex.cs = 118; goto _test_eof
	_test_eof29:  lex.cs = 29; goto _test_eof
	_test_eof30:  lex.cs = 30; goto _test_eof
	_test_eof31:  lex.cs = 31; goto _test_eof
	_test_eof32:  lex.cs = 32; goto _test_eof
	_test_eof33:  lex.cs = 33; goto _test_eof
	_test_eof34:  lex.cs = 34; goto _test_eof
	_test_eof35:  lex.cs = 35; goto _test_eof
	_test_eof36:  lex.cs = 36; goto _test_eof
	_test_eof37:  lex.cs = 37; goto _test_eof
	_test_eof38:  lex.cs = 38; goto _test_eof
	_test_eof39:  lex.cs = 39; goto _test_eof
	_test_eof40:  lex.cs = 40; goto _test_eof
	_test_eof41:  lex.cs = 41; goto _test_eof
	_test_eof42:  lex.cs = 42; goto _test_eof
	_test_eof43:  lex.cs = 43; goto _test_eof
	_test_eof44:  lex.cs = 44; goto _test_eof
	_test_eof45:  lex.cs = 45; goto _test_eof
	_test_eof46:  lex.cs = 46; goto _test_eof
	_test_eof47:  lex.cs = 47; goto _test_eof
	_test_eof119:  lex.cs = 119; goto _test_eof
	_test_eof120:  lex.cs = 120; goto _test_eof
	_test_eof121:  lex.cs = 121; goto _test_eof
	_test_eof122:  lex.cs = 122; goto _test_eof
	_test_eof48:  lex.cs = 48; goto _test_eof
	_test_eof49:  lex.cs = 49; goto _test_eof
	_test_eof50:  lex.cs = 50; goto _test_eof
	_test_eof51:  lex.cs = 51; goto _test_eof
	_test_eof52:  lex.cs = 52; goto _test_eof
	_test_eof53:  lex.cs = 53; goto _test_eof
	_test_eof54:  lex.cs = 54; goto _test_eof
	_test_eof55:  lex.cs = 55; goto _test_eof
	_test_eof56:  lex.cs = 56; goto _test_eof
	_test_eof57:  lex.cs = 57; goto _test_eof
	_test_eof58:  lex.cs = 58; goto _test_eof
	_test_eof59:  lex.cs = 59; goto _test_eof
	_test_eof60:  lex.cs = 60; goto _test_eof
	_test_eof61:  lex.cs = 61; goto _test_eof
	_test_eof62:  lex.cs = 62; goto _test_eof
	_test_eof63:  lex.cs = 63; goto _test_eof
	_test_eof64:  lex.cs = 64; goto _test_eof
	_test_eof65:  lex.cs = 65; goto _test_eof
	_test_eof66:  lex.cs = 66; goto _test_eof
	_test_eof123:  lex.cs = 123; goto _test_eof
	_test_eof67:  lex.cs = 67; goto _test_eof
	_test_eof124:  lex.cs = 124; goto _test_eof
	_test_eof125:  lex.cs = 125; goto _test_eof
	_test_eof68:  lex.cs = 68; goto _test_eof
	_test_eof69:  lex.cs = 69; goto _test_eof
	_test_eof70:  lex.cs = 70; goto _test_eof
	_test_eof71:  lex.cs = 71; goto _test_eof
	_test_eof126:  lex.cs = 126; goto _test_eof
	_test_eof127:  lex.cs = 127; goto _test_eof
	_test_eof128:  lex.cs = 128; goto _test_eof
	_test_eof129:  lex.cs = 129; goto _test_eof
	_test_eof130:  lex.cs = 130; goto _test_eof
	_test_eof72:  lex.cs = 72; goto _test_eof
	_test_eof73:  lex.cs = 73; goto _test_eof
	_test_eof74:  lex.cs = 74; goto _test_eof
	_test_eof75:  lex.cs = 75; goto _test_eof
	_test_eof76:  lex.cs = 76; goto _test_eof
	_test_eof77:  lex.cs = 77; goto _test_eof
	_test_eof78:  lex.cs = 78; goto _test_eof
	_test_eof79:  lex.cs = 79; goto _test_eof
	_test_eof80:  lex.cs = 80; goto _test_eof
	_test_eof81:  lex.cs = 81; goto _test_eof
	_test_eof82:  lex.cs = 82; goto _test_eof
	_test_eof83:  lex.cs = 83; goto _test_eof
	_test_eof84:  lex.cs = 84; goto _test_eof
	_test_eof85:  lex.cs = 85; goto _test_eof
	_test_eof131:  lex.cs = 131; goto _test_eof
	_test_eof86:  lex.cs = 86; goto _test_eof
	_test_eof87:  lex.cs = 87; goto _test_eof
	_test_eof88:  lex.cs = 88; goto _test_eof
	_test_eof89:  lex.cs = 89; goto _test_eof
	_test_eof90:  lex.cs = 90; goto _test_eof
	_test_eof91:  lex.cs = 91; goto _test_eof
	_test_eof92:  lex.cs = 92; goto _test_eof
	_test_eof93:  lex.cs = 93; goto _test_eof
	_test_eof94:  lex.cs = 94; goto _test_eof
	_test_eof95:  lex.cs = 95; goto _test_eof
	_test_eof96:  lex.cs = 96; goto _test_eof
	_test_eof97:  lex.cs = 97; goto _test_eof
	_test_eof98:  lex.cs = 98; goto _test_eof
	_test_eof99:  lex.cs = 99; goto _test_eof
	_test_eof100:  lex.cs = 100; goto _test_eof
	_test_eof101:  lex.cs = 101; goto _test_eof
	_test_eof102:  lex.cs = 102; goto _test_eof
	_test_eof103:  lex.cs = 103; goto _test_eof
	_test_eof104:  lex.cs = 104; goto _test_eof
	_test_eof105:  lex.cs = 105; goto _test_eof
	_test_eof106:  lex.cs = 106; goto _test_eof
	_test_eof107:  lex.cs = 107; goto _test_eof
	_test_eof108:  lex.cs = 108; goto _test_eof
	_test_eof109:  lex.cs = 109; goto _test_eof
	_test_eof132:  lex.cs = 132; goto _test_eof
	_test_eof110:  lex.cs = 110; goto _test_eof
	_test_eof133:  lex.cs = 133; goto _test_eof
	_test_eof134:  lex.cs = 134; goto _test_eof
	_test_eof135:  lex.cs = 135; goto _test_eof
	_test_eof136:  lex.cs = 136; goto _test_eof
	_test_eof137:  lex.cs = 137; goto _test_eof
	_test_eof138:  lex.cs = 138; goto _test_eof
	_test_eof139:  lex.cs = 139; goto _test_eof
	_test_eof140:  lex.cs = 140; goto _test_eof

	_test_eof: {}
	if ( lex.p) == eof {
		switch  lex.cs {
		case 112:
			goto tr149
		case 113:
			goto tr150
		case 114:
			goto tr37
		case 115:
			goto tr152
		case 116:
			goto tr153
		case 117:
			goto tr153
		case 118:
			goto tr155
		case 29:
			goto tr33
		case 30:
			goto tr33
		case 31:
			goto tr33
		case 32:
			goto tr37
		case 33:
			goto tr37
		case 34:
			goto tr37
		case 35:
			goto tr37
		case 36:
			goto tr37
		case 37:
			goto tr37
		case 38:
			goto tr37
		case 39:
			goto tr37
		case 40:
			goto tr37
		case 41:
			goto tr37
		case 42:
			goto tr37
		case 43:
			goto tr37
		case 44:
			goto tr37
		case 45:
			goto tr37
		case 119:
			goto tr157
		case 120:
			goto tr157
		case 121:
			goto tr157
		case 122:
			goto tr157
		case 48:
			goto tr37
		case 49:
			goto tr37
		case 50:
			goto tr37
		case 51:
			goto tr37
		case 52:
			goto tr37
		case 53:
			goto tr37
		case 54:
			goto tr37
		case 55:
			goto tr37
		case 56:
			goto tr37
		case 57:
			goto tr37
		case 58:
			goto tr37
		case 59:
			goto tr37
		case 60:
			goto tr37
		case 61:
			goto tr37
		case 62:
			goto tr37
		case 63:
			goto tr37
		case 64:
			goto tr37
		case 65:
			goto tr37
		case 66:
			goto tr37
		case 123:
			goto tr164
		case 67:
			goto tr73
		case 124:
			goto tr164
		case 125:
			goto tr157
		case 68:
			goto tr75
		case 69:
			goto tr75
		case 70:
			goto tr75
		case 71:
			goto tr75
		case 126:
			goto tr166
		case 127:
			goto tr153
		case 128:
			goto tr153
		case 129:
			goto tr153
		case 130:
			goto tr153
		case 131:
			goto tr170
		case 86:
			goto tr96
		case 132:
			goto tr172
		case 133:
			goto tr37
		case 134:
			goto tr175
		case 135:
			goto tr175
		case 136:
			goto tr177
		case 137:
			goto tr177
		case 138:
			goto tr180
		case 140:
			goto tr184
		}
	}

	_out: {}
	}

// line 98 "scanner/lex.rl"


    return tok;
}

func (lex *lexer) Error(e string) {
    fmt.Println("error:", e)
}

func ParseString(s string) string{
	var decoded []byte
	var err error
	var strippedStr string
	strippedStr = strings.Replace(s, `\x`, "", -1)
	strippedStr = strings.Replace(strippedStr, ".", "", -1)
	strippedStr = strings.Replace(strippedStr, `"`, "", -1)
	if len(strippedStr) % 2 > 0 {
	    strippedStr += "0"
	}
	if decoded, err = hex.DecodeString(strippedStr); err != nil {
		panic(fmt.Sprintf("Failed to decode string: %s, with error: %s\n", s, err.Error()))
	}
	decoded = append(decoded, '\x00')
	return string(decoded)
}
