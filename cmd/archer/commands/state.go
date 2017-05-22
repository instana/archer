package commands

import (
	"strings"
)

type ArcherStateCommand struct {
	*Meta
}

func (a *ArcherStateCommand) Run(args []string) int {
	var err error

	args, cmdFlags := a.setup(args, "state")
	cmdFlags.Usage = func() { a.Ui.Output(a.Help()) }

	if err = cmdFlags.Parse(args); err != nil {
		return 1
	}

	if cmdFlags.NArg() == 0 {
		cmdFlags.Usage()
		return 1
	}

	return 0
}

func (a *ArcherStateCommand) Help() string {
	help := `
Usage: archer state OPERATION

  Manages archer state database.

`
	return strings.TrimSpace(help)
}

func (a *ArcherStateCommand) Synopsis() string {
	return "Manages archer state database"
}
