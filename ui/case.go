package ui

import "github.com/ohzqq/rename/cfg"

func CaseForm() *Form {
	inputs := make([]*Input, 1)
	inputs[0] = NewInput(cfg.SetCase)
	inputs[0].Placeholder = "3"
	inputs[0].Prompt = casePrompt
	inputs[0].Validate = ValidateInt
	return &Form{
		inputs:  inputs,
		focused: 0,
	}
}

const casePrompt = `[0] CamelCase
[1] kebab-case
[2] lowerCamel
[3] snake_case (default)
[4] lower
[5] UPPER
> `
