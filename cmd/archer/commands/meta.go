package commands

import (
	"flag"

	"github.com/mitchellh/cli"
	"github.com/mitchellh/colorstring"

	"github.com/instana/archer/cmd"
)

type Meta struct {
	Color bool
	Debug bool
	Ui    cli.Ui

	oldUi cli.Ui
}

func (m *Meta) setup(args []string, name string) ([]string, *flag.FlagSet) {
	m.Color = true

	for i, v := range args {
		if v == "-no-color" {
			m.Color = false
			args = append(args[:i], args[i+1:]...)
		}

		if v == "-debug" {
			m.Debug = true
			args = append(args[:i], args[i+1:]...)
		}
	}

	// Setup the ui
	m.oldUi = m.Ui
	m.Ui = &cmd.ColorizeUi{
		Colorize:   m.Colorize(),
		InfoColor:  "[green]",
		ErrorColor: "[red]",
		WarnColor:  "[yellow]",
		Ui:         m.oldUi,
	}

	cmdFlags := flag.NewFlagSet(name, flag.ContinueOnError)
	cmdFlags.Usage = func() {}

	return args, cmdFlags
}

func (m *Meta) Colorize() *colorstring.Colorize {
	return &colorstring.Colorize{
		Colors:  colorstring.DefaultColors,
		Disable: !m.Color,
		Reset:   true,
	}
}
