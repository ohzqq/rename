package batch

import (
	"fmt"
	"log"
	"path/filepath"
	"strconv"

	"github.com/danielgtaylor/casing"
	"github.com/ohzqq/rename/cfg"
	"github.com/ohzqq/rename/name"
	"github.com/ohzqq/rename/xform"
	"github.com/spf13/viper"
)

type Names struct {
	Files []*name.Name
}

func New() *Names {
	name := &Names{}
	return name
}

func Glob(dir string) *Names {
	return New().Glob(dir)
}

func Files(files []string) *Names {
	return New().SetFiles(files)
}

func (r *Names) Glob(g string) *Names {
	files, err := filepath.Glob(g)
	if err != nil {
		log.Fatal(err)
	}
	r.SetFiles(files)
	return r
}

func (r *Names) SetFiles(files []string) *Names {
	for _, file := range files {
		r.Files = append(r.Files, name.New(file))
	}
	if cfg.Padding().Zeroes == 0 {
		cfg.Padding().SetZeroes(len(strconv.Itoa(len(r.Files))))
	}
	return r
}

func (b *Names) Transform() []map[string]string {
	var names []map[string]string
	var trans []casing.TransformFunc

	if viper.GetBool("sanitize") || viper.GetBool("asciiify") {
		trans = append(trans, xform.Asciiify)
	}

	num := cfg.Padding().Start
	for _, file := range b.Files {
		n := file.Transform(trans...)

		if viper.IsSet("find") {
			n = xform.Replace(n)
		}

		var padding string
		if p := cfg.Padding().Zeroes; p >= 0 {
			padding = xform.Pad(num)
			num++
		}

		var pre string
		if viper.IsSet("prefix") {
			pre = viper.GetString("prefix")
		}

		var suf string
		if viper.IsSet("suffix") {
			suf = viper.GetString("suffix")
		}

		switch pos := cfg.Padding().Position; name.PadPosition(pos) {
		case name.PosBeginning:
			n = fmt.Sprint(padding, pre, n, suf)
		case name.PosBeforeName:
			n = fmt.Sprint(pre, padding, n, suf)
		case name.PosEnd:
			n = fmt.Sprint(pre, n, suf, padding)
		case name.PosAfterName:
			n = fmt.Sprint(pre, n, padding, suf)
		}

		n = filepath.Join(file.Dir, n+file.Ext)

		names = append(names, map[string]string{file.Original: n})
	}
	return names
}
