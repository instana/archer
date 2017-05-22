package commands

type ArcherCollectionListCommand struct {
	*Meta
}

func (a *ArcherCollectionListCommand) Run(args []string) int {
	args, cmdFlags := a.setup(args, "collection list")
	cmdFlags.Usage = func() { a.Ui.Output(a.Help()) }

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	return 0
}

func (a *ArcherCollectionListCommand) Help() string {
	return `List configured collections`
}

func (a *ArcherCollectionListCommand) Synopsis() string {
	return "List configured collections"
}
