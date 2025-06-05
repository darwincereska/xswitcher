package config_test

import (
	"fmt"
	"testing"
	c "xswitcher/core/config"
)

func TestParse(t *testing.T) {
	cfg, _ := c.ParseConfig("")
	fmt.Println(cfg.Layouts.Main)
	fmt.Println(cfg.Layouts.Secondary)
}
