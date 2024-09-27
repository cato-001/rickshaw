package lsp

import (
	"errors"
	"os"
	"path/filepath"
	"rickshaw/csharp"

	"github.com/tliron/commonlog"
	"github.com/tliron/glsp"
	protocol "github.com/tliron/glsp/protocol_3_16"
)

var solutions []csharp.Solution

func initialize(context *glsp.Context, params *protocol.InitializeParams) (any, error) {
	commonlog.NewInfoMessage(0, "Initializing server...")

	solutions = solutions_in_workspaces(params.WorkspaceFolders)
	file, err := os.OpenFile(".rick.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, errors.New("")
	}

	capabilities := handler.CreateServerCapabilities()
	capabilities.CompletionProvider = &protocol.CompletionOptions{}
	capabilities.HoverProvider = &protocol.HoverOptions{}

	return protocol.InitializeResult{
		Capabilities: capabilities,
		ServerInfo: &protocol.InitializeResultServerInfo{
			Name:    SERVER_NAME,
			Version: &version,
		},
	}, nil
}

func solutions_in_workspaces(workspaces []protocol.WorkspaceFolder) []csharp.Solution {
	var solutions []csharp.Solution
	for _, workspace := range workspaces {
		path := filepath.Join(workspace.URI, "**/*.sln")
		matches, err := filepath.Glob(path)
		if err != nil {
			continue
		}
		for _, match := range matches {
			solutions = append(solutions, csharp.Solution{
				Path: match,
			})
		}
	}
	return solutions
}
