package lsp

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

var details string = "This is a test item"
var insert string = "Please insert this cool text from Rickshaw^tm !"

func text_document_completion(context *glsp.Context, params *protocol.CompletionParams) (interface{}, error) {
	var result []protocol.CompletionItem

	result = append(result, protocol.CompletionItem{
		Label:      "Rickshaw",
		Detail:     &details,
		InsertText: &insert,
	})

	return result, nil
}
