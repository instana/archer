package commands

type ArcherCollectionShowCommand struct {
	*Meta
}

func (a *ArcherCollectionShowCommand) Run(args []string) int {
	args, cmdFlags := a.setup(args, "collection show")
	cmdFlags.Usage = func() { a.Ui.Output(a.Help()) }

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	return 0
}

func (a *ArcherCollectionShowCommand) Help() string {
	return `Show configured collection COLLECTION`
}

func (a *ArcherCollectionShowCommand) Synopsis() string {
	return "Show configured collection COLLECTION"
}
