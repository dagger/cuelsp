package semantictoken

// type TokenModifier int

// const (
// 	Declaration TokenModifier = 1 << iota // Declaration   = 0
// 	Definition
// 	Readonly
// 	Static
// 	Deprecated
// 	Abstract
// 	Async
// 	Modification
// 	Documentation
// 	DefaultLibrary
// )

// var (
// 	tokenModifiersMap map[string]TokenModifier
// 	tokenTypeKeys     []string
// )

// func init() {
// 	tokenModifiersMap = map[string]TokenModifier{
// 		"declaration":    Declaration,
// 		"definition":     Definition,
// 		"readonly":       Readonly,
// 		"static":         Static,
// 		"deprecated":     Deprecated,
// 		"abstract":       Abstract,
// 		"async":          Async,
// 		"modification":   Modification,
// 		"documentation":  Documentation,
// 		"defaultLibrary": DefaultLibrary,
// 	}

// }

// // // Retrieve token index in O(1), memory complexity O(n)
// // func tokenModifierIndex(token string) (TokenModifier, bool) {
// // 	i, ok := tokenModifiersMap[strings.ToLower(token)]
// // 	return i, ok
// // }

// // // verify that clientCapabilities.TextDocument.SemanticTokens.TokenModifiers
// // // is the same as our expected one
// // func VerifyClientTokenModifiers(tokens []string) bool {
// // 	return len(tokens) == len(tokenModifiersMap)
// // }
