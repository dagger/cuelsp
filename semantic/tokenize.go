package semantic

import (
	"context"

	"github.com/dagger/dlsp/parser"
	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/cue"
	"github.com/tliron/kutil/logging"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func initScanner(URI string, log logging.Logger) (*sitter.Node, *Scanner, error) {
	sitterParser := sitter.NewParser()
	sitterParser.SetLanguage(cue.GetLanguage())

	code := parser.ReadFile(URI)
	tree, err := sitterParser.ParseCtx(context.Background(), nil, []byte(code))
	if err != nil {
		return nil, nil, err
	}

	return tree.RootNode(), &Scanner{
		startLine: 0,
		startCol:  0,
		logger:    log,
	}, nil
}

func Tokenize(URI string, log logging.Logger) ([]protocol.UInteger, error) {
	rootNode, scanner, err := initScanner(URI, log)
	if err != nil {
		return nil, err
	}

	// capture all tokens from query grammar
	q, _ := sitter.NewQuery([]byte(grammar), cue.GetLanguage())
	qc := sitter.NewQueryCursor()
	qc.Exec(q, rootNode)

	// transform tokens to protocol.UInteger
	for {
		m, ok := qc.NextMatch()
		if !ok {
			break
		}
		for _, c := range m.Captures {
			// log.Debugf("capture: %v", c)
			scanner.Scan(m.PatternIndex, c)
		}
	}
	log.Debugf("finished")
	return scanner.tokens, nil
}
