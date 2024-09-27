package state

import (
	"fmt"
	"rickshaw/csharp"
	"rickshaw/util"
)

type Workspace struct {
	path      string
	solutions []csharp.Solution
}

func NewWorkspace(path string) (Workspace, error) {
	solutions := make([]csharp.Solution, 0)
	workspace := Workspace{path, solutions}
	if !util.IsGitRepo(path) {
		return workspace, fmt.Errorf("Workspaces is not a git repository: %s", path)
	}
	return workspace, nil
}
