package commands

type ArcherCollectionPurgeCommand struct {
	*Meta
}

func (a *ArcherCollectionPurgeCommand) Run(args []string) int {
	args, cmdFlags := a.setup(args, "collection purge")
	cmdFlags.Usage = func() { a.Ui.Output(a.Help()) }

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	return 0
}

func (a *ArcherCollectionPurgeCommand) Help() string {
	return `Purge collection COLLECTION`
}

func (a *ArcherCollectionPurgeCommand) Synopsis() string {
	return "Purge collection COLLECTION"
}
