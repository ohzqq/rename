package ui

import (
	"strconv"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/londek/reactea"
	"github.com/spf13/viper"
)

type Padding struct {
	reactea.BasicComponent
	reactea.BasicPropfulComponent[reactea.NoProps]

	inputs []textinput.Model

	focused int
}

const (
	zeroes = iota
	start
	position
)

func NewPaddingForm() *Padding {
	inputs := make([]textinput.Model, 3)
	inputs[zeroes] = textinput.New()
	inputs[zeroes].SetValue("0")
	if p := viper.GetInt("pad"); p > 0 {
		inputs[zeroes].SetValue(strconv.Itoa(p))
	}
	inputs[zeroes].Width = 5
	inputs[zeroes].Prompt = ""

	inputs[start] = textinput.New()
	inputs[start].SetValue(viper.GetString("min"))
	inputs[start].Width = 5
	inputs[start].Prompt = ""

	inputs[position] = textinput.New()
	inputs[position].SetValue("1")
	inputs[position].Width = 5
	inputs[position].Prompt = ""
	return &Padding{
		inputs:  inputs,
		focused: 0,
	}
}

func (c *Padding) Init(reactea.NoProps) tea.Cmd {
	return c.inputs[c.focused].Focus()
}

func (c *Padding) Update(msg tea.Msg) tea.Cmd {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyEnter {
			if c.focused == len(c.inputs)-1 {
				reactea.SetCurrentRoute("default")
				return nil
			}
			c.nextInput()
		}
		switch key := msg.String(); key {
		case "ctrl+c":
			return reactea.Destroy
		}
	}

	for i := range c.inputs {
		c.inputs[i], cmds[i] = c.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (c *Padding) Render(int, int) string {
	var v []string

	v = append(v, c.inputs[zeroes].View())
	v = append(v, c.inputs[start].View())
	v = append(v, c.inputs[position].View())

	return lipgloss.JoinVertical(lipgloss.Left, v...)
}

func (c *Padding) nextInput() {
	c.focused = (c.focused + 1) % len(c.inputs)
}

func (c *Padding) prevInput() {
	c.focused--
	// Wrap around
	if c.focused < 0 {
		c.focused = len(c.inputs) - 1
	}
}
