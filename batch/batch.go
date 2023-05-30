package batch

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/danielgtaylor/casing"
	"github.com/ohzqq/rename"
	"github.com/ohzqq/rename/name"
)

type Names struct {
	Files       []*name.Name
	PadPosition rename.PadPosition
	Case        rename.Casing
	Sep         string
	Min         int
	Max         int
	Pad         bool
	PadFmt      string
	Sanitize    bool
}

func New() *Names {
	name := &Names{
		PadPosition: rename.PosAfterName,
		Case:        -1,
		Min:         1,
		Sep:         "",
		PadFmt:      "%03d",
	}
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

func (b *Names) Rename(trans ...casing.TransformFunc) {
	num := b.Min
	for _, file := range b.Files {
		if b.Sep != "" {
			file.Sep(b.Sep)
		}
		name := file.Rename(trans...)
		if b.Pad {
			//name = Pad(name, b.PadFmt, num, b.PadPosition)
			num++
		}
		fmt.Printf("%s%s\n", name, file.Ext)
	}
}
