package ui

import (
	"strconv"

	"github.com/ohzqq/rename/cfg"
)

const (
	zeroes = iota
	start
	position
)

func NewPaddingForm() []*Input {
	inputs := make([]*Input, 3)
	inputs[zeroes] = NewInput()
	inputs[zeroes].Init(InputProps{SetValue: cfg.SetZeroes})
	inputs[zeroes].Placeholder = "0"
	if p := cfg.Zeroes(); p > 0 {
		inputs[zeroes].Placeholder = strconv.Itoa(p)
	}
	inputs[zeroes].Prompt = "zeroes: "
	inputs[zeroes].Validate = ValidateInt

	inputs[start] = NewInput()
	inputs[start].Init(InputProps{SetValue: cfg.SetStart})
	inputs[start].Placeholder = strconv.Itoa(cfg.Start())
	inputs[start].Prompt = "start: "
	inputs[start].Validate = ValidateInt

	inputs[position] = NewInput()
	inputs[position].Init(InputProps{SetValue: cfg.SetPosition})
	inputs[position].Placeholder = "3"
	inputs[position].Prompt = padMenu
	inputs[position].Validate = ValidateInt

	return inputs
}

func ValidateInt(a string) error {
	_, err := strconv.Atoi(a)
	if err != nil {
		return err
	}
	return nil
}

const padMenu = `pos:   
  [0] Start 
  [1] Before Name
  [2] After Name
  [3] End
> `
