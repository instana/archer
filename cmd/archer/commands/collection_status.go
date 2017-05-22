package commands

type ArcherCollectionStatusCommand struct {
	*Meta
}

func (a *ArcherCollectionStatusCommand) Run(args []string) int {
	args, cmdFlags := a.setup(args, "collection status")
	cmdFlags.Usage = func() { a.Ui.Output(a.Help()) }

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	return 0
}

func (a *ArcherCollectionStatusCommand) Help() string {
	return `Start collection COLLECTION`
}

func (a *ArcherCollectionStatusCommand) Synopsis() string {
	return "Start collection COLLECTION"
}
