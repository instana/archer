package commands

import (
	"strings"
)

type ArcherRealizeCommand struct {
	*Meta
}

func (a *ArcherRealizeCommand) Run(args []string) int {
	var err error

	args, cmdFlags := a.setup(args, "realize")
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

func (a *ArcherRealizeCommand) Help() string {
	help := `
Usage: archer realize PHASE

  Realizes passed PHASE.

`
	return strings.TrimSpace(help)
}

func (a *ArcherRealizeCommand) Synopsis() string {
	return "Realizes provided archer PHASE"
}
