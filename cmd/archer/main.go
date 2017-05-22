package main

import (
	"fmt"
	"os"

	"github.com/instana/archer/cmd/archer/commands"
	"github.com/mitchellh/cli"
)

func main() {
	ui := &cli.BasicUi{
		Reader:      os.Stdin,
		Writer:      os.Stdout,
		ErrorWriter: os.Stderr,
	}

	c := cli.NewCLI("archer", "0.0.1")
	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"build": func() (cli.Command, error) {
			return &commands.ArcherBuildCommand{
				&commands.Meta{Ui: ui},
			}, nil
		},
		"collection": func() (cli.Command, error) {
			return &commands.ArcherCollectionCommand{
				&commands.Meta{Ui: ui},
			}, nil
		},
		"hook": func() (cli.Command, error) {
			return &commands.ArcherHookCommand{
				&commands.Meta{Ui: ui},
			}, nil
		},
		"realize": func() (cli.Command, error) {
			return &commands.ArcherRealizeCommand{
				&commands.Meta{Ui: ui},
			}, nil
		},
		"state": func() (cli.Command, error) {
			return &commands.ArcherStateCommand{
				&commands.Meta{Ui: ui},
			}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}

	os.Exit(exitStatus)
}
