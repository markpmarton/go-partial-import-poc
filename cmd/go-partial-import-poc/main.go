package main

import (
	"github.com/markpmarton/go-partial-import-poc/internal/server"
)

func main() {
	// processing.ProcessProto()
	ss := server.SampleServer{}
	ss.StartServer()
}
