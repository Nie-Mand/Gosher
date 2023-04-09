package main

import (
	"Nie-Mand/Gosher/server/core"
	"Nie-Mand/Gosher/server/utils"
)

func main() {
	undo := utils.InitLogger()
	defer undo()

	core.StartServer()
}
