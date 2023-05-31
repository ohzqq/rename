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
	inputs[zeroes].SetValue("0")
	if p := cfg.Padding().Zeroes; p > 0 {
		inputs[zeroes].SetValue(strconv.Itoa(p))
	}
	inputs[zeroes].Width = 5
	inputs[zeroes].Prompt = "zeroes: "

	inputs[start] = NewInput(cfg.SetStart)
	inputs[start].SetValue(strconv.Itoa(cfg.Padding().Start))
	inputs[start].Width = 5
	inputs[start].Prompt = "start: "

	inputs[position] = NewInput(cfg.SetPosition)
	inputs[position].SetValue(strconv.Itoa(cfg.Padding().Position))
	inputs[position].Width = 5
	inputs[position].Prompt = padMenu

	return &Form{
		inputs:  inputs,
		focused: 0,
	}
}

const padMenu = `pos:   
  [0] Start 
  [1] Before Name
  [2] After Name
  [3] End
> `
