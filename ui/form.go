package ui

import (
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/londek/reactea"
	"github.com/londek/reactea/router"
	"github.com/ohzqq/rename/cfg"
)

type Form struct {
	reactea.BasicComponent
	reactea.BasicPropfulComponent[reactea.NoProps]

	inputs []*Input

	focused int
}

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
	inputs[position].Prompt = "pos: "
	return &Form{
		inputs:  inputs,
		focused: 0,
	}
}

func FormRoute() router.RouteInitializer {
	return func(router.Params) (reactea.SomeComponent, tea.Cmd) {
		cmpnt := NewPaddingForm()
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
		case tea.KeyShiftTab, tea.KeyCtrlP:
			c.prevInput()
		case tea.KeyTab, tea.KeyCtrlN:
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

	v = append(v, c.inputs[zeroes].View())
	v = append(v, c.inputs[start].View())
	v = append(v, c.inputs[position].View())
	v = append(v, padMenu)

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

const padMenu = `  [0] Start 
  [1] Before Name
  [2] After Name
  [3] End`
