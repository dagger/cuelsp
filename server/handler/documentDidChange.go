package handler

import (
	"fmt"
	"os"

	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"go.lsp.dev/uri"
)

func (h *Handler) documentDidChange(_ *glsp.Context, params *protocol.DidChangeTextDocumentParams) error {
	h.log.Debugf("Document saved")
	h.log.Debugf("params: %#v", params)

	_uri, err := uri.Parse(params.TextDocument.URI)
	if err != nil {
		return h.wrapError(err)
	}

	p := h.workspace.GetPlan(_uri.Filename())
	if p == nil {
		return h.wrapError(fmt.Errorf("plan not found"))
	}

	ogf, err := os.ReadFile(_uri.Filename())
	if err != nil {
		return h.wrapError(err)
	}
	content := string(ogf)

	for _, change := range params.ContentChanges {
		if change_, ok := change.(protocol.TextDocumentContentChangeEvent); ok {
			startIndex, endIndex := change_.Range.IndexesIn(content)
			content = content[:startIndex] + change_.Text + content[endIndex:]
		} else if change_, ok := change.(protocol.TextDocumentContentChangeEventWhole); ok {
			content = change_.Text
		}
	}
	p.AddOverride(h.workspace.TrimRootPath(_uri.Filename()), []byte(content))

	return p.Reload()
}
