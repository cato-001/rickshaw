package state

import (
	"path/filepath"
	"rickshaw/util"
)

type Store struct {
	path      string
	workspace Workspace
}

func InitStore(workspace Workspace) (Store, error) {
	path := storepath(workspace.path)
	store := Store{path, workspace}
	err := util.EnsureGitRepo(path)
	if err != nil {
		return store, err
	}
	return store, nil
}

func storepath(wsPath string) string {
	_, file := filepath.Split(wsPath)
	storedir := "~/.local/state/rickshaw/"
	result := storedir + file
	return result
}
