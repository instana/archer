package archer

import (
	"errors"
	"io/ioutil"
	"os"
	"path"

	"github.com/chuckpreslar/emission"
	"github.com/instana/archer/action"
)

type Builder struct {
	*emission.Emitter

	debug bool

	afPath     string
	workPath   string
	outputPath string
	targetPath string

	buildRpm bool
	buildDeb bool

	fileUser  string
	fileGroup string

	config *action.Config
}

func NewBuilder() *Builder {
	builder := &Builder{Emitter: emission.NewEmitter()}

	builder.config = action.NewConfig()

	return builder
}

func (b *Builder) Debug(debug bool) *Builder {
	b.debug = debug
	return b
}

func (b *Builder) AfPath(af string) *Builder {
	b.afPath = af
	return b
}

func (b *Builder) WorkPath(wp string) *Builder {
	b.workPath = wp
	return b
}

func (b *Builder) TargetPath(tp string) *Builder {
	b.targetPath = tp
	return b
}

func (b *Builder) OutputPath(op string) *Builder {
	b.outputPath = op
	return b
}

// Set default paths
func (b *Builder) setPaths() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	if b.afPath == "" {
		b.afPath = path.Join(wd, DefaultAfPath)
	}

	// If the path is a directory append the default
	// Archerfile path
	stat, err := os.Stat(b.afPath)
	if err != nil {
		return err
	}
	if stat.IsDir() {
		b.afPath = path.Join(b.afPath, DefaultAfPath)
	}

	if b.workPath == "" {
		b.workPath = wd
	}
	if b.outputPath == "" {
		b.outputPath = path.Join(wd, DefaultOutPath)
	}
	if b.targetPath == "" {
		b.targetPath = path.Join(wd, DefaultTargetPath)
	}

	return err
}

func (b *Builder) loadArcherfile() error {
	af := NewArcherFile(b.afPath)

	afResult, err := af.Load()

	err = b.config.Load(afResult)
	if err != nil {
		return err
	}

	return err
}

func (b *Builder) setBuild() {
	filter := b.config.Section("build")
	if len(filter) == 0 {
		return
	}

	build := filter[0].(*action.Build)

	if build.WorkPath != "" {
		b.workPath = build.WorkPath
	}

	if build.OutPath != "" {
		b.outputPath = build.OutPath
	}

	if build.TargetPath != "" {
		b.targetPath = build.TargetPath
	}

	b.buildDeb = build.Deb
	b.buildRpm = build.Rpm

	if b.fileUser == "" {
		b.fileUser = build.FileUser
	}

	if b.fileGroup == "" {
		b.fileGroup = build.FileGroup
	}

}

func (b *Builder) writeScripts() error {
	var err error

	hooks := []string{
		"before-install",
		"after-install",
		"before-remove",
		"after-remove",
	}

	scriptDir := path.Join(b.workPath, ScriptDir)
	if _, err = os.Stat(scriptDir); os.IsNotExist(err) {
		os.Mkdir(scriptDir, os.FileMode(0750))
	}

	for _, hook := range hooks {
		content := "#!/bin/bash\n" +
			"which archer > /dev/null || exit 0\n" +
			"archer hook " +
			hook +
			" << EOF" +
			"\n" +
			b.config.Json() +
			"\n" +
			"EOF" +
			"\n"

		err = ioutil.WriteFile(path.Join(scriptDir, hook), []byte(content), 0640)
		if err != nil {
			return err
		}
	}

	return err
}

func (b *Builder) writePackageConf(pkg string) error {
	var err error

	confDir := path.Join(b.targetPath, ConfPath)

	if _, err = os.Stat(confDir); os.IsNotExist(err) {
		os.MkdirAll(confDir, os.FileMode(0750))
	}

	err = ioutil.WriteFile(path.Join(confDir, pkg+".conf"), []byte(b.config.Json()+"\n"), 0644)
	if err != nil {
		return err
	}

	return err
}

func (b *Builder) Build() error {
	err := b.setPaths()
	if err != nil {
		return err
	}

	err = b.loadArcherfile()
	if err != nil {
		return err
	}

	b.setBuild()

	err = b.writeScripts()
	if err != nil {
		return err
	}

	filter := b.config.Section("pkg")
	if len(filter) == 0 {
		return errors.New("builder: no pkg definition found")
	}
	pkg := filter[0].(*action.Pkg)

	err = b.writePackageConf(pkg.Name)
	if err != nil {
		return err
	}

	if _, err = os.Stat(b.outputPath); os.IsNotExist(err) {
		os.Mkdir(b.outputPath, os.FileMode(0750))
	}

	if b.buildRpm == true {
		rpmBuilder, err := NewFpm("rpm", b.workPath, b.outputPath, b.targetPath, b.debug)
		if err != nil {
			return err
		}

		rpmBuilder.Name(pkg.Name).
			FileGroup(b.fileGroup).
			FileUser(b.fileGroup).
			Arch(pkg.Arch).
			Version(pkg.Version).
			Iteration(pkg.Iteration).
			Description(pkg.Description).
			Vendor(pkg.Vendor).
			Maintainer(pkg.Maintainer).
			Url(pkg.Url).
			License(pkg.License)

		err = rpmBuilder.Run()
		if err != nil {
			return errors.New("fpm: rpm build failed")
		}
	}

	if b.buildDeb == true {
		debBuilder, err := NewFpm("deb", b.workPath, b.outputPath, b.targetPath, b.debug)
		if err != nil {
			return err
		}

		debBuilder.Name(pkg.Name).
			FileGroup(b.fileGroup).
			FileUser(b.fileGroup).
			Arch(pkg.Arch).
			Version(pkg.Version).
			Iteration(pkg.Iteration).
			Description(pkg.Description).
			Vendor(pkg.Vendor).
			Maintainer(pkg.Maintainer).
			Url(pkg.Url).
			License(pkg.License)

		err = debBuilder.Run()
		if err != nil {
			return errors.New("fpm: deb build failed")
		}
	}

	return err
}