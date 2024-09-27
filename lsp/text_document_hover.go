package lsp

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func text_document_hover(context *glsp.Context, params *protocol.HoverParams) (*protocol.Hover, error) {
	result := protocol.Hover{
		Contents: protocol.MarkupContent{},
	}
	return &result, nil
}
