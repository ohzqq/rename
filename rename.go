package rename

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/danielgtaylor/casing"
)

type FileName struct {
	Ext          string
	name         string
	dir          string
	Base         string
	Sep          string
	Regex        string
	NewName      string
	Search       string
	Replace      string
	Split        []string
	Case         Casing
	Sanitize     bool
	Cwd          bool
	PadPosition  PadPosition
	MergeNumbers bool
	Prefix       string
	Suffix       string
	Num          int
}

func New(n string) *FileName {
	name := &FileName{
		Num:         1,
		PadPosition: -1,
		Case:        -1,
	}
	err := name.Parse(n)
	if err != nil {
		panic(err)
	}
	return name
}

func (fn *FileName) SetName(n string) *FileName {
	fn.NewName = n
	return fn
}

func (fn *FileName) Parse(n string) error {
	fn.dir, fn.name = filepath.Split(n)
	fn.Ext = filepath.Ext(n)
	fn.Base = strings.TrimSuffix(fn.name, fn.Ext)
	fn.Split = casing.Split(fn.Base)
	fn.NewName = n
	return nil
}

func (name *FileName) Format(opts ...Option) (string, error) {
	if name.name == name.NewName {
		return "", fmt.Errorf("old name (%s) is the same as new name (%s)\n", name.name, name.NewName)
	}

	for _, opt := range opts {
		opt(name)
	}
	name.NewName = name.Base

	var buf bytes.Buffer
	err := nameTmpl.Execute(&buf, name)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

var nameTmpl = template.Must(template.New("name").Parse(`
{{- if eq .PadPosition 0}}{{ printf .Padding .Num }}{{end -}}
{{- with .Prefix }}{{.}}{{end -}}
{{- if eq .PadPosition 1}}{{ printf .Padding .Num }}{{end -}}
{{- with .NewName }}{{.}}{{end -}}
{{- if eq .PadPosition 2}}{{ printf .Padding .Num }}{{end -}}
{{- with .Suffix }}{{.}}{{end -}}
{{- if eq .PadPosition 4}}{{ printf .Padding .Num }}{{end -}}
{{- with .Ext }}{{.}}{{end -}}
`))

//go:generate stringer -type Casing
type Casing int

const (
	Camel Casing = iota
	Kebab
	LowerCamel
	Snake
)
