package commands

type ArcherCollectionStopCommand struct {
	*Meta
}

func (a *ArcherCollectionStopCommand) Run(args []string) int {
	args, cmdFlags := a.setup(args, "collection stop")
	cmdFlags.Usage = func() { a.Ui.Output(a.Help()) }

	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	return 0
}

func (a *ArcherCollectionStopCommand) Help() string {
	return `Stop collection COLLECTION`
}

func (a *ArcherCollectionStopCommand) Synopsis() string {
	return "Stop collection COLLECTION"
}
