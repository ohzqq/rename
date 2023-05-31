package ui

import "github.com/ohzqq/rename/cfg"

func SuffixForm() *Form {
	inputs := make([]*Input, 1)
	inputs[0] = NewInput(cfg.SetSuffix)
	inputs[0].Prompt = "suffix: "
	return &Form{
		inputs:  inputs,
		focused: 0,
	}
}
