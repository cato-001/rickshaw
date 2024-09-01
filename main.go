package main

import (
	"rickshaw/lsp"

	"github.com/tliron/commonlog"
	protocol "github.com/tliron/glsp/protocol_3_16"
	"github.com/tliron/glsp/server"
)

func main() {
	server := lsp.NewRickshawServer()
	server.RunStdio()
}
