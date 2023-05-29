package yyutils

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/danielgtaylor/casing"
	"github.com/gosimple/unidecode"
)

type FileName struct {
	Ext          string
	name         string
	Base         string
	Padding      string
	Sep          string
	Regex        string
	NewName      string
	Search       string
	Replace      string
	Split        []string
	PadPosition  PadPosition
	Case         Case
	Pad          bool
	Sanitize     bool
	Cwd          bool
	MergeNumbers bool
	Prefix       string
	Suffix       string
	Num          int
	Min          int
	Max          int
}

type RenameOpt func(*FileName)

func Rename() *FileName {
	name := &FileName{
		Padding:     "%03d",
		Num:         1,
		PadPosition: -1,
		Case:        -1,
	}
	return name
}

func (fn *FileName) Parse(n string) error {
	fn.name = unidecode.Unidecode(filepath.Base(n))
	fn.Ext = filepath.Ext(n)
	fn.NewName = fn.name

	if fn.Cwd {
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}
		fn.NewName = cwd
	}

	fn.Base = strings.TrimSuffix(fn.name, fn.Ext)
	fn.Split = casing.Split(fn.name)

	return nil
}

func (name FileName) Format(opts ...RenameOpt) (string, error) {
	if name.name == name.NewName {
		return "", fmt.Errorf("old name (%s) is the same as new name (%s)\n", name.name, name.NewName)
	}

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
{{- with .Sep }}{{.}}{{end -}}
{{- if eq .PadPosition 1}}{{ printf .Padding .Num }}{{end -}}
{{- with .NewName }}{{.}}{{end -}}
{{- if eq .PadPosition 2}}{{ printf .Padding .Num }}{{end -}}
{{- with .Sep }}{{.}}{{end -}}
{{- with .Suffix }}{{.}}{{end -}}
{{- if eq .PadPosition 4}}{{ printf .Padding .Num }}{{end -}}
{{- with .Ext }}{{.}}{{end -}}
`))

type PadPosition int

//go:generate stringer -type PadPosition -trimprefix PadPos
const (
	PadPosStart PadPosition = iota
	PadPosBeforeName
	PadPosAfterName
	PadPosEnd
)

//go:generate stringer -type Case
type Case int

const (
	Camel Case = iota
	Kebab
	LowerCamel
	Snake
)
