package name

import (
	"path/filepath"
	"strings"

	"github.com/danielgtaylor/casing"
	"github.com/spf13/viper"
)

type Name struct {
	Ext          string
	name         string
	dir          string
	Base         string
	sep          string
	Regex        string
	Search       string
	Replace      string
	Split        []string
	Cwd          bool
	MergeNumbers bool
	prefix       string
	suffix       string
	num          int
}

type TransformFunc func(*Name) string

func New(n string) *Name {
	name := &Name{
		num: 1,
		sep: "",
	}
	err := name.Parse(n)
	if err != nil {
		panic(err)
	}
	return name
}

func (fn *Name) SetName(n string) *Name {
	fn.Base = n
	return fn
}

func (fn *Name) Sep(sep string) *Name {
	fn.sep = sep
	return fn
}

func (fn *Name) Num(n int) *Name {
	fn.num = n
	return fn
}

func (fn *Name) Parse(n string) error {
	fn.dir, fn.name = filepath.Split(n)
	fn.Ext = filepath.Ext(n)
	fn.Base = strings.TrimSpace(strings.TrimSuffix(fn.name, fn.Ext))
	fn.Split = casing.Split(fn.Base)
	return nil
}

func (fn *Name) NewName() string {
	return fn.Base
}

func (name *Name) Rename(trans ...casing.TransformFunc) string {
	switch c := viper.GetInt("casing"); c {
	case Camel:
		return casing.Camel(name.Base, trans...)
	case Kebab:
		return casing.Kebab(name.Base, trans...)
	case LowerCamel:
		return casing.LowerCamel(name.Base, trans...)
	case Snake:
		return casing.Snake(name.Base, trans...)
	default:
		return casing.Join(name.Split, viper.GetString("sep"), trans...)
	}
}

const (
	Camel = iota
	Kebab
	LowerCamel
	Snake
)

type PadPosition int

const (
	PosStart = iota
	PosBeforeName
	PosAfterName
	PosEnd
)
