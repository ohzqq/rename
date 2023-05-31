package ui

import "github.com/ohzqq/rename/cfg"

func PrefixForm() *Form {
	inputs := make([]*Input, 1)
	inputs[0] = NewInput(cfg.SetPrefix)
	inputs[0].Prompt = "prefix: "
	return &Form{
		inputs:  inputs,
		focused: 0,
	}
}
