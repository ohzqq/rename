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
	if viper.GetInt("pad") == 0 {
	}
	for _, file := range files {
		r.Files = append(r.Files, name.New(file))
	}
	return r
}

func (b *Names) Generate() []string {
	var names []string
	var trans []casing.TransformFunc

	if viper.GetBool("sanitize") || viper.GetBool("asciiify") {
		trans = append(trans, xform.Asciiify)
	}

	switch p := viper.GetInt("pad"); p {
	case 0:
		d := strconv.Itoa(len(strconv.Itoa(len(b.Files))))
		viper.Set("pad_fmt", "%0"+d+"d")
	default:
		viper.Set("pad_fmt", "%0"+strconv.Itoa(p)+"d")
	}

	num := viper.GetInt("min")
	for _, file := range b.Files {
		name := file.Build(trans...)

		if viper.IsSet("pad_fmt") {
			name = xform.Pad(name, num)
			num++
		}

		names = append(names, name+file.Ext)
	}
	return names
}
