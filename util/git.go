package util

import (
	"os"
	"os/exec"
	"path/filepath"
)

func EnsureGitRepo(path string) error {
	err := EnsureDirExists(path)
	if err != nil {
		return err
	}
	if IsGitRepo(path) {
		return nil
	}
	return InitGitRepo(path)
}

func IsGitRepo(path string) bool {
	git_directory := filepath.Join(path, "/.git")
	fileInfo, err := os.Stat(git_directory)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

func InitGitRepo(path string) error {
	cmd := exec.Command("git", "init")
	cmd.Dir = path
	return cmd.Run()
}
