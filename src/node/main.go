package main

import (
	"Nie-Mand/Gosher/cli/cmd"
	"Nie-Mand/Gosher/cli/utils"
)

func main() {
	undo := utils.InitLogger()
	defer undo()
	
	cmd.Execute()
	
}