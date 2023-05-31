package cfg

import "github.com/spf13/viper"

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

func (p *PaddingCfg) SetZeroes(num any) {
	viper.Set("pad.zeroes", num)
}

func (p *PaddingCfg) SetStart(num any) {
	viper.Set("pad.start", num)
}

func (p *PaddingCfg) SetPosition(num any) {
	viper.Set("pad.position", num)
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
