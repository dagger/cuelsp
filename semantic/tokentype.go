package semantic

import (
	"strings"

	sitter "github.com/smacker/go-tree-sitter"
)

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
	VALUE
	LABEL
	ATTRIBUTE
	DEFINITION
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
		"value",
		"label",
		"attribute",
		"definition",
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
		25: VALUE, // not in grammar, computed manually
		26: LABEL,
		27: LABEL,
		28: ATTRIBUTE,
		29: DEFINITION,
	}
}

// Checks whethger a node is a definition
func isDefinition(node *sitter.Node, code []byte) (uint32, bool) {
	codeStr := node.Content(code)
	if (strings.HasPrefix(codeStr, "#") || strings.HasPrefix(codeStr, "_#")) && !strings.Contains(codeStr, ".") {
		return uint32(DEFINITION), true
	}
	return 0, false
}

// Retrieve token index in O(1), memory complexity O(n)
// index is the position in the query grammar
func tokenTypeIndex(index uint16, node *sitter.Node, code []byte) uint32 {
	// if it is a field or a value (position 24/25 in grammar), manually check if it's a definition
	if index == 24 || index == 25 {
		if val, ok := isDefinition(node, code); ok {
			return val
		}
	}

	val, ok := tokenMap[index]
	if !ok {
		return 0
	}
	return uint32(val)
}
