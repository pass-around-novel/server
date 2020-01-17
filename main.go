package main

import (
	"./cmd"
	"./logger/console"
)

func main() {
	console.Init()
	cmd.Execute()
}
