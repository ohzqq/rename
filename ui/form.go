// Heavily borrowed from: https://github.com/charmbracelet/bubbletea/tree/master/examples/credit-card-form
package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/londek/reactea"
	"github.com/londek/reactea/router"
)

type Form struct {
	reactea.BasicComponent
	reactea.BasicPropfulComponent[reactea.NoProps]

	inputs []*Input

	focused int
}

func NewForm(inputs []*Input) *Form {
	return &Form{
		inputs:  inputs,
		focused: 0,
	}
}

func FormRoute(cmpnt *Form) router.RouteInitializer {
	return func(router.Params) (reactea.SomeComponent, tea.Cmd) {
		return cmpnt, cmpnt.Init(reactea.NoProps{})
	}
}

func (c *Form) Init(reactea.NoProps) tea.Cmd {
	return c.inputs[c.focused].Focus()
}

func (c *Form) Update(msg tea.Msg) tea.Cmd {
	var cmds []tea.Cmd = make([]tea.Cmd, len(c.inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if c.focused == len(c.inputs)-1 {
				for i := range c.inputs {
					c.inputs[i].Save()
				}
				reactea.SetCurrentRoute("preview")
				return nil
			}
			c.nextInput()
		case tea.KeyShiftTab:
			c.prevInput()
		case tea.KeyTab:
			c.nextInput()
		}
		switch key := msg.String(); key {
		case "ctrl+c":
			return reactea.Destroy
		}
		for i := range c.inputs {
			c.inputs[i].Blur()
		}
		c.inputs[c.focused].Focus()
	}

	for i := range c.inputs {
		c.inputs[i], cmds[i] = c.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (c *Form) Render(int, int) string {
	var v []string

	for i := range c.inputs {
		v = append(v, c.inputs[i].View())
	}

	return lipgloss.JoinVertical(lipgloss.Left, v...)
}

func (c *Form) nextInput() {
	c.focused = (c.focused + 1) % len(c.inputs)
}

func (c *Form) prevInput() {
	c.focused--
	// Wrap around
	if c.focused < 0 {
		c.focused = len(c.inputs) - 1
	}
}
