package main

import (
	// "Nie-Mand/Gosher/cli/utils"
	"Nie-Mand/Gosher/cli/cmd"
	"Nie-Mand/Gosher/cli/utils"
	// "os"
)

func main() {
	undo := utils.InitLogger()
	defer undo()
	
	cmd.Execute()
	
}