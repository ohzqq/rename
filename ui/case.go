package ui

import "github.com/ohzqq/rename/cfg"

func CaseForm() []*Input {
	inputs := make([]*Input, 3)

	inputs[0] = NewInput()
	inputs[0].Init(InputProps{SetValue: cfg.ToUpper})
	inputs[0].Placeholder = "n"
	inputs[0].Prompt = "uppercase (y/n): "
	inputs[0].Validate = ValidateBool

	inputs[1] = NewInput()
	inputs[1].Init(InputProps{SetValue: cfg.ToLower})
	inputs[1].Placeholder = "n"
	inputs[1].Prompt = "lowercase (y/n): "
	inputs[1].Validate = ValidateBool

	inputs[2] = NewInput()
	inputs[2].Init(InputProps{SetValue: cfg.SetCase})
	inputs[2].Placeholder = "3"
	inputs[2].Prompt = casePrompt
	inputs[2].Validate = ValidateInt
	return inputs
}

const casePrompt = `[0] CamelCase
[1] kebab-case
[2] lowerCamel
[3] snake_case (default)
> `
