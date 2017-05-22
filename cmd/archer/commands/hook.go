package commands

import (
	"strings"
)

type ArcherHookCommand struct {
	*Meta
}

func (a *ArcherHookCommand) Run(args []string) int {
	var err error

	args, cmdFlags := a.setup(args, "hook")
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

func (a *ArcherHookCommand) Help() string {
	help := `
Usage: archer hook HOOK

  Realizes resources for assigned HOOK.

`
	return strings.TrimSpace(help)
}

func (a *ArcherHookCommand) Synopsis() string {
	return "Realizes resources for assigned HOOK"
}
