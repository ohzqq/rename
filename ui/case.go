package ui

import "github.com/ohzqq/rename/cfg"

func CaseForm() *Input {
	input := NewInput()
	input.Init(InputProps{SetValue: cfg.SetCase})
	input.Placeholder = "3"
	input.Prompt = casePrompt
	input.Validate = ValidateInt
	return input
}

const casePrompt = `[0] CamelCase
[1] kebab-case
[2] lowerCamel
[3] snake_case (default)
[4] lower
[5] UPPER
> `
