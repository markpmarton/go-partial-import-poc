package main

import (
	"gitlab.com/mark-p-marton/go-partial-import-poc/internal/server"
)

func main() {
	// processing.ProcessProto()
	ss := server.SampleServer{}
	ss.StartServer()
}
