package main

import (
	"./cmd"
	"./initdb"
	"./logger/console"
	"./server"
)

func main() {
	console.Init()
	initdb.Init()
	server.Init()
	cmd.Execute()
}
