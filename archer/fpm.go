package archer

import (
	"errors"
	"fmt"
	"os/exec"
	"path"
)

type Fpm struct {
	debug  bool
	args   []string
	format string
}

func NewFpm(format string, workPath string, outputPath string, targetPath string, debug bool) (*Fpm, error) {
	_, err := exec.LookPath("fpm")
	if err != nil {
		return nil, errors.New("fpm: fpm is required to build packages, ensure that fpm is in your PATH")
	}

	fpm := &Fpm{}
	fpm.debug = debug

	switch format {
	case "deb":
		fpm.format = "deb"
		fpm.addArg("-t", "deb")
		fpm.addArg("--deb-no-default-config-files")
	case "rpm":
		fpm.format = "rpm"
		fpm.addArg("-t", "rpm")
	default:
		return nil, errors.New("fpm: format not supported")

	}

	fpm.addArg("-s", "dir")
	fpm.addArg("-p", outputPath)
	fpm.addArg("-C", targetPath)

	fpm.addArg("--after-install", path.Join(workPath, ScriptDir, "after-install"))
	fpm.addArg("--before-install", path.Join(workPath, ScriptDir, "before-install"))
	fpm.addArg("--after-remove", path.Join(workPath, ScriptDir, "after-remove"))
	fpm.addArg("--before-remove", path.Join(workPath, ScriptDir, "before-remove"))

	return fpm, nil
}

func (f *Fpm) addArg(args ...string) {
	f.args = append(f.args, args...)
}

func (f *Fpm) Arch(arch string) *Fpm {
	switch arch {
	case "x86_64":
		switch f.format {
		case "deb":
			f.addArg("-a", "amd64")
		case "rpm":
			f.addArg("-a", arch)

		}
	case "none":
		switch f.format {
		case "deb":
			f.addArg("-a", "all")
		case "rpm":
			f.addArg("-a", "noarch")

		}
	default:
		f.addArg("-a", "amd64")
	}

	return f
}

func (f *Fpm) Name(name string) *Fpm {
	f.addArg("-n", name)

	return f
}

func (f *Fpm) Description(description string) *Fpm {
	f.addArg("--description", description)

	return f
}

func (f *Fpm) Version(version string) *Fpm {
	f.addArg("-v", version)

	return f
}

func (f *Fpm) Iteration(iteration string) *Fpm {
	f.addArg("--iteration", iteration)

	return f
}

func (f *Fpm) License(license string) *Fpm {
	f.addArg("--license", license)

	return f
}

func (f *Fpm) Vendor(vendor string) *Fpm {
	f.addArg("--vendor", vendor)

	return f
}

func (f *Fpm) Maintainer(maintainer string) *Fpm {
	f.addArg("--maintainer", maintainer)

	return f
}

func (f *Fpm) Branch(branch string) *Fpm {
	switch f.format {
	case "deb":
		f.addArg("--deb-field", fmt.Sprint("Branch: ", branch))
	case "rpm":
		f.addArg("--rpm-tag", fmt.Sprint("Branch: ", branch))
	}

	return f
}

func (f *Fpm) VcsRevision(vcsRev string) *Fpm {
	switch f.format {
	case "deb":
		f.addArg("--deb-field", fmt.Sprint("VcsRevision: ", vcsRev))
	case "rpm":
		f.addArg("--rpm-tag", fmt.Sprint("VcsRevision: ", vcsRev))
	}
	return f
}

func (f *Fpm) Url(url string) *Fpm {
	f.addArg("--url", url)

	return f
}

func (f *Fpm) FileUser(user string) *Fpm {
	f.addArg(fmt.Sprint("--", f.format, "-user"), user)

	return f
}

func (f *Fpm) FileGroup(group string) *Fpm {
	f.addArg(fmt.Sprint("--", f.format, "-group"), group)

	return f
}

func (f *Fpm) Run() error {
	f.addArg(".")

	if out, err := exec.Command("fpm", f.args...).CombinedOutput(); err != nil {
		if f.debug {
			fmt.Println(string(out))
		}
		return err
	}

	return nil
}
