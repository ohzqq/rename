package rename

import (
	"github.com/danielgtaylor/casing"
	"github.com/gosimple/unidecode"
)

type Option func(*FileName)

func Case(c string) Option {
	return func(fn *FileName) {
		switch c {
		case "Camel", "camel", "c":
			fn.Base = casing.Camel(fn.Base)
		case "Kebab", "kebab", "k":
			fn.Base = casing.Kebab(fn.Base)
		case "lowerCamel", "lowercamel", "Lowercamel", "LowerCamel", "lc":
			fn.Base = casing.LowerCamel(fn.Base)
		case "snake", "Snake", "s":
			fn.Base = casing.Snake(fn.Base)
		}
	}
}

func Prefix(pre string) Option {
	return func(fn *FileName) {
		fn.Prefix = pre
	}
}

func Suffix(suf string) Option {
	return func(fn *FileName) {
		fn.Suffix = suf
	}
}

func Join(s string) Option {
	return func(fn *FileName) {
		fn.Base = casing.Join(fn.Split, s)
	}
}

func Ascii() Option {
	return func(fn *FileName) {
		var s []string
		for _, w := range fn.Split {
			if dec := unidecode.Unidecode(w); dec != "" {
				s = append(s, casing.Split(dec)...)
			}
		}
		fn.Split = s
	}
}
