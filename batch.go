package rename

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/danielgtaylor/casing"
)

type Batch struct {
	Names       []*Name
	PadPosition PadPosition
	Case        Casing
	Sep         string
	Min         int
	Max         int
	Pad         bool
	Padding     string
	Sanitize    bool
}

func Rename() *Batch {
	name := &Batch{
		PadPosition: PosAfterName,
		Case:        -1,
		Min:         1,
		Sep:         "",
		Padding:     "%03d",
	}
	return name
}

func (r *Batch) Glob(g string) *Batch {
	files, err := filepath.Glob(g)
	if err != nil {
		log.Fatal(err)
	}
	r.Files(files)
	return r
}

func (r *Batch) Files(files []string) *Batch {
	for _, file := range files {
		r.Names = append(r.Names, New(file))
	}
	return r
}

func (b *Batch) Rename(trans ...casing.TransformFunc) {
	num := b.Min
	for _, file := range b.Names {
		file.Sep(b.Sep)
		name := file.Rename(trans...)
		if b.Pad {
			name = Pad(name, b.Padding, num, b.PadPosition)
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
