package scanner

import (
	"fmt"

	"cuelang.org/go/cue/scanner"
	"cuelang.org/go/cue/token"
	"github.com/dagger/dlsp/parser"
	"github.com/tliron/kutil/logging"

	protocol "github.com/tliron/glsp/protocol_3_16"
)

var cnt = 0

func FillData(scan *Scan, pos token.Pos, tok token.Token, lit string, log logging.Logger) {

	// len
	tokenLength := uint32(len(lit))
	if tokenLength == 0 {
		tokenLength = uint32(len(tok.String()))
	}

	var diffCol uint32
	var diffLine uint32

	// if startPos is the mocked one
	if !scan.startPos.IsValid() {
		// line diff
		diffLine = 0
		diffCol = 0
		log.Debugf("CNT-notvalid:|%d|[%d:%d]\n", cnt, diffLine, diffCol)
		cnt++
	} else {
		// line diff
		diffLine = (uint32(pos.Line()) - 1) - (uint32(scan.startPos.Line()) - 1)

		if diffLine > 0 {
			diffCol = uint32(pos.Column()) - 1
		} else {
			diffCol = (uint32(pos.Column()) - 1) - (uint32((*scan).startPos.Column()) - 1)
		}
		log.Debugf("CNT-valid:|%d|\n", cnt)
		cnt++
	}

	// tok.String()
	// log.Debugf("l:%d|c: %d|len(%d)|t:[%s]|deb:[[%s]]|tok:%d, %d, %d - %d\n", diffLine, diffCol, tokenLength, tok.String(), lit, tok.IsKeyword(), tok.IsLiteral(), tok.IsOperator(), tok)

	// find token type
	tokenType, ok := TokenTypeIndex(tok)
	if ok {
		(*scan).tokens = append((*scan).tokens, diffLine, diffCol, tokenLength, uint32(tokenType), 0) // IParsedToken{
	}
}

func initScanner(URI string) Scan {
	src := parser.ReadFile(URI)
	file := token.NewFile(URI, 1, len(src))

	// Initialize the scanner.
	scan := Scan{
		h:          []ErrorCollector{}, // collects all errors
		startToken: (token.Token)(51),
		startPos:   token.NoPos,
		startLit:   "",
	}

	// Error callback
	eh := func(pos token.Pos, msg string, args []interface{}) {
		scan.h = append(scan.h, ErrorCollector{
			msg: fmt.Sprintf(msg, args...),
			pos: pos,
		})
	}

	// Initialize scanner with proper options
	// s.Init(file, src, nil, scanner.ScanComments|scanner.dontInsertCommas)
	scan.s.Init(file, src, eh, scanner.Mode(ScanComments|dontInsertCommas))
	return scan
}

func ScanFile(URI string, log logging.Logger) []protocol.UInteger {
	scan := initScanner(URI)

	// Loop over all tokens in File
	for {
		p, tok, lit := scan.s.Scan()
		if tok == token.EOF {
			break
		}
		FillData(&scan, p, tok, lit, log)
		scan.startPos, scan.startToken, scan.startLit = p, tok, lit
	}

	return scan.tokens
}
