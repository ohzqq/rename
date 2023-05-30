package batch

import (
	"log"
	"path/filepath"
	"strconv"

	"github.com/danielgtaylor/casing"
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
	if viper.GetInt("pad") == 0 {
		viper.Set("pad", len(strconv.Itoa(len(r.Files))))
	}
	return r
}

func (b *Names) Transform() []map[string]string {
	var names []map[string]string
	var trans []casing.TransformFunc

	if viper.GetBool("sanitize") || viper.GetBool("asciiify") {
		trans = append(trans, xform.Asciiify)
	}

	num := viper.GetInt("min")
	for _, file := range b.Files {
		name := file.Transform(trans...)

		if p := viper.GetInt("pad"); p >= 0 {
			name = xform.Pad(name, num)
			num++
		}

		name = name + file.Ext

		if viper.IsSet("find") {
			name = xform.Replace(name)
		}

		names = append(names, map[string]string{file.Original: name})
	}
	return names
}
