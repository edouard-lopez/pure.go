package main

import (
	"fmt"

	pure "github.com/edouard-lopez/pure.go/pkg"
	"github.com/leaanthony/clir"
)

type Flags struct {
	Version           bool `name:"version" description:"version of pure"`
	LastCommandStatus int  `name:"last-command-status" description:"status code of last command"`
}

func main() {
	// Create new cli
	cli := clir.NewCli("Pure", "Pure prompt in go", "v0.0.1")

	var version bool
	cli.BoolFlag("version", "Pure version", &version)

	var lastStatusCommand int
	cli.IntFlag("last-command-status", "Are you awesome?", &lastStatusCommand)

	// Define action for the command
	cli.Action(func() error {
		fmt.Printf("%s\n", pure.Get(lastStatusCommand))
		return nil
	})

	if err := cli.Run(); err != nil {
		fmt.Printf("Error encountered: %v\n", err)
	}
}
