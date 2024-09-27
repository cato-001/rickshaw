package lsp

import (
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

func text_document_completion(context *glsp.Context, params *protocol.CompletionParams) (interface{}, error) {
	var result []protocol.CompletionItem
	for _, solution := range solutions {
		result = append(result, protocol.CompletionItem{
			Label:      solution.Name(),
			Detail:     &solution.Path,
			InsertText: &solution.Path,
		})
	}
	return result, nil
}
