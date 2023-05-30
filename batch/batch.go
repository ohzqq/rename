package batch

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/danielgtaylor/casing"
	"github.com/ohzqq/rename/name"
	"github.com/spf13/viper"
)

type Names struct {
	Files       []*name.Name
	PadPosition name.PadPosition
	Case        name.Casing
	Sep         string
	Min         int
	Max         int
	Pad         bool
	PadFmt      string
	Sanitize    bool
}

func New() *Names {
	name := &Names{
		PadPosition: name.PosAfterName,
		Case:        -1,
		Min:         1,
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
	num := viper.GetInt("min")
	for _, file := range b.Files {
		name := file.Rename(trans...)
		if viper.GetBool("pad") {
			name = Pad(name, num)
			num++
		}
		fmt.Printf("%s%s\n", name, file.Ext)
	}
}

func Pad(in string, num int) string {
	pos := viper.GetInt("pad_position")
	switch pos {
	case name.PosStart, name.PosBeforeName:
		return fmt.Sprintf(
			viper.GetString("pad_fmt")+"%s",
			num,
			in,
		)
	case name.PosEnd, name.PosAfterName:
		return fmt.Sprintf(
			"%s"+viper.GetString("pad_fmt"),
			in,
			num,
		)
	}
	return in
}
