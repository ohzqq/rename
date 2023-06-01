package ui

import (
	"fmt"

	"github.com/ohzqq/rename/cfg"
)

func MiscForm() []*Input {
	inputs := make([]*Input, 3)

	inputs[0] = NewInput()
	inputs[0].Init(InputProps{SetValue: cfg.SetPrefix})
	inputs[0].Prompt = "prefix: "

	inputs[1] = NewInput()
	inputs[1].Init(InputProps{SetValue: cfg.SetSuffix})
	inputs[1].Prompt = "suffix: "

	inputs[2] = NewInput()
	inputs[2].Init(InputProps{SetValue: cfg.Sanitize})
	inputs[2].Prompt = "sanitize (y/n): "
	inputs[2].Validate = ValidateBool

	return inputs
}

func ValidateBool(v string) error {
	if v != "y" {
		return fmt.Errorf("invalid")
	}
	return nil
}
