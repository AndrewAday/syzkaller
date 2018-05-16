%{
package scanner

import (
    //"fmt"
    types "github.com/google/syzkaller/tools/moonshine/strace_types"
)
%}

%start syscall

%union {
    data string
    val_int int64
    val_double float64
    val_uint uint64
    val_field *types.Field
    val_call *types.Call
    val_int_type *types.IntType
    val_identifiers []*types.BufferType
    val_buf_type *types.BufferType
    val_struct_type *types.StructType
    val_array_type *types.ArrayType
    val_pointer_type *types.PointerType
    val_flag_type *types.FlagType
    val_expr_type *types.Expression
    val_type types.Type
    val_ipv4_type *types.Ipv4Type
    val_types []types.Type
    val_syscall *types.Syscall
}

%token <data> STRING_LITERAL IPV4 IDENTIFIER FLAG DATETIME SIGNAL_PLUS SIGNAL_MINUS
%token <val_int> INT
%token <val_uint> UINT
%token <val_double> DOUBLE
%type <val_field> field_type
%type <val_identifiers> identifiers
%type <val_int_type> int_type
%type <val_buf_type> buf_type
%type <val_struct_type> struct_type
%type <val_array_type> array_type
%type <val_flag_type> flag_type
%type <val_expr_type> expr_type
%type <val_call> call_type
%type <val_type> type
%type <val_pointer_type> pointer_type
%type <val_ipv4_type> ipv4_type
%type <val_types> types
%type <val_syscall> syscall

%token STRING_LITERAL IPV4 IDENTIFIER FLAG INT UINT QUESTION DOUBLE OR AND LOR TIMES LAND NOT ONESCOMP LSHIFT RSHIFT
%token COMMA LBRACKET RBRACKET LBRACKET_SQUARE RBRACKET_SQUARE LPAREN RPAREN EQUALS
%token UNFINISHED UNFINISHED_W_COMMA RESUMED
%token SIGNAL_PLUS SIGNAL_MINUS NULL

%nonassoc FLAG
%nonassoc NOFLAG

%%
syscall:
    INT IDENTIFIER LPAREN UNFINISHED %prec NOFLAG { $$ = types.NewSyscall($1, $2, nil, int64(-1), true, false);
                                                        Stracelex.(*lexer).result = $$ }
    | INT IDENTIFIER LPAREN types UNFINISHED %prec NOFLAG { $$ = types.NewSyscall($1, $2, $4, int64(-1), true, false);
                                                        Stracelex.(*lexer).result = $$ }
    | INT IDENTIFIER LPAREN types UNFINISHED_W_COMMA %prec NOFLAG { $$ = types.NewSyscall($1, $2, $4, int64(-1), true, false);
                                                        Stracelex.(*lexer).result = $$ }
    | INT RESUMED RPAREN EQUALS INT %prec NOFLAG { $$ = types.NewSyscall($1, "tmp", nil, int64($5), false, true);
                                                        Stracelex.(*lexer).result = $$ }
    | INT RESUMED RPAREN EQUALS UINT %prec NOFLAG { $$ = types.NewSyscall($1, "tmp", nil, int64($5), false, true);
                                                        Stracelex.(*lexer).result = $$ }
    | INT RESUMED RPAREN EQUALS INT LPAREN FLAG RPAREN { $$ = types.NewSyscall($1, "tmp", nil, int64($5), false, true);
                                                        Stracelex.(*lexer).result = $$ }
    | INT RESUMED RPAREN EQUALS UINT LPAREN FLAG RPAREN { $$ = types.NewSyscall($1, "tmp", nil, int64($5), false, true);
                                                        Stracelex.(*lexer).result = $$ }
    | INT RESUMED types RPAREN EQUALS INT %prec NOFLAG { $$ = types.NewSyscall($1, "tmp", $3, int64($6), false, true);
                                                        Stracelex.(*lexer).result = $$ }
    | INT RESUMED types RPAREN EQUALS UINT %prec NOFLAG { $$ = types.NewSyscall($1, "tmp", $3, int64($6), false, true);
                                                        Stracelex.(*lexer).result = $$ }
    | INT RESUMED types RPAREN EQUALS INT LPAREN FLAG RPAREN { $$ = types.NewSyscall($1, "tmp", $3, int64($6), false, true);
                                                        Stracelex.(*lexer).result = $$ }
    | INT RESUMED types RPAREN EQUALS UINT LPAREN FLAG RPAREN { $$ = types.NewSyscall($1, "tmp", $3, int64($6), false, true);
                                                        Stracelex.(*lexer).result = $$ }
    | INT IDENTIFIER LPAREN RPAREN EQUALS INT %prec NOFLAG { $$ = types.NewSyscall($1, $2, nil, $6, false, false);
                                                            Stracelex.(*lexer).result = $$;}
    | INT IDENTIFIER LPAREN types RPAREN EQUALS INT %prec NOFLAG{
                                                        $$ = types.NewSyscall($1, $2, $4, $7, false, false);
                                                        Stracelex.(*lexer).result = $$;}
    | INT IDENTIFIER LPAREN types RPAREN EQUALS UINT %prec NOFLAG {
                                                        $$ = types.NewSyscall($1, $2, $4, int64($7), false, false);
                                                        Stracelex.(*lexer).result = $$;}
    | INT IDENTIFIER LPAREN types RPAREN EQUALS QUESTION %prec NOFLAG {
                                                            $$ = types.NewSyscall($1, $2, $4, -1, false, false);
                                                            Stracelex.(*lexer).result = $$;}
    | INT IDENTIFIER LPAREN types RPAREN EQUALS INT FLAG LPAREN identifiers RPAREN {
                                                              $$ = types.NewSyscall($1, $2, $4, $7, false, false);
                                                              Stracelex.(*lexer).result = $$;}
    | INT IDENTIFIER LPAREN types RPAREN EQUALS UINT FLAG LPAREN identifiers RPAREN {
                                                              $$ = types.NewSyscall($1, $2, $4, int64($7), false, false);
                                                              Stracelex.(*lexer).result = $$;}
    | INT IDENTIFIER LPAREN types RPAREN EQUALS INT LPAREN flag_type RPAREN {
                                                                  $$ = types.NewSyscall($1, $2, $4, $7, false, false);
                                                                  Stracelex.(*lexer).result = $$;}
    | INT IDENTIFIER LPAREN types RPAREN EQUALS UINT LPAREN flag_type RPAREN {
                                                                  $$ = types.NewSyscall($1, $2, $4, int64($7), false, false);
                                                                  Stracelex.(*lexer).result = $$;}



types:
    type {types := make([]types.Type, 0); types = append(types, $1); $$ = types;}
    | type COMMA types {$3 = append([]types.Type{$1}, $3...); $$ = $3;}

type:
    buf_type {$$ = $1}
    | field_type {$$ = $1}
    | expr_type {$$ = $1}
    | pointer_type {$$ = $1}
    | array_type {$$ = $1}
    | struct_type {$$ = $1}
    | call_type {$$ = $1}
    | ipv4_type {$$ = $1}

call_type:
    IDENTIFIER LPAREN types RPAREN {$$ = types.NewCallType($1, $3)}

pointer_type:
    AND UINT EQUALS type {$$ = types.NewPointerType($2, $4)}
    | NULL {$$ = types.NullPointer()}

array_type:
    LBRACKET_SQUARE types RBRACKET_SQUARE {arr := types.NewArrayType($2); $$ = arr}
    | LBRACKET_SQUARE RBRACKET_SQUARE {arr := types.NewArrayType(nil); $$ = arr}

struct_type:
    LBRACKET types RBRACKET {$$ = types.NewStructType($2)}

field_type:
    IDENTIFIER EQUALS type {$$ = types.NewField($1, $3);}

buf_type:
    STRING_LITERAL {$$ = types.NewBufferType($1)}
    | DATETIME {$$ = types.NewBufferType($1)}

expr_type:
    flag_type {$$ = types.NewExpression($1);}
    | int_type {$$ = types.NewExpression($1);}
    | LPAREN expr_type RPAREN {$$ = types.NewExpression($2);}
    | flag_type OR expr_type {$$ = types.NewExpression(types.NewBinop(types.NewExpression($1), types.OR, $3));}
    | int_type OR expr_type {$$ = types.NewExpression(types.NewBinop(types.NewExpression($1), types.OR, $3));}
    | flag_type AND expr_type {$$ = types.NewExpression(types.NewBinop(types.NewExpression($1), types.AND, $3));}
    | int_type AND expr_type {$$ = types.NewExpression(types.NewBinop(types.NewExpression($1), types.AND, $3));}
    | flag_type LSHIFT expr_type {$$ = types.NewExpression(types.NewBinop(types.NewExpression($1), types.LSHIFT, $3));}
    | int_type LSHIFT expr_type {$$ = types.NewExpression(types.NewBinop(types.NewExpression($1), types.LSHIFT, $3));}
    | flag_type RSHIFT expr_type {$$ = types.NewExpression(types.NewBinop(types.NewExpression($1), types.RSHIFT, $3));}
    | int_type RSHIFT expr_type {$$ = types.NewExpression(types.NewBinop(types.NewExpression($1), types.RSHIFT, $3));}
    | int_type TIMES expr_type {$$ = types.NewExpression(types.NewBinop(types.NewExpression($1), types.TIMES, $3));}
    | ONESCOMP expr_type {$$ = types.NewExpression(types.NewUnop(types.NewExpression($2), types.ONESCOMP));}

int_type:
      INT {$$ = types.NewIntType($1)}
    | UINT {$$ = types.NewIntType(int64($1))}

flag_type:
      FLAG {$$ = types.NewFlagType($1)}

ipv4_type:
    IPV4 {$$ = types.NewIpv4Type($1)}

identifiers:
    IDENTIFIER {ids := make([]*types.BufferType, 0); ids = append(ids, types.NewBufferType($1)); $$ = ids}
    | IDENTIFIER identifiers {$2 = append($2, types.NewBufferType($1)); $$ = $2}

