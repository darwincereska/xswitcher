package core

import (
	"fmt"
)

const (
	NAME       = "XSwitcher"
	VERSION    = "v0.1.2"
	AUTHOR     = "Darwin Cereska"
	GITHUB     = "https://github.com/darwincereska"
	CONFIG_DIR = ".config"
	FILENAME   = "xswitcher.yml"
)

func PrintConstants() {
	fmt.Printf("%s %s\n", NAME, VERSION)
	fmt.Println("_______________")
	fmt.Printf("Made By: %s\n", AUTHOR)
	fmt.Println(GITHUB)
	fmt.Println("_______________")
}
