package xform

import (
	"strings"

	"github.com/danielgtaylor/casing"
	"github.com/gosimple/unidecode"
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
	return casing.Camel(s, Asciiify)
}
