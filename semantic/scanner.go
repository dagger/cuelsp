package semantic

import (
	sitter "github.com/smacker/go-tree-sitter"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/kutil/logging"
)

type Scanner struct {
	code      []byte
	startLine uint32
	startCol  uint32
	tokens    []protocol.UInteger
	logger    logging.Logger
}

// Scan a QueryCapture and append the token to the scanner's tokens slice
// => Normal case: append the token at the end of the tokens slice
// => Edge case: append the token anywhere in the tokens slice ; update the token positions accordingly
func (s *Scanner) Scan(patternIndex uint16, c sitter.QueryCapture) {
	deltaLine := s.deltaLine(c)
	deltaCol := s.deltaColumn(c, deltaLine)

	switch {
	case deltaLine >= 0 && deltaCol >= 0: // Normal case: append token at end of slice
		s.append(deltaLine, deltaCol, c.Node, patternIndex)
		s.updateTokenPos(c)

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
}
