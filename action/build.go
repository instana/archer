package action

import (
	"strings"
)

type Build struct {
	TargetPath string `hcl:"target_path"`
	WorkPath   string `hcl:"work_path"`
	OutPath    string `hcl:"out_path"`
	FileUser   string `hcl:"file_user"`
	FileGroup  string `hcl:"file_group"`
	Rpm        bool
	Deb        bool
}

func NewBuild() *Build {
	build := &Build{}
	build.FileUser = "root"
	build.FileGroup = "root"
	build.Rpm = true
	build.Deb = true
	return build
}

func (b *Build) Key() string {
	key := []string{b.Type(), b.Type()}
	return strings.Join(key, ":")
}

func (b *Build) Columns() string {
	return strings.Join([]string{
		strings.ToUpper(b.Type()),
		b.WorkPath,
		b.OutPath,
		b.TargetPath,
	}, "|")
}

func (b *Build) Unique() string {
	key := []string{b.Type(), b.Type()}
	return strings.Join(key, ":")
}

func (b *Build) Type() string {
	return "build"
}

func (b *Build) Valid() bool {
	return true
}
