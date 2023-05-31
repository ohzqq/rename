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

func (p *PaddingCfg) SetZeroes(num any) *PaddingCfg {
	viper.Set("pad.zeroes", num)
	return p
}

func (p *PaddingCfg) SetStart(num any) *PaddingCfg {
	viper.Set("pad.start", num)
	return p
}

func (p *PaddingCfg) SetPosition(num any) *PaddingCfg {
	viper.Set("pad.position", num)
	return p
}
