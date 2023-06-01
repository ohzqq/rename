package cfg

import (
	"log"
	"regexp"

	"github.com/ohzqq/rename/opt"
	"github.com/spf13/viper"
)

type PaddingCfg struct {
	Zeroes   int
	Start    int
	Position int
}

func Padding() *PaddingCfg {
	return &PaddingCfg{
		Zeroes:   viper.GetInt(opt.Zeroes),
		Start:    viper.GetInt(opt.Start),
		Position: viper.GetInt(opt.Position),
	}
}

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

func Position() opt.PadPosition {
	return opt.PadPosition(viper.GetInt(opt.Start))
}

func SetPosition(num any) {
	viper.Set(opt.Position, num.(int))
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

func Case() int {
	return viper.GetInt(opt.Casing)
}

func SetCase(c any) {
	viper.Set(opt.Casing, c.(int))
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

func Sanitize(c any) {
	viper.Set(opt.Clean, true)
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
