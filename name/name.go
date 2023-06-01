package name

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/danielgtaylor/casing"
	"github.com/ohzqq/rename/cfg"
	"github.com/ohzqq/rename/opt"
	"github.com/spf13/viper"
)

type Name struct {
	Ext          string
	name         string
	Original     string
	Dir          string
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
	name := &Name{
		Original: n,
	}
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
	fn.Dir, fn.name = filepath.Split(n)
	fn.Ext = filepath.Ext(n)
	return nil
}

func (fn *Name) Base() string {
	if viper.IsSet("name") {
		return cfg.Name()
	}
	if viper.GetBool(opt.Dir) {
		if fn.Dir != "" {
			return filepath.Base(fn.Dir)
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

	switch c := cfg.Case(); c {
	case opt.Camel:
		n = casing.Camel(base, trans...)
	case opt.Kebab:
		n = casing.Kebab(base, trans...)
	case opt.LowerCamel:
		n = casing.LowerCamel(base, trans...)
	case opt.Snake:
		n = casing.Snake(base, trans...)
	case opt.Lower:
		trans = append(trans, strings.ToLower)
		n = casing.Join(casing.Split(base), cfg.Sep(), trans...)
	case opt.Upper:
		trans = append(trans, strings.ToUpper)
		n = casing.Join(casing.Split(base), cfg.Sep(), trans...)
	default:
		n = casing.Join(casing.Split(base), cfg.Sep(), trans...)
	}
	return n
}
