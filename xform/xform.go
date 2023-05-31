package xform

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/danielgtaylor/casing"
	"github.com/gosimple/unidecode"
	"github.com/ohzqq/rename/cfg"
	"github.com/ohzqq/rename/name"
)

func Asciiify(s string) string {
	var ascii []string
	for _, w := range casing.Split(unidecode.Unidecode(s)) {
		ascii = append(ascii, casing.Split(w)...)
	}
	return strings.Join(ascii, "")
}

func Sanitize(s string) string {
	return casing.Snake(s, Asciiify)
}

func Pad(in string, num int) string {
	pad := "%0" + strconv.Itoa(cfg.Padding().Zeroes) + "d"
	switch pos := cfg.Padding().Position; name.PadPosition(pos) {
	case name.PosStart, name.PosBeforeName:
		pad = pad + "%s"
	case name.PosEnd, name.PosAfterName:
		pad = "%s" + pad
	}
	return fmt.Sprintf(pad, in, num)
}

func Replace(n string) string {
	regex := cfg.Find()
	return regex.ReplaceAllString(n, cfg.Replace())
}
