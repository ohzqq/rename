package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Input struct {
	textinput.Model
	Set func(any)
}

const (
	zeroes = iota
	start
	position
)

func NewInput(set func(any)) *Input {
	return &Input{
		Model: textinput.New(),
		Set:   set,
	}
}

func (c *Input) Save() {
	c.Set(c.Model.Value())
}

func (c *Input) Update(msg tea.Msg) (*Input, tea.Cmd) {
	var cmd tea.Cmd

	c.Model, cmd = c.Model.Update(msg)

	return c, cmd
}
