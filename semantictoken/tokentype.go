package semantictoken

type TokenType int

const (
	Namespace TokenType = iota // Namespace   = 0
	Type
	Class
	Enum
	Interface
	Struct
	TypeParameter
	Parameter
	Variable
	Property
	EnumMember
	Event
	Function
	Method
	Macro
	Keyword
	Modifier
	Comment
	String
	Number
	Regexp
	Operator
	Decorator
)

var (
	tokenTypeMap  map[string]TokenType
	TokenTypeKeys []string
)

func init() {
	TokenTypeKeys = []string{
		"namespace",
		"type",
		"class",
		"enum",
		"interface",
		"struct",
		"typeParameter",
		"parameter",
		"variable",
		"property",
		"enumMember",
		"event",
		"function",
		"method",
		"macro",
		"keyword",
		"modifier",
		"comment",
		"string",
		"number",
		"regexp",
		"operator",
		"decorator",
	}

	tokenTypeMap = map[string]TokenType{
		"namespace":     Namespace,
		"type":          Type,
		"class":         Class,
		"enum":          Enum,
		"interface":     Interface,
		"struct":        Struct,
		"typeParameter": TypeParameter,
		"parameter":     Parameter,
		"variable":      Variable,
		"property":      Property,
		"enumMember":    EnumMember,
		"event":         Event,
		"function":      Function,
		"method":        Method,
		"macro":         Macro,
		"keyword":       Keyword,
		"modifier":      Modifier,
		"comment":       Comment,
		"string":        String,
		"number":        Number,
		"regexp":        Regexp,
		"operator":      Operator,
		"decorator":     Decorator,
	}
}

// Retrieve token index in O(1), memory complexity O(n)
// token will always be lowercase
func tokenTypeIndex(token string) (TokenType, bool) {
	i, ok := tokenTypeMap[token]
	return i, ok
}

// verify that clientCapabilities.TextDocument.SemanticTokens.TokenTypes
// is the same as our expected one
// func VerifyClientTokenTypes(tokens []string) bool {
// 	return len(tokens) == len(tokenTypesMap)
// }
