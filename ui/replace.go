package ui

import (
	"github.com/ohzqq/rename/cfg"
)

const (
	find = iota
	replace
)

func FindReplaceForm() *Form {
	inputs := make([]*Input, 2)

	inputs[find] = NewInput()
	inputs[find].Init(InputProps{SetValue: cfg.SetFind})
	inputs[find].Prompt = "regex search: "

	inputs[replace] = NewInput()
	inputs[replace].Init(InputProps{SetValue: cfg.SetReplace})
	inputs[replace].Prompt = "replace with: "

	return &Form{
		inputs:  inputs,
		focused: 0,
	}
}
