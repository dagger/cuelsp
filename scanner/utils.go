package scanner

import (
	"cuelang.org/go/cue/scanner"
	"cuelang.org/go/cue/token"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

// A Mode value is a set of flags (or 0).
// They control scanner behavior.
type Mode scanner.Mode

// These constants are options to the Init function.
const (
	ScanComments     Mode = 1 << iota // return comments as COMMENT tokens
	dontInsertCommas                  // do not automatically insert commas - for testing only
)

type ErrorCollector struct {
	msg string    // last error message encountered
	pos token.Pos // last error position encountered
}

type Scan struct {
	s          scanner.Scanner
	h          []ErrorCollector
	startToken token.Token
	startPos   token.Pos
	startLit   string
	tokens     []protocol.UInteger
}
