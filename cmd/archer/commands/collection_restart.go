package commands

type ArcherCollectionRestartCommand struct {
	*Meta
}

func (a *ArcherCollectionRestartCommand) Run(args []string) int {
	args, cmdFlags := a.setup(args, "collection restart")
	cmdFlags.Usage = func() { a.Ui.Output(a.Help()) }

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	return 0
}

func (a *ArcherCollectionRestartCommand) Help() string {
	return `List configured collections`
}

func (a *ArcherCollectionRestartCommand) Synopsis() string {
	return "List configured collections"
}
