package batch

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/danielgtaylor/casing"
)

type Names struct {
	Names       []*Name
	PadPosition PadPosition
	Case        Casing
	Sep         string
	Min         int
	Max         int
	Pad         bool
	PadFmt      string
	Sanitize    bool
}

func Rename() *Names {
	name := &Names{
		PadPosition: PosAfterName,
		Case:        -1,
		Min:         1,
		Sep:         "",
		PadFmt:      "%03d",
	}
	return name
}

func (r *Names) Glob(g string) *Names {
	files, err := filepath.Glob(g)
	if err != nil {
		log.Fatal(err)
	}
	r.Files(files)
	return r
}

func (r *Names) Files(files []string) *Names {
	for _, file := range files {
		r.Names = append(r.Names, New(file))
	}
	return r
}

func (b *Names) Rename(trans ...casing.TransformFunc) {
	num := b.Min
	for _, file := range b.Names {
		if b.Sep != "" {
			file.Sep(b.Sep)
		}
		name := file.Rename(trans...)
		if b.Pad {
			name = Pad(name, b.PadFmt, num, b.PadPosition)
			num++
		}
		fmt.Printf("%s%s\n", name, file.Ext)
	}
}

type PadPosition int

//go:generate stringer -type PadPosition -trimprefix PadPos
const (
	PosStart PadPosition = iota
	PosBeforeName
	PosAfterName
	PosEnd
)
