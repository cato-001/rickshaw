package csharp

import "path/filepath"

type Solution struct {
	path     string
	projects []Project
}

type Project struct {
	path    string
	targets []string
}

func InitSolution(path string) (Solution, error) {
}

func (s *Solution) Name() string {
	_, file := filepath.Split(s.path)
	return file
}

func (s *Solution) Root() string {
	root, _ := filepath.Split(s.path)
	return root
}
