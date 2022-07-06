package semantic

// Token types according to Treesitter grammar (grammar.go)
type TokenType uint32

const (
	INCLUDE TokenType = iota
	VARIABLE
	OPERATOR
	FUNCTION
	KEYWORD
	REPEAT
	CONDITIONAL
	COMMENT
	STRING
	NUMBER
	BOOLEAN
	CONSTANT
	PUNCTUATION
	NONE
	FIELD
	DEFINITION
	LABEL
	ATTRIBUTE
)

var (
	tokenMap  map[uint16]TokenType
	TokenKeys []string
)

// Treesitter grammar to semantic token translation
// shall be equal to Token enum above (especially the order!)
func init() {
	// send back keys that the language server speaks to client
	TokenKeys = []string{
		"include",
		"variable",
		"operator",
		"function",
		"keyword",
		"repeat",
		"conditional",
		"comment",
		"string",
		"number",
		"boolean",
		"constant",
		"punctuation",
		"none",
		"field",
		"definition",
		"label",
		"attribute",
	}

	// binds grammar index to semantic token type
	tokenMap = map[uint16]TokenType{
		0:  INCLUDE,
		1:  VARIABLE,
		2:  VARIABLE,
		3:  OPERATOR,
		4:  OPERATOR,
		5:  OPERATOR,
		6:  OPERATOR,
		7:  OPERATOR,
		8:  OPERATOR,
		9:  FUNCTION,
		10: FUNCTION,
		11: KEYWORD,
		12: REPEAT,
		13: REPEAT,
		14: CONDITIONAL,
		15: COMMENT,
		16: STRING,
		17: NUMBER,
		18: BOOLEAN,
		19: CONSTANT,
		20: PUNCTUATION,
		21: PUNCTUATION,
		22: PUNCTUATION,
		23: NONE,
		24: FIELD,
		25: DEFINITION,
		26: LABEL,
		27: LABEL,
		28: ATTRIBUTE,
	}
}

// Retrieve token index in O(1), memory complexity O(n)
// index is the position in the query grammar
func TokenTypeIndex(index uint16) uint32 {
	val, ok := tokenMap[index]
	if !ok {
		return 0
	}
	return uint32(val)
}
