package commands

type ArcherCollectionStartCommand struct {
	*Meta
}

func (a *ArcherCollectionStartCommand) Run(args []string) int {
	args, cmdFlags := a.setup(args, "collection start")
	cmdFlags.Usage = func() { a.Ui.Output(a.Help()) }

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	return 0
}

func (a *ArcherCollectionStartCommand) Help() string {
	return `Start collection COLLECTION`
}

func (a *ArcherCollectionStartCommand) Synopsis() string {
	return "Start collection COLLECTION"
}
