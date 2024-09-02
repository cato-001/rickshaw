package main

import (
	"fmt"
	"rickshaw/lsp"
)

func main() {
	server := lsp.NewRickshawServer()
	err := server.RunStdio()
	if err != nil {
		fmt.Println("An error occured while running the language server:", err)
	}
}
