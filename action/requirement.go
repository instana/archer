package action

import (
	"errors"
	"strings"

	"github.com/hashicorp/go-multierror"
)

type Requirement struct {
	Name      string `json:"name"`
	Method    string `json:"method"`
	Operation string `json:"operation"`
	Version   string `json:"version"`
}

func NewRequirement() *Requirement {
	return &Requirement{}
}

func (r *Requirement) Key() string {
	key := []string{r.Name, r.Method, r.Operation}
	return strings.Join(key, ":")
}

func (r *Requirement) Columns() string {
	return strings.Join([]string{
		strings.ToUpper(r.Type()),
		r.Name,
		r.Method,
		r.Operation,
		r.Version,
	}, "|")
}

func (r *Requirement) Unique() string {
	key := []string{"requirement", r.Name, r.Method, r.Operation}
	return strings.Join(key, ":")
}

func (r *Requirement) Type() string {
	return "requirement"
}

func (r *Requirement) Valid() bool {
	err := r.Validate()
	if err != nil {
		return false
	}

	return true
}

func (r *Requirement) Validate() error {
	var error *multierror.Error

	err := r.validateName()
	if err != nil {
		error = multierror.Append(error, err)
	}

	err = r.validateMethod()
	if err != nil {
		error = multierror.Append(error, err)
	}

	return error.ErrorOrNil()
}

func (r *Requirement) validateName() error {
	if r.Name == "" {
		return errors.New("action requirement:name required")
	}

	return nil
}

func (r *Requirement) validateMethod() error {

	if r.Method == "" {
		return errors.New("action requirement:method required")
	}

	if r.Method != "depends" {
		return errors.New("action requirement:method not supported")
	}

	return nil
}

func (r *Requirement) validateOperation() error {

	if r.Operation == "" {
		return errors.New("action requirement:operation required")
	}

	if r.Operation != "ANY" {
		return errors.New("action requirement:operation not supported")
	}

	return nil
}
