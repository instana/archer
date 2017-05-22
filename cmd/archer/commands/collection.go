package commands

import (
	"github.com/mitchellh/cli"
)

type ArcherCollectionCommand struct {
	*Meta
}

func (a *ArcherCollectionCommand) Run(args []string) int {
	repoCli := cli.NewCLI("archer collection", "")
	repoCli.Args = args

	repoCli.Commands = map[string]cli.CommandFactory{
		"list": func() (cli.Command, error) {
			return &ArcherCollectionListCommand{&Meta{Ui: a.Ui}}, nil
		},
	}

	if exitStatus, err := repoCli.Run(); err != nil {
		a.Ui.Error(err.Error())
		return exitStatus
	} else {
		return exitStatus
	}
}

func (a *ArcherCollectionCommand) Help() string {
	return `Manage archer collections`
}

func (a *ArcherCollectionCommand) Synopsis() string {
	return "Manage archer collections"
}
