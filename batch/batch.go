package batch

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

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
	return r
}

func (b *Names) Rename() {
	var names []string
	var trans []casing.TransformFunc

	if viper.GetBool("sanitize") || viper.GetBool("asciiify") {
		trans = append(trans, xform.Asciiify)
	}

	num := viper.GetInt("min")
	for _, file := range b.Files {
		name := file.Rename(trans...)

		if viper.GetBool("pad") {
			name = xform.Pad(name, num)
			num++
		}

		names = append(names, name+file.Ext)
	}
	fmt.Printf("%v\n", strings.Join(names, "\n"))
}
