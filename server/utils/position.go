package utils

import (
	"github.com/dagger/cuelsp/loader"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"go.lsp.dev/uri"
)

func CueLocationToLSPLocation(v *loader.Value) protocol.Location {
	pos := protocol.Position{
		Line:      IntToUInt(v.Pos().Line()),
		Character: IntToUInt(v.Pos().Column()),
	}

	return protocol.Location{
		URI: string(uri.File(v.Pos().Filename())),
		Range: protocol.Range{
			Start: pos,
			End:   pos,
		},
	}
}
