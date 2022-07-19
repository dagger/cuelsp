package semantic

import (
	sitter "github.com/smacker/go-tree-sitter"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/kutil/logging"
)

type NodeType uint32

const (
	MULTILINE_STRING TokenType = iota
	OTHER
)

type Scanner struct {
	code      []byte
	startLine uint32
	startCol  uint32
	endLine   uint32
	endCol    uint32
	nodeType  NodeType
	tokens    []protocol.UInteger
	logger    logging.Logger
}

// Scan a QueryCapture and append the token to the scanner's tokens slice
// => Normal case: append the token at the end of the tokens slice
// => Edge case: append the token anywhere in the tokens slice ; update the token positions accordingly
// If the token is a multiline string, append their children as arbitrary string tokens, and skip them in next iterations
func (s *Scanner) Scan(patternIndex uint16, c sitter.QueryCapture) {
	deltaLine := s.deltaLine(c)
	deltaCol := s.deltaColumn(c, deltaLine)

	if skip := s.skipMultilineStringChildren(c); skip {
		return
	}

	switch {
	case deltaLine >= 0 && deltaCol >= 0: // Normal case: append token at end of slice
		s.append(deltaLine, deltaCol, c.Node, patternIndex)
		s.updateTokenPos(c)
		if isMultilineStringToken(c.Node.String()) {
			s.appendArbitraryMultilineToken(c)
		}

	case deltaLine == 0 && deltaCol < 0: // Edge case: append current token anywhere in token slice
		tokenArr := s.newToken(patternIndex, c.Node, s.code)
		if newTokens, ok := shiftTokens(s.tokens, tokenArr, deltaCol, s.logger); ok {
			s.tokens = newTokens
		} else { // Never reached
			s.logger.Errorf("could not shift tokens", "dLine", deltaLine, "dCol", deltaCol)
		}

	default: // Never reached
		s.logger.Errorf("could not shift tokens", "dLine", deltaLine, "dCol", deltaCol)
	}
}

// Retrieve token length (end - start)
// => manages `multiline string` token. Particularity: (end < start)
func (s *Scanner) tokenLen(node *sitter.Node) uint32 {
	if node.StartPoint().Column <= node.EndPoint().Column {
		return node.EndPoint().Column - node.StartPoint().Column
	}
	return uint32(len(node.Content(s.code)))
}

// Creates a new token aimed to be appended in the middle of the token's slice
// => DeltaCol will be updated prior being appended. The 0 value is temporary
func (s *Scanner) newToken(patternIndex uint16, node *sitter.Node, code []byte) []protocol.UInteger {
	return []uint32{
		0,
		0,
		s.tokenLen(node),
		tokenTypeIndex(patternIndex, node, code),
		tokenModifierIndex(patternIndex),
	}
}

// Append the token according to the LSP protocol
func (s *Scanner) append(diffLine int, diffCol float64, node *sitter.Node, patternIndex uint16) {
	s.tokens = append(
		s.tokens,
		uint32(diffLine),
		uint32(diffCol),
		s.tokenLen(node),
		tokenTypeIndex(patternIndex, node, s.code),
		tokenModifierIndex(patternIndex),
	)
}

// Delta of lines between current and previous token
func (s *Scanner) deltaLine(c sitter.QueryCapture) int {
	return int(c.Node.StartPoint().Row) - int(s.startLine)
}

// Delta of columns between current and previous token
// => Takes into account whether we changed line or not
func (s *Scanner) deltaColumn(c sitter.QueryCapture, deltaLine int) float64 {
	if deltaLine == 0 {
		return float64(c.Node.StartPoint().Column) - float64(s.startCol)
	}
	return float64(c.Node.StartPoint().Column)
}

// Update token position inside scanner struct
// => current token becomes the old token for next iteration
func (s *Scanner) updateTokenPos(c sitter.QueryCapture) {
	s.startLine, s.startCol = c.Node.StartPoint().Row, c.Node.StartPoint().Column
	s.endLine, s.endCol = c.Node.EndPoint().Row, c.Node.EndPoint().Column
	s.nodeType = NodeType(OTHER)
}

// Append arbitrary multiline token, representing the children of a multiline string token
func (s *Scanner) appendArbitraryMultilineToken(c sitter.QueryCapture) {
	for i := 0; i < s.deltaEndLine(c); i++ {
		s.tokens = append(
			s.tokens,
			uint32(1),      // token exists in a new line
			uint32(0),      // start at col 0
			200,            // token has arbitrary len of 200
			uint32(STRING), // Token of string type
			0,
		)
	}
	s.updateMultilineTokenPos(c)
}

// Delta of lines between end and start positions of current token
func (s *Scanner) deltaEndLine(c sitter.QueryCapture) int {
	return int(c.Node.EndPoint().Row) - int(c.Node.StartPoint().Row)
}

// Update token position to Ending position of multiline token
func (s *Scanner) updateMultilineTokenPos(c sitter.QueryCapture) {
	s.startLine, s.startCol = c.Node.EndPoint().Row, c.Node.EndPoint().Column
	s.endLine, s.endCol = c.Node.EndPoint().Row, c.Node.EndPoint().Column
	s.nodeType = NodeType(MULTILINE_STRING)
}

// Check if we need to current token as part of multiline tokens
func (s *Scanner) skipMultilineStringChildren(c sitter.QueryCapture) bool {
	if s.nodeType == NodeType(MULTILINE_STRING) {
		if s.isInsideMultilineByRow(c) || s.isInsideMultilineByCol(c) {
			return true
		} else if !s.isInsideMultilineByCol(c) {
			s.nodeType = NodeType(OTHER)
		}
	}
	return false
}

// Check if current token is inside a multiline string
func (s *Scanner) isInsideMultilineByRow(c sitter.QueryCapture) bool {
	return c.Node.EndPoint().Row < s.endLine
}

// Check if current token is inside a multiline string
func (s *Scanner) isInsideMultilineByCol(c sitter.QueryCapture) bool {
	return c.Node.EndPoint().Row == s.endLine && c.Node.EndPoint().Column <= s.endCol
}
