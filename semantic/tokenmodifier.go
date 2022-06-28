package semantic

// Token Modifier according to Treesitter grammar (grammar.go)
type TokenModifier uint32

const (
	DEFAULT TokenModifier = 1 << iota
	REGEXP
	UNIFY
	DISJUNCT
	BUILTIN
	SPECIAL
	DELIMITER
	BRACKET
)

var (
	tokenModifierMap  map[uint16]TokenModifier
	TokenModifierKeys []string
)

func init() {
	TokenModifierKeys = []string{
		"default",
		"regexp",
		"unify",
		"disjunct",
		"builtin",
		"special",
		"delimiter",
		"bracket",
	}

	// binds grammar index to semantic token modifier
	tokenModifierMap = map[uint16]TokenModifier{
		4:  DEFAULT,
		5:  REGEXP,
		6:  REGEXP,
		7:  UNIFY,
		8:  DISJUNCT,
		9:  BUILTIN,
		10: BUILTIN,
		19: BUILTIN,
		20: SPECIAL,
		21: DELIMITER,
		22: BRACKET,
		23: SPECIAL,
	}
}

// Retrieve token index in O(1), memory complexity O(n)
// index is the position in the query grammar
func tokenModifierIndex(index uint16) uint32 {
	val, ok := tokenModifierMap[index]
	if !ok {
		return 0
	}
	return uint32(val)
}
