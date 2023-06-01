package xform

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/danielgtaylor/casing"
	"github.com/gosimple/unidecode"
	"github.com/ohzqq/rename/cfg"
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

func Pad(num int) string {
	pad := "%0" + strconv.Itoa(cfg.Zeroes()) + "d"
	return fmt.Sprintf(pad, num)
}

func Replace(n string) string {
	regex := cfg.Find()
	return regex.ReplaceAllString(n, cfg.Replace())
}
