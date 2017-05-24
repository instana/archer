package archer

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/hashicorp/hil"
	"github.com/hashicorp/hil/ast"
)

type ArcherFile struct {
	path string
}

func NewArcherFile(path string) *ArcherFile {
	return &ArcherFile{path}
}

func (a *ArcherFile) variables(root ast.Node) ([]string, error) {
	var result []string
	var resultErr error

	fn := func(n ast.Node) ast.Node {
		if resultErr != nil {
			return n
		}
		switch vn := n.(type) {
		case *ast.VariableAccess:
			v := vn.Name
			result = append(result, v)
		case *ast.Index:
			if va, ok := vn.Target.(*ast.VariableAccess); ok {
				v := va.Name
				result = append(result, v)
			}
			if va, ok := vn.Key.(*ast.VariableAccess); ok {
				v := va.Name

				result = append(result, v)
			}
		default:
			return n
		}

		return n
	}

	root.Accept(fn)

	if resultErr != nil {
		return nil, resultErr
	}

	return result, nil
}

func (a *ArcherFile) Load() (string, error) {
	afbytes, err := ioutil.ReadFile(a.path)
	if err != nil {
		return "", err
	}

	tree, err := hil.Parse(string(afbytes))
	if err != nil {
		return "", err
	}

	vars, err := a.variables(tree)
	if err != nil {
		return "", err
	}

	timestamp := ast.Function{
		ArgTypes:   []ast.Type{},
		ReturnType: ast.TypeString,
		Variadic:   false,
		Callback: func(inputs []interface{}) (interface{}, error) {
			return time.Now().UTC().Format("20060102150405"), nil
		},
	}

	config := &hil.EvalConfig{
		GlobalScope: &ast.BasicScope{
			FuncMap: map[string]ast.Function{
				"timestamp": timestamp,
			},
		},
	}

	config.GlobalScope.VarMap = make(map[string]ast.Variable)

	for _, v := range vars {
		if strings.HasPrefix(v, "env.") {
			val, ok := os.LookupEnv(strings.Split(v, ".")[1])
			if !ok {
				return "", errors.New(fmt.Sprint("Archerfile no value defined for: ", v))
			} else {
				config.GlobalScope.VarMap[v] = ast.Variable{
					Type:  ast.TypeString,
					Value: val,
				}
			}

		}
	}

	result, err := hil.Eval(tree, config)
	if err != nil {
		return "", err
	}

	return result.Value.(string), nil
}
