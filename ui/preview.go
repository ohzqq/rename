package ui

import (
	"os"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/londek/reactea"
	"golang.org/x/term"
)

type Preview struct {
	reactea.BasicComponent
	reactea.BasicPropfulComponent[PreviewProps]

	view viewport.Model
}

type PreviewProps struct {
	Names []map[string]string
}

func NewPreview() *Preview {
	return &Preview{
		view: viewport.New(TermSize()),
	}
}

func (c *Preview) Init(props PreviewProps) tea.Cmd {
	c.UpdateProps(props)
	return nil
}

func (c *Preview) Update(msg tea.Msg) tea.Cmd {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return reactea.Destroy
		}
	}
}

func TermSize() (int, int) {
	w, h, _ := term.GetSize(int(os.Stdin.Fd()))
	return w, h
}
