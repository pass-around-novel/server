package main

import (
	"./cmd"
	"./logger/console"
	"./server"
)

func main() {
	console.Init()
	server.Init()
	cmd.Execute()
}
