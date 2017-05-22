package commands

import (
	"fmt"
	"strings"

	"github.com/instana/archer/archer"
)

type ArcherBuildCommand struct {
	*Meta
}

func (a *ArcherBuildCommand) Run(args []string) int {
	var err error

	args, cmdFlags := a.setup(args, "build")
	cmdFlags.Usage = func() { a.Ui.Output(a.Help()) }

	if err = cmdFlags.Parse(args); err != nil {
		return 1
	}

	builder := archer.NewBuilder()
	builder.AfPath(cmdFlags.Arg(0)).Debug(a.Debug)

	err = builder.Build()
	if err != nil {
		a.Ui.Error(fmt.Sprint("Error => ", err))
	}

	return 0
}

func (a *ArcherBuildCommand) Help() string {
	help := `
Usage: archer build PATH_TO_ARCHERFILE

  Builds archer packages.

`
	return strings.TrimSpace(help)
}

func (a *ArcherBuildCommand) Synopsis() string {
	return "Builds archer packages"
}
