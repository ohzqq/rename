package name

import (
	"log"
	"os"
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
	Regex        string
	Search       string
	Replace      string
	Split        []string
	Cwd          bool
	MergeNumbers bool
	prefix       string
	suffix       string
}

func New(n string) *Name {
	name := &Name{}
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

func (fn *Name) Parse(n string) error {
	fn.dir, fn.name = filepath.Split(n)
	fn.Ext = filepath.Ext(n)
	fn.Base = strings.TrimSpace(strings.TrimSuffix(fn.name, fn.Ext))
	var err error
	if viper.GetBool("cwd") {
		viper.Set("pad", true)
		var wd string
		if fn.dir != "" {
			wd = fn.dir
		} else {
			wd, err = os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
		}
		fn.Base = filepath.Base(wd)
	}
	return nil
}

func (name *Name) Rename(trans ...casing.TransformFunc) string {
	var n string
	base := name.Base

	switch c := viper.GetInt("casing"); c {
	case Camel:
		n = casing.Camel(base, trans...)
	case Kebab:
		n = casing.Kebab(base, trans...)
	case LowerCamel:
		n = casing.LowerCamel(base, trans...)
	case Snake:
		n = casing.Snake(base, trans...)
	default:
		n = casing.Join(casing.Split(base), viper.GetString("sep"), trans...)
	}

	return filepath.Join(name.dir, n)
}

const (
	Camel = iota
	Kebab
	LowerCamel
	Snake
)

const (
	PosStart = iota
	PosBeforeName
	PosAfterName
	PosEnd
)
