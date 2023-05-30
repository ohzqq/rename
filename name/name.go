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
	base         string
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
	fn.base = n
	return fn
}

func (fn *Name) Parse(n string) error {
	fn.dir, fn.name = filepath.Split(n)
	fn.Ext = filepath.Ext(n)
	return nil
}

func (fn *Name) Base() string {
	if viper.GetBool("cwd") {
		if fn.dir != "" {
			return filepath.Base(fn.dir)
		} else {
			wd, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			return filepath.Base(wd)
		}
	}
	return strings.TrimSpace(strings.TrimSuffix(fn.name, fn.Ext))
}

func (name *Name) Transform(trans ...casing.TransformFunc) string {
	var n string
	base := name.Base()

	switch c := viper.GetInt("casing"); c {
	case Camel:
		n = casing.Camel(base, trans...)
	case Kebab:
		n = casing.Kebab(base, trans...)
	case LowerCamel:
		n = casing.LowerCamel(base, trans...)
	case Snake:
		n = casing.Snake(base, trans...)
	case Lower:
		trans = append(trans, strings.ToLower)
		n = casing.Join(casing.Split(base), viper.GetString("sep"), trans...)
	case Upper:
		trans = append(trans, strings.ToUpper)
		n = casing.Join(casing.Split(base), viper.GetString("sep"), trans...)
	default:
		n = casing.Join(casing.Split(base), viper.GetString("sep"), trans...)
	}

	if viper.IsSet("prefix") {
		n = viper.GetString("prefix") + n
	}

	if viper.IsSet("suffix") {
		n = n + viper.GetString("suffix")
	}

	return filepath.Join(name.dir, n)
}

const (
	Camel = iota
	Kebab
	LowerCamel
	Snake
	Lower
	Upper
)

const (
	PosStart = iota
	PosBeforeName
	PosAfterName
	PosEnd
)
