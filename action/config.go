package action

import (
	"errors"
	"sort"

	"encoding/json"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
)

type Config struct {
	Actions Actions
	Af      *ast.File

	index map[string]int
}

type JsonConfig struct {
	Pkg         []Action `json:"pkg"`
	Requirement []Action `json:"requirement"`
}

func NewConfig() *Config {
	config := &Config{}
	config.index = make(map[string]int)
	return config
}

func (c *Config) Load(ahcl string) error {
	var err error

	c.Af, err = hcl.Parse(ahcl)
	if err != nil {
		return err
	}

	_, ok := c.Af.Node.(*ast.ObjectList)
	if !ok {
		return errors.New("Archerfile doesn't contain a root object")
	}

	err = c.LoadPkg()
	if err != nil {
		return err
	}

	err = c.LoadRequirement()
	if err != nil {
		return err
	}

	err = c.LoadBuild()
	if err != nil {
		return err
	}

	return err
}

func (c *Config) Add(action Action) {
	if index, ok := c.index[action.Unique()]; !ok {
		c.Actions = append(c.Actions, action)
		c.index[action.Unique()] = len(c.Actions) - 1
	} else {
		// Shuffle stuff around as this is an override
		action = c.Actions[index]
		c.Actions = append(c.Actions[:index], c.Actions[index+1:]...)
		c.Actions = append(c.Actions, action)
		c.index[action.Unique()] = len(c.Actions) - 1
		c.Index()
	}
}

func (c *Config) Section(filters ...string) []Action {
	var items []Action

	for _, item := range c.Actions {
		for _, filter := range filters {
			if item.Type() == filter {
				items = append(items, item)
			}
		}
	}

	return items
}

func (c *Config) Sort() {
	sort.Sort(c.Actions)

	// Rebuild the index
	c.Index()
}

func (c *Config) Index() {
	for index, action := range c.Actions {
		c.index[action.Unique()] = index
	}
}

func (c *Config) LoadPkg() error {
	actions, _ := c.Af.Node.(*ast.ObjectList)

	if pkg := actions.Filter("pkg"); len(pkg.Items) > 0 {
		for _, item := range pkg.Items {
			var object *Pkg

			err := hcl.DecodeObject(&object, item.Val)
			if err != nil {
				return err
			}

			if object.Valid() {
				c.Add(object)
			} else {
				return errors.New("invalid package action")
			}
			break
		}
	}

	return nil
}

func (c *Config) LoadBuild() error {
	actions, _ := c.Af.Node.(*ast.ObjectList)

	if build := actions.Filter("build"); len(build.Items) > 0 {
		for _, item := range build.Items {
			var object *Build

			err := hcl.DecodeObject(&object, item.Val)
			if err != nil {
				return err
			}

			if object.Valid() {
				c.Add(object)
			} else {
				return errors.New("invalid build action")
			}
			break
		}
	}

	return nil
}

func (c *Config) LoadRequirement() error {
	actions, _ := c.Af.Node.(*ast.ObjectList)

	if requirement := actions.Filter("requirement"); len(requirement.Items) > 0 {
		for _, item := range requirement.Items {
			var object *Requirement

			err := hcl.DecodeObject(&object, item.Val)
			if err != nil {
				return err
			}

			if object.Valid() {
				c.Add(object)
			} else {
				return errors.New("invalid requirement action")
			}
		}
	}

	return nil
}

func (c *Config) Json() string {
	manifest := &JsonConfig{}
	c.Sort()

	manifest.Pkg = c.Section("pkg")

	requirement := c.Section("requirement")
	if len(requirement) > 0 {
		manifest.Requirement = requirement
	}

	out, _ := json.Marshal(manifest)

	return string(out)
}
