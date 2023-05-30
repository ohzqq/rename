package xform

import (
	"fmt"
	"strings"

	"github.com/danielgtaylor/casing"
	"github.com/gosimple/unidecode"
	"github.com/ohzqq/rename/name"
	"github.com/spf13/viper"
)

//func Case(c string) Option {
//  return func(fn *Batch) {
//    switch c {
//    case "Camel", "camel", "c":
//      fn.Case = Camel
//    case "Kebab", "kebab", "k":
//      fn.Case = Camel
//    case "lowerCamel", "lowercamel", "Lowercamel", "LowerCamel", "lc":
//      fn.Case = Camel
//    case "snake", "Snake", "s":
//      fn.Case = Camel
//    }
//  }
//}

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
	var pad string
	switch pos := viper.GetInt("pad_position"); pos {
	case name.PosStart, name.PosBeforeName:
		pad = viper.GetString("pad_fmt") + "%s"
	case name.PosEnd, name.PosAfterName:
		pad = "%s" + viper.GetString("pad_fmt")
	}
	return fmt.Sprintf(pad, in, num)
}
