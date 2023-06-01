package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/londek/reactea"
)

type Input struct {
	reactea.BasicComponent
	reactea.BasicPropfulComponent[InputProps]

	textinput.Model
	SetValue func(any)
}

type InputProps struct {
	SetValue func(any)
}

func NewInput() *Input {
	return &Input{
		Model: textinput.New(),
	}
}

func (c *Input) Init(props InputProps) tea.Cmd {
	c.UpdateProps(props)
	return nil
}

func (c *Input) Save() {
	if val := c.Model.Value(); val != "" {
		c.Props().SetValue(val)
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

func (c *Input) Render(int, int) string {
	return c.Model.View()
}
