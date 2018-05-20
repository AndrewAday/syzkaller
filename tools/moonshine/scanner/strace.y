%{
package scanner

import (
    "fmt"
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
    val_macro *types.Macro
    val_int_type *types.IntType
    val_identifiers []*types.BufferType
    val_buf_type *types.BufferType
    val_struct_type *types.StructType
    val_dynamic_type *types.DynamicType
    val_array_type *types.ArrayType
    val_pointer_type *types.PointerType
    val_flag_type *types.FlagType
    val_expr_type *types.Expression
    val_type types.Type
    val_ipv4_type *types.Ipv4Type
    val_types []types.Type
    val_parenthetical *types.Parenthetical
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
%type <val_dynamic_type> dynamic_type
%type <val_array_type> array_type
%type <val_flag_type> flag_type
%type <val_expr_type> expr_type
%type <val_call> call_type
%type <val_parenthetical> parenthetical
%type <val_macro> macro_type
%type <val_type> type
%type <val_pointer_type> pointer_type
%type <val_ipv4_type> ipv4_type
%type <val_types> types
%type <val_syscall> syscall

%token STRING_LITERAL IPV4 IDENTIFIER FLAG INT UINT QUESTION DOUBLE ARROW
%token OR AND LOR TIMES LAND LEQUAL ONESCOMP LSHIFT RSHIFT TIMES NOT
%token COMMA LBRACKET RBRACKET LBRACKET_SQUARE RBRACKET_SQUARE LPAREN RPAREN EQUALS
%token UNFINISHED RESUMED
%token SIGNAL_PLUS SIGNAL_MINUS NULL

%nonassoc FLAG
%nonassoc NOFLAG

%%
syscall:
    INT IDENTIFIER LPAREN UNFINISHED %prec NOFLAG { $$ = types.NewSyscall($1, $2, nil, int64(-1), true, false);
                                                        Stracelex.(*lexer).result = $$ }
    | INT IDENTIFIER LPAREN types UNFINISHED %prec NOFLAG { $$ = types.NewSyscall($1, $2, $4, int64(-1), true, false);
                                                        Stracelex.(*lexer).result = $$ }
    | INT RESUMED UNFINISHED RPAREN EQUALS QUESTION %prec NOFLAG
        {
            $$ = types.NewSyscall($1, "tmp", nil, -1, true, true);
            Stracelex.(*lexer).result = $$;
        }
    | INT IDENTIFIER LPAREN RESUMED RPAREN EQUALS INT %prec NOFLAG
        {
            $$ = types.NewSyscall($1, $2, nil, int64($7), false, false);
            Stracelex.(*lexer).result = $$;
        }
    | INT RESUMED RPAREN EQUALS INT %prec NOFLAG { $$ = types.NewSyscall($1, "tmp", nil, int64($5), false, true);
                                                        Stracelex.(*lexer).result = $$ }
    | INT RESUMED RPAREN EQUALS UINT %prec NOFLAG { $$ = types.NewSyscall($1, "tmp", nil, int64($5), false, true);
                                                        Stracelex.(*lexer).result = $$ }
    | INT RESUMED RPAREN EQUALS QUESTION %prec NOFLAG { $$ = types.NewSyscall($1, "tmp", nil, -1, false, true);
                                                              Stracelex.(*lexer).result = $$ }
    | INT RESUMED RPAREN EQUALS INT LPAREN FLAG RPAREN { $$ = types.NewSyscall($1, "tmp", nil, int64($5), false, true);
                                                        Stracelex.(*lexer).result = $$ }
    | INT RESUMED RPAREN EQUALS UINT LPAREN FLAG RPAREN { $$ = types.NewSyscall($1, "tmp", nil, int64($5), false, true);
                                                        Stracelex.(*lexer).result = $$ }
    | INT RESUMED types RPAREN EQUALS INT %prec NOFLAG { $$ = types.NewSyscall($1, "tmp", $3, int64($6), false, true);
                                                        Stracelex.(*lexer).result = $$ }
    | INT RESUMED types RPAREN EQUALS UINT %prec NOFLAG { $$ = types.NewSyscall($1, "tmp", $3, int64($6), false, true);
                                                        Stracelex.(*lexer).result = $$ }
    | INT RESUMED types RPAREN EQUALS QUESTION %prec NOFLAG { $$ = types.NewSyscall($1, "tmp", $3, -1, false, true);
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
    | INT IDENTIFIER LPAREN types RPAREN EQUALS INT FLAG LPAREN parenthetical RPAREN {
                                                              $$ = types.NewSyscall($1, $2, $4, $7, false, false);
                                                              Stracelex.(*lexer).result = $$;}
    | INT IDENTIFIER LPAREN types RPAREN EQUALS UINT FLAG LPAREN parenthetical RPAREN {
                                                              $$ = types.NewSyscall($1, $2, $4, int64($7), false, false);
                                                              Stracelex.(*lexer).result = $$;}
    | INT IDENTIFIER LPAREN types RPAREN EQUALS INT LPAREN parenthetical RPAREN {
                                                                  $$ = types.NewSyscall($1, $2, $4, $7, false, false);
                                                                  Stracelex.(*lexer).result = $$;}
    | INT IDENTIFIER LPAREN types RPAREN EQUALS UINT LPAREN parenthetical RPAREN {
                                                                  $$ = types.NewSyscall($1, $2, $4, int64($7), false, false);
                                                                  Stracelex.(*lexer).result = $$;}


parenthetical:
    identifiers {$$ = types.NewParenthetical();}
    | struct_type {$$ = types.NewParenthetical();}
    | array_type {$$ = types.NewParenthetical();}
    | expr_type {$$ = types.NewParenthetical();}
    | identifiers expr_type {$$ = types.NewParenthetical();}

types:
    type {types := make([]types.Type, 0); types = append(types, $1); $$ = types;}
    | types COMMA type {$1 = append($1, $3); $$ = $1;}



type:
    buf_type {$$ = $1}
    | field_type {$$ = $1}
    | expr_type {$$ = $1}
    | pointer_type {$$ = $1}
    | array_type {$$ = $1}
    | struct_type {$$ = $1}
    | dynamic_type {$$ = $1}
    | call_type {$$ = $1}
    | ipv4_type {$$ = $1}

dynamic_type:
    expr_type ARROW type {fmt.Printf("DYNAMIC TYPE\n"); $$ = types.NewDynamicType($1, $3)}

call_type:
    IDENTIFIER LPAREN types RPAREN {fmt.Printf("Call Type\n"); $$ = types.NewCallType($1, $3)}

macro_type:
    FLAG LPAREN types RPAREN {fmt.Printf("Macro Type\n"); $$ = types.NewMacroType($1, $3)}
    | FLAG LPAREN identifiers RPAREN {fmt.Printf("Macro Type\n"); $$ = types.NewMacroType($1, nil)}

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
    | LBRACKET_SQUARE FLAG FLAG RBRACKET_SQUARE {
                expr1 := types.NewExpression(types.NewFlagType($2));
                expr2 := types.NewExpression(types.NewFlagType($3));
                fmt.Printf("HAVE EXPR1 EXPR2\n");
                bs := types.NewBinarySet(expr1, expr2)
                $$=types.NewExpression(bs);}
    | int_type {$$ = types.NewExpression($1);}
    | macro_type {$$ = types.NewExpression($1);}
    | LPAREN expr_type RPAREN {$$ = $2;}
    | macro_type OR expr_type {$$ = types.NewExpression(types.NewBinop(types.NewExpression($1), types.OR, $3));}
    | flag_type OR expr_type {$$ = types.NewExpression(types.NewBinop(types.NewExpression($1), types.OR, $3));}
    | int_type OR expr_type {$$ = types.NewExpression(types.NewBinop(types.NewExpression($1), types.OR, $3));}
    | macro_type LAND expr_type {$$ = types.NewExpression(types.NewBinop(types.NewExpression($1), types.OR, $3));}
    | flag_type LAND expr_type {$$ = types.NewExpression(types.NewBinop(types.NewExpression($1), types.LAND, $3));}
    | int_type LAND expr_type {$$ = types.NewExpression(types.NewBinop(types.NewExpression($1), types.LAND, $3));}
    | macro_type LOR expr_type {$$ = types.NewExpression(types.NewBinop(types.NewExpression($1), types.LOR, $3));}
    | flag_type LOR expr_type {$$ = types.NewExpression(types.NewBinop(types.NewExpression($1), types.LOR, $3));}
    | int_type LOR expr_type {$$ = types.NewExpression(types.NewBinop(types.NewExpression($1), types.LOR, $3));}
    | macro_type LEQUAL expr_type {$$ = types.NewExpression(types.NewBinop(types.NewExpression($1), types.LEQUAL, $3));}
    | flag_type LEQUAL expr_type {$$ = types.NewExpression(types.NewBinop(types.NewExpression($1), types.LEQUAL, $3));}
    | int_type LEQUAL expr_type {$$ = types.NewExpression(types.NewBinop(types.NewExpression($1), types.LEQUAL, $3));}
    | macro_type AND expr_type {$$ = types.NewExpression(types.NewBinop(types.NewExpression($1), types.AND, $3));}
    | flag_type AND expr_type {$$ = types.NewExpression(types.NewBinop(types.NewExpression($1), types.AND, $3));}
    | int_type AND expr_type {$$ = types.NewExpression(types.NewBinop(types.NewExpression($1), types.AND, $3));}
    | macro_type LSHIFT expr_type {$$ = types.NewExpression(types.NewBinop(types.NewExpression($1), types.LSHIFT, $3));}
    | flag_type LSHIFT expr_type {$$ = types.NewExpression(types.NewBinop(types.NewExpression($1), types.LSHIFT, $3));}
    | int_type LSHIFT expr_type {$$ = types.NewExpression(types.NewBinop(types.NewExpression($1), types.LSHIFT, $3));}
    | macro_type RSHIFT expr_type {$$ = types.NewExpression(types.NewBinop(types.NewExpression($1), types.RSHIFT, $3));}
    | flag_type RSHIFT expr_type {$$ = types.NewExpression(types.NewBinop(types.NewExpression($1), types.RSHIFT, $3));}
    | int_type RSHIFT expr_type {$$ = types.NewExpression(types.NewBinop(types.NewExpression($1), types.RSHIFT, $3));}
    | int_type TIMES expr_type {$$ = types.NewExpression(types.NewBinop(types.NewExpression($1), types.TIMES, $3));}
    | ONESCOMP expr_type {$$ = types.NewExpression(types.NewUnop($2, types.ONESCOMP));}



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

