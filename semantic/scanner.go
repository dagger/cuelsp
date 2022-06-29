package semantic

import (
	"math"

	sitter "github.com/smacker/go-tree-sitter"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/kutil/logging"
)

type Scanner struct {
	startLine uint32
	startCol  uint32
	tokens    []protocol.UInteger
	logger    logging.Logger
}

func (s *Scanner) append(diffLine uint32, diffCol uint32, node *sitter.Node, patternIndex uint16) {
	s.tokens = append(
		s.tokens,
		diffLine,
		diffCol,
		node.EndPoint().Column-node.StartPoint().Column,
		tokenTypeIndex(patternIndex),
		tokenModifierIndex(patternIndex),
	)
}

func (s *Scanner) Scan(patternIndex uint16, c sitter.QueryCapture) {
	deltaLine := int(c.Node.StartPoint().Row) - int(s.startLine)
	var deltaCol float64
	if deltaLine == 0 {
		deltaCol = float64(c.Node.StartPoint().Column) - float64(s.startCol)
	} else {
		deltaCol = float64(c.Node.StartPoint().Column)

	}
	s.logger.Debugf("------------------------1-----------------------\n")
	s.logger.Debugf("titi:[debLine:|%d|debCol:|%d|]patternIndex:|%d|len:%d|str:%s\n", c.Node.StartPoint().Row, c.Node.StartPoint().Column, patternIndex, c.Node.EndPoint().Column-c.Node.StartPoint().Column, c.Node.String())
	s.logger.Debugf("toto:[%v]|\n", s.tokens)
	s.logger.Debugf("-------------------------1----------------------\n")

	switch {
	case deltaLine >= 0 && deltaCol >= 0: // Normal case: append token at end of slice
		s.append(uint32(deltaLine), uint32(deltaCol), c.Node, patternIndex)
		s.startLine, s.startCol = c.Node.StartPoint().Row, c.Node.StartPoint().Column

	case deltaLine == 0 && deltaCol < 0: // Edge case: append current token in middle of slice
		tokenArr := newToken(patternIndex, c.Node)
		if newTokens, ok := shiftTokens(s.tokens, tokenArr, int(math.Abs(deltaCol)), s.logger); ok {
			s.tokens = newTokens
		} else {
			panic("Should never happen")
		}

	default: // Never reached
		s.logger.Errorf("could not shift tokens", "dLine", deltaLine, "dCol", deltaCol)
		panic("Should never happen")
	}
	s.logger.Debugf("------------------------2-----------------------\n")
	// s.logger.Debugf("titi:[debLine:|%d|debCol:|%d|]patternIndex:|%d|\n", c.Node.StartPoint().Row, c.Node.StartPoint().Column, patternIndex)
	s.logger.Debugf("toto:[%v]|\n", s.tokens)
	s.logger.Debugf("-------------------------2----------------------\n")
}

// s.logger.Debugf("---------YOOOOOOOO---------")
// s.logger.Debugf("toto:[startLine:|%d|startCol:|%d|]|[%d:%d]-len:|%d|patternIndex:|%d|\n", s.startLine, s.startCol, dLine, dCol, c.Node.EndPoint().Column-c.Node.StartPoint().Column, patternIndex)
