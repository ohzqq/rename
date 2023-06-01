package cfg

import (
	"log"
	"regexp"

	"github.com/spf13/viper"
)

type PaddingCfg struct {
	Zeroes   int
	Start    int
	Position int
}

func Padding() *PaddingCfg {
	return &PaddingCfg{
		Zeroes:   viper.GetInt("pad.zeroes"),
		Start:    viper.GetInt("pad.start"),
		Position: viper.GetInt("pad.position"),
	}
}

func SetZeroes(num any) {
	viper.Set("pad.zeroes", num)
}

func SetStart(num any) {
	viper.Set("pad.start", num)
}

func SetPosition(num any) {
	viper.Set("pad.position", num)
}

func Sep() string {
	return viper.GetString("sep")
}

func SetSep(c any) {
	viper.Set("sep", c)
}

func Find() *regexp.Regexp {
	regex, err := regexp.Compile(viper.GetString("find"))
	if err != nil {
		log.Fatal(err)
	}
	return regex
}

func SetFind(f any) {
	viper.Set("find", f)
}

func Replace() string {
	return viper.GetString("replace")
}

func SetReplace(f any) {
	viper.Set("replace", f)
}

func Case() int {
	return viper.GetInt("casing")
}

func SetCase(c any) {
	viper.Set("casing", c)
}

func SetSuffix(c any) {
	viper.Set("suffix", c)
}

func SetPrefix(c any) {
	viper.Set("prefix", c)
}

func Sanitize(c any) {
	viper.Set("sanitize", true)
}

func UseCWD(c any) {
	viper.Set("cwd", true)
}

func NewName(c any) {
	viper.Set("name", c)
}
