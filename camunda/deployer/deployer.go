package deployer

import (
	"fmt"
	"os"
	"path/filepath"
)

// Bpmn -
type Bpmn struct {
	Path string
}

// Workflows -
type Workflows struct {
	Root      string
	Extension string
}

func (w *Workflows) ext(path string) bool {
	return filepath.Ext(path) == fmt.Sprintf(".%s", w.Extension)
}

// Itens -
func (w *Workflows) Itens() []Bpmn {
	itens := make([]Bpmn, 0)

	err := filepath.Walk(w.Root, func(path string, info os.FileInfo, err error) error {
		if w.ext(path) {
			fullPath, _ := filepath.Abs(path)
			bpmn := Bpmn{
				Path: fullPath,
			}

			itens = append(itens, bpmn)
		}
		return nil
	})

	if err != nil {
		panic(err)
	}

	return itens
}

// New -
func New(params ...string) *Workflows {
	root, _ := filepath.Abs("../bpmn")

	if len(params) == 1 {
		root = params[0]
	}

	return &Workflows{
		Root:      root,
		Extension: "bpmn",
	}
}
