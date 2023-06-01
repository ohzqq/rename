package ui

import (
	"fmt"

	"github.com/ohzqq/rename/cfg"
)

func DefaultForm() []*Input {
	inputs := make([]*Input, 3)

	inputs[0] = NewInput()
	inputs[0].Init(InputProps{SetValue: cfg.SetPrefix})
	inputs[0].Prompt = "prefix: "

	inputs[1] = NewInput()
	inputs[1].Init(InputProps{SetValue: cfg.Sanitize})
	inputs[1].Prompt = "sanitize (y/n): "
	inputs[1].Validate = ValidateBool

	return inputs
}

func PrefixSuffixForm() []*Input {
	inputs := make([]*Input, 3)

	inputs[0] = NewInput()
	inputs[0].Init(InputProps{SetValue: cfg.SetPrefix})
	inputs[0].Prompt = "prefix: "

	inputs[1] = NewInput()
	inputs[1].Init(InputProps{SetValue: cfg.SetSuffix})
	inputs[1].Prompt = "suffix: "

	return inputs
}

func ValidateBool(v string) error {
	if v != "y" {
		return fmt.Errorf("invalid")
	}
	return nil
}
