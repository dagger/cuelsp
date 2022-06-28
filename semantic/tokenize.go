package semantic

import (
	"context"

	sitter "github.com/smacker/go-tree-sitter"
	"github.com/smacker/go-tree-sitter/cue"
	"github.com/tliron/kutil/logging"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

func initScanner(URI string, log logging.Logger) (*sitter.Node, *Scanner, error) {
	// Initialize Cue parser
	sitterParser := sitter.NewParser()
	sitterParser.SetLanguage(cue.GetLanguage())

	// Read file
	code, err := readFile(URI, log)
	if err != nil {
		return nil, nil, err
	}
	tree, err := sitterParser.ParseCtx(context.Background(), nil, code)
	if err != nil {
		return nil, nil, err
	}

	return tree.RootNode(), &Scanner{
		startLine: 0,
		startCol:  0,
		code:      code,
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

	// iterate over queryied tokens
	for {
		m, ok := qc.NextMatch()
		if !ok {
			break
		}
		for _, c := range m.Captures {
			scanner.Scan(m.PatternIndex, c)
		}
	}
	return scanner.tokens, nil
}
