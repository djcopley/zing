package main

import (
	"github.com/djcopley/zing/cmd"
	"github.com/djcopley/zing/config"
)

func main() {
	conf := &config.Config{}
	cmd.Execute(conf)
}
