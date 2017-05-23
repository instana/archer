package action

import (
	"errors"
	"strings"

	"github.com/hashicorp/go-multierror"
)

type Pkg struct {
	Name        string `json:"name"`
	Arch        string `json:"arch"`
	Description string `json:"description"`
	Vendor      string `json:"vendor"`
	Maintainer  string `json:"maintainer"`
	Url         string `json:"url"`
	License     string `json:"license"`
	Version     string `json:"version"`
	Iteration   string `json:"iteration"`
	Branch      string `json:"branch"`
	VcsRevision string `json:"vcs_revision" hcl:"vcs_revision"`
}

func NewPkg() *Pkg {
	return &Pkg{}
}

func (p *Pkg) Key() string {
	key := []string{p.Type(), p.Type()}
	return strings.Join(key, ":")
}

func (p *Pkg) Columns() string {
	return strings.Join(
		[]string{
			p.Name,
			p.Arch,
		},
		"|",
	)
}

func (p *Pkg) Unique() string {
	key := []string{p.Type(), p.Type()}
	return strings.Join(key, ":")
}

func (p *Pkg) Type() string {
	return "pkg"
}

func (p *Pkg) Validate() error {
	var error *multierror.Error

	err := p.validateName()
	if err != nil {
		error = multierror.Append(error, err)
	}

	err = p.validateArch()
	if err != nil {
		error = multierror.Append(error, err)
	}

	err = p.validateDescription()
	if err != nil {
		error = multierror.Append(error, err)
	}

	return error.ErrorOrNil()
}

func (p *Pkg) Valid() bool {
	err := p.Validate()
	if err != nil {
		return false
	}

	return true
}

func (p *Pkg) validateName() error {
	if p.Name == "" {
		return errors.New("action pkg:name required")
	}

	return nil
}

func (p *Pkg) validateArch() error {
	if p.Arch != "x86_64" && p.Arch != "none" {
		return errors.New("action pkg:arch is unsupported")
	}

	return nil
}

func (p *Pkg) validateDescription() error {
	if p.Description == "" {
		return errors.New("action pkg:description required")
	}

	return nil
}
