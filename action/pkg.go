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

	err = p.validateVendor()
	if err != nil {
		error = multierror.Append(error, err)
	}

	err = p.validateMaintainer()
	if err != nil {
		error = multierror.Append(error, err)
	}

	err = p.validateUrl()
	if err != nil {
		error = multierror.Append(error, err)
	}

	err = p.validateLicense()
	if err != nil {
		error = multierror.Append(error, err)
	}

	err = p.validateVersion()
	if err != nil {
		error = multierror.Append(error, err)
	}

	err = p.validateIteration()
	if err != nil {
		error = multierror.Append(error, err)
	}

	err = p.validateBranch()
	if err != nil {
		error = multierror.Append(error, err)
	}

	err = p.validateVcsRevision()
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

func (p *Pkg) validateVendor() error {
	if p.Vendor == "" {
		return errors.New("action pkg:vendor required")
	}

	return nil
}

func (p *Pkg) validateMaintainer() error {
	if p.Maintainer == "" {
		return errors.New("action pkg:maintainer required")
	}

	return nil
}

func (p *Pkg) validateUrl() error {
	if p.Url == "" {
		return errors.New("action pkg:url required")
	}

	return nil
}

func (p *Pkg) validateLicense() error {
	if p.License == "" {
		return errors.New("action pkg:license required")
	}

	return nil
}

func (p *Pkg) validateVersion() error {
	if p.Version == "" {
		return errors.New("action pkg:version required")
	}

	return nil
}

func (p *Pkg) validateIteration() error {
	if p.Iteration == "" {
		return errors.New("action pkg:iteration required")
	}

	return nil
}

func (p *Pkg) validateBranch() error {
	if p.Branch == "" {
		return errors.New("action pkg:branch required")
	}

	return nil
}

func (p *Pkg) validateVcsRevision() error {
	if p.Branch == "" {
		return errors.New("action pkg:vcs_revision required")
	}

	return nil
}
