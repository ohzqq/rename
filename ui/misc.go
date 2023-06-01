package ui

import (
	"fmt"

	"github.com/ohzqq/rename/cfg"
)

func DefaultForm() []*Input {
	inputs := make([]*Input, 4)

	inputs[0] = NewInput()
	inputs[0].Init(InputProps{SetValue: cfg.Sanitize})
	inputs[0].Prompt = "sanitize (y/n): "
	inputs[0].Validate = ValidateBool

	inputs[1] = NewInput()
	inputs[1].Init(InputProps{SetValue: cfg.NewName})
	inputs[1].Prompt = "new name: "

	inputs[2] = NewInput()
	inputs[2].Init(InputProps{SetValue: cfg.SetPrefix})
	inputs[2].Prompt = "prefix: "

	inputs[3] = NewInput()
	inputs[3].Init(InputProps{SetValue: cfg.SetSuffix})
	inputs[3].Prompt = "suffix: "

	return inputs
}

func ValidateBool(v string) error {
	if v != "y" {
		return fmt.Errorf("invalid")
	}
	return nil
}
