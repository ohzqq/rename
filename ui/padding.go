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

func NewPaddingForm() *Form {
	inputs := make([]*Input, 3)
	inputs[zeroes] = NewInput(cfg.SetZeroes)
	inputs[zeroes].Placeholder = "0"
	if p := cfg.Padding().Zeroes; p > 0 {
		inputs[zeroes].SetValue(strconv.Itoa(p))
	}
	inputs[zeroes].Width = 5
	inputs[zeroes].Prompt = "zeroes: "
	inputs[zeroes].Validate = ValidateInt

	inputs[start] = NewInput(cfg.SetStart)
	inputs[start].Placeholder = strconv.Itoa(cfg.Padding().Start)
	inputs[start].Width = 5
	inputs[start].Prompt = "start: "
	inputs[start].Validate = ValidateInt

	inputs[position] = NewInput(cfg.SetPosition)
	inputs[position].Placeholder = strconv.Itoa(cfg.Padding().Position)
	inputs[position].Width = 5
	inputs[position].Prompt = padMenu
	inputs[position].Validate = ValidateInt

	return &Form{
		inputs:  inputs,
		focused: 0,
	}
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
