package ui

import (
	"os"
	"strings"

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
	var names []string
	for _, name := range c.Props().Names {
		for og, n := range name {
			names = append(names, og+" => "+n)
		}
	}
	c.view.SetContent(strings.Join(names, "\n"))

	return nil
}

func (c *Preview) Update(msg tea.Msg) tea.Cmd {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "ctrl+c" {
			return reactea.Destroy
		}
	}
	c.view, cmd = c.view.Update(msg)
	cmds = append(cmds, cmd)

	return tea.Batch(cmds...)
}

func (c *Preview) Render(w, h int) string {
	return c.view.View()
}

func TermSize() (int, int) {
	w, h, _ := term.GetSize(int(os.Stdin.Fd()))
	return w, h
}
