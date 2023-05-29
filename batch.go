package rename

import (
	"log"
	"path/filepath"
)

type Rename struct {
	Names   []*FileName
	Padding string
	Min     int
	Max     int
	Pad     bool
}

func Batch() *Rename {
	name := &Rename{
		Padding: "%03d",
	}
	return name
}

func (r *Rename) Glob(g string) *Rename {
	files, err := filepath.Glob(g)
	if err != nil {
		log.Fatal(err)
	}
	r.Files(files)
	return r
}

func (r *Rename) Files(files []string) *Rename {
	for _, file := range files {
		r.Names = append(r.Names, New(file))
	}
	return r
}

type PadPosition int

//go:generate stringer -type PadPosition -trimprefix PadPos
const (
	PadPosStart PadPosition = iota
	PadPosBeforeName
	PadPosAfterName
	PadPosEnd
)
