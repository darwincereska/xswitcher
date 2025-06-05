package main

import (
	"fmt"
	"os"
	"xswitcher/core"
	"xswitcher/core/config"
	"xswitcher/core/utils"
)

func main() {
	cfg, _ := config.ParseConfig("")
	if len(os.Args) < 2 {
		core.PrintConstants()
		fmt.Println("Usage: xswitcher <command> [args...]")
		os.Exit(1)
	}

	mainLayout := cfg.Layouts.Main
	secondaryLayout := cfg.Layouts.Secondary
	command := os.Args[1:]

	exitCode := utils.RunProgram(command, mainLayout, secondaryLayout)
	os.Exit(exitCode)
}
