package ui

import (
	"fmt"

	"github.com/ohzqq/rename/cfg"
)

func MiscForm() *Form {
	inputs := make([]*Input, 3)

	inputs[0] = NewInput(cfg.SetPrefix)
	inputs[0].Prompt = "prefix: "

	inputs[1] = NewInput(cfg.SetSuffix)
	inputs[1].Prompt = "suffix: "

	inputs[2] = NewInput(cfg.Sanitize)
	inputs[2].Prompt = "sanitize (y/n): "
	inputs[2].Validate = ValidateBool

	return &Form{
		inputs:  inputs,
		focused: 0,
	}
}

func ValidateBool(v string) error {
	if v != "y" {
		return fmt.Errorf("invalid")
	}
	return nil
}
