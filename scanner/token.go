package scanner

import (
	"cuelang.org/go/cue/token"
	tk "cuelang.org/go/cue/token"
)

// Scanner package type to semantic package type
type Token int

const (
	ILLEGAL Token = iota
	EOF
	COMMENT
	ATTRIBUTE
	IDENT
	INT
	FLOAT
	STRING
	INTERPOLATION
	OPERATOR
	DELIMITER
	KEYWORD
)

var (
	tokenMap  map[tk.Token]Token
	TokenKeys []string
)

// Scanner token to semantic token translation
func init() {

	// shall be equal to Token above (especially the order)
	TokenKeys = []string{
		"ILLEGAL",
		"EOF",
		"COMMENT",
		"ATTRIBUTE",
		"IDENT",
		"INT",
		"FLOAT",
		"STRING",
		"INTERPOLATION",
		"OPERATOR",
		"DELIMITER",
		"KEYWORD",
	}

	tokenMap = map[tk.Token]Token{
		tk.ILLEGAL:       ILLEGAL,
		tk.EOF:           EOF,
		tk.COMMENT:       COMMENT,
		tk.ATTRIBUTE:     ATTRIBUTE,
		tk.IDENT:         IDENT,
		tk.INT:           INT,
		tk.FLOAT:         FLOAT,
		tk.STRING:        STRING,
		tk.INTERPOLATION: INTERPOLATION,
		tk.BOTTOM:        OPERATOR,
		tk.ADD:           OPERATOR,
		tk.SUB:           OPERATOR,
		tk.MUL:           OPERATOR,
		tk.POW:           OPERATOR,
		tk.QUO:           OPERATOR,
		tk.IQUO:          OPERATOR,
		tk.IREM:          OPERATOR,
		tk.IDIV:          OPERATOR,
		tk.IMOD:          OPERATOR,
		tk.AND:           OPERATOR,
		tk.OR:            OPERATOR,
		tk.LAND:          OPERATOR,
		tk.LOR:           OPERATOR,
		tk.BIND:          OPERATOR,
		tk.EQL:           OPERATOR,
		tk.LSS:           OPERATOR,
		tk.GTR:           OPERATOR,
		tk.NOT:           OPERATOR,
		tk.ARROW:         OPERATOR,
		tk.NEQ:           OPERATOR,
		tk.LEQ:           OPERATOR,
		tk.GEQ:           OPERATOR,
		tk.MAT:           OPERATOR,
		tk.NMAT:          OPERATOR,
		tk.LPAREN:        DELIMITER,
		tk.LBRACK:        DELIMITER,
		tk.LBRACE:        DELIMITER,
		tk.COMMA:         DELIMITER,
		tk.PERIOD:        DELIMITER,
		tk.ELLIPSIS:      DELIMITER,
		tk.RPAREN:        DELIMITER,
		tk.RBRACK:        DELIMITER,
		tk.RBRACE:        DELIMITER,
		tk.SEMICOLON:     DELIMITER,
		tk.COLON:         DELIMITER,
		tk.ISA:           DELIMITER,
		tk.OPTION:        DELIMITER,
		tk.IF:            KEYWORD,
		tk.FOR:           KEYWORD,
		tk.IN:            KEYWORD,
		tk.LET:           KEYWORD,
		tk.TRUE:          KEYWORD,
		tk.FALSE:         KEYWORD,
		tk.NULL:          KEYWORD,
	}
}

// Retrieve token index in O(1), memory complexity O(n)
// token will always be lowercase
func TokenTypeIndex(tok token.Token) (Token, bool) {
	i, ok := tokenMap[tok]
	return i, ok
}
