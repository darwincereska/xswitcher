package main

import (
	"flag"
	"fmt"
	"os"
	"xswitcher/core"
	"xswitcher/core/config"
	"xswitcher/core/utils"
)

type Flags struct {
	HELP bool
	INIT bool
}

func parseFlags() *Flags {
	flags := &Flags{}
	flag.BoolVar(&flags.HELP, "help", false, "Shows help menu")
	flag.BoolVar(&flags.INIT, "init", false, "Generates a Config File")
	flag.Parse()

	switch {
	case flags.HELP:
		flag.Usage()
		os.Exit(1)
	case flags.INIT:
		config.Init("")
		os.Exit(1)
	}

	return flags
}

func main() {
	// Parse Flags First
	parseFlags()

	// Load Config File
	cfg, _ := config.ParseConfig("")
	if len(os.Args) < 2 {
		core.PrintConstants()
		fmt.Println("Usage: xswitcher <command> [args...]")
		os.Exit(1)
	}

	// Set Layouts
	mainLayout := cfg.Layouts.Main
	secondaryLayout := cfg.Layouts.Secondary

	// Rest of args to run
	command := os.Args[1:]
	exitCode := utils.RunProgram(command, mainLayout, secondaryLayout)
	os.Exit(exitCode)
}
