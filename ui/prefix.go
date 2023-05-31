package ui

import (
	"fmt"

	"github.com/ohzqq/rename/cfg"
)

func PrefixForm() *Form {
	inputs := make([]*Input, 1)
	inputs[0] = NewInput(cfg.SetPrefix)
	inputs[0].Prompt = "prefix: "
	return &Form{
		inputs:  inputs,
		focused: 0,
	}
}

func SanitizeForm() *Form {
	inputs := make([]*Input, 1)
	inputs[0] = NewInput(cfg.Sanitize)
	inputs[0].Prompt = "sanitize (y/n): "
	inputs[0].Validate = ValidateBool
	return &Form{
		inputs:  inputs,
		focused: 0,
	}
}

func ValidateBool(v string) error {
	if v != "y" {
		return fmt.Errorf("invalide")
	}
	return nil
}
