package cfg

import (
	"log"
	"regexp"

	"github.com/ohzqq/rename/opt"
	"github.com/spf13/viper"
)

func Zeroes() int {
	return viper.GetInt(opt.Zeroes)
}

func SetZeroes(num any) {
	viper.Set(opt.Zeroes, num)
}

func Start() int {
	return viper.GetInt(opt.Start)
}

func SetStart(num any) {
	viper.Set(opt.Start, num)
}

func Position() string {
	return viper.GetString(opt.Position)
}

func SetPosition(num any) {
	var pos string
	switch n := num.(string); n {
	case "0":
		pos = opt.Beginning
	case "1":
		pos = opt.BeforeName
	case "2":
		pos = opt.AfterName
	case "3":
		pos = opt.End
	}
	viper.Set(opt.Position, pos)
}

func Sep() string {
	return viper.GetString(opt.Sep)
}

func SetSep(c any) {
	viper.Set(opt.Sep, c)
}

func Find() *regexp.Regexp {
	regex, err := regexp.Compile(viper.GetString(opt.Find))
	if err != nil {
		log.Fatal(err)
	}
	return regex
}

func SetFind(f any) {
	viper.Set(opt.Find, f)
}

func Replace() string {
	return viper.GetString(opt.Replace)
}

func SetReplace(f any) {
	viper.Set(opt.Replace, f)
}

func Case() string {
	return viper.GetString(opt.Casing)
}

func SetCase(cs any) {
	var pos string
	switch n := cs.(string); n {
	case "0", opt.Camel:
		pos = opt.Camel
	case "1", opt.Kebab:
		pos = opt.Kebab
	case "2", opt.LowerCamel:
		pos = opt.LowerCamel
	case "3", opt.Snake:
		pos = opt.Snake
	}
	viper.Set(opt.Casing, pos)
}

func Suffix() string {
	return viper.GetString(opt.Suffix)
}

func SetSuffix(c any) {
	viper.Set(opt.Suffix, c)
}

func Prefix() string {
	return viper.GetString(opt.Prefix)
}

func SetPrefix(c any) {
	viper.Set(opt.Prefix, c)
}

func Tidy(c any) {
	SetCase(opt.Snake)
	viper.Set(opt.Tidy, true)
}

func UseDir(c any) {
	viper.Set(opt.Dir, true)
}

func NewName(c any) {
	viper.Set(opt.Name, c)
}

func Name() string {
	return viper.GetString(opt.Name)
}

func ToUpper(c any) {
	viper.Set(opt.Upper, true)
}

func ToLower(c any) {
	viper.Set(opt.Lower, true)
}
