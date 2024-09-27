package util

import (
	"os"
)

func EnsureDirExists(path string) error {
	return os.MkdirAll(path, 0644)
}
