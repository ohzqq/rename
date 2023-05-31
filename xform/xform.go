package xform

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/danielgtaylor/casing"
	"github.com/gosimple/unidecode"
	"github.com/ohzqq/rename/name"
	"github.com/spf13/viper"
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
	pad := "%0" + strconv.Itoa(viper.GetInt("zeroes")) + "d"
	switch pos := viper.GetInt("pad_position"); name.PadPosition(pos) {
	case name.PosStart, name.PosBeforeName:
		pad = pad + "%s"
	case name.PosEnd, name.PosAfterName:
		pad = "%s" + pad
	}
	return fmt.Sprintf(pad, in, num)
}

func Replace(n string) string {
	regex, err := regexp.Compile(viper.GetString("find"))
	if err != nil {
		log.Fatal(err)
	}
	return regex.ReplaceAllString(n, viper.GetString("replace"))
}
