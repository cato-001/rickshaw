package lsp

import (
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/glsp/server"
)

const SERVER_NAME = "Rickshaw - CSharp Language Server"

var version string = "0.0.1"
var handler protocol.Handler

func NewRickshawServer() *server.Server {
	handler := protocol.Handler{
		Initialize:             initialize,
		Shutdown:               shutdown,
		TextDocumentCompletion: text_document_completion,
	}
	return server.NewServer(&handler, SERVER_NAME, true)
}
