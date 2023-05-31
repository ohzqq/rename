package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/londek/reactea"
)

type Input struct {
	textinput.Model
	Set func(any)
}

func NewInput(set func(any)) *Input {
	return &Input{
		Model: textinput.New(),
		Set:   set,
	}
}

func (c *Input) Save() {
	if val := c.Model.Value(); val != "" {
		c.Set(val)
	}
}

func (c *Input) Update(msg tea.Msg) (*Input, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch key := msg.String(); key {
		case "ctrl+c":
			return c, reactea.Destroy
		}
	}

	c.Model, cmd = c.Model.Update(msg)

	return c, cmd
}
