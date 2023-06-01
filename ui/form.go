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
	reactea.BasicPropfulComponent[FormProps]

	focused int
}

type FormProps struct {
	Inputs []*Input
}

func NewForm() *Form {
	return &Form{
		focused: 0,
	}
}

func FormRoute(inputs ...*Input) router.RouteInitializer {
	return func(router.Params) (reactea.SomeComponent, tea.Cmd) {
		cmpnt := NewForm()
		return cmpnt, cmpnt.Init(FormProps{Inputs: inputs})
	}
}

func (c *Form) Init(props FormProps) tea.Cmd {
	c.UpdateProps(props)
	return c.Props().Inputs[c.focused].Focus()
}

func (c *Form) Update(msg tea.Msg) tea.Cmd {
	var cmds []tea.Cmd = make([]tea.Cmd, len(c.Props().Inputs))

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if c.focused == len(c.Props().Inputs)-1 {
				for i := range c.Props().Inputs {
					c.Props().Inputs[i].Save()
				}
				reactea.SetCurrentRoute(View.String())
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
		for i := range c.Props().Inputs {
			c.Props().Inputs[i].Blur()
		}
		c.Props().Inputs[c.focused].Focus()
	}

	for i := range c.Props().Inputs {
		c.Props().Inputs[i], cmds[i] = c.Props().Inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func (c *Form) Render(int, int) string {
	var v []string

	for i := range c.Props().Inputs {
		v = append(v, c.Props().Inputs[i].View())
	}

	return lipgloss.JoinVertical(lipgloss.Left, v...)
}

func (c *Form) nextInput() {
	c.focused = (c.focused + 1) % len(c.Props().Inputs)
}

func (c *Form) prevInput() {
	c.focused--
	// Wrap around
	if c.focused < 0 {
		c.focused = len(c.Props().Inputs) - 1
	}
}
