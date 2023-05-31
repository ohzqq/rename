package ui

import (
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/londek/reactea"
	"github.com/londek/reactea/router"
	"github.com/ohzqq/rename/batch"
	"golang.org/x/term"
)

type Preview struct {
	reactea.BasicComponent
	reactea.BasicPropfulComponent[PreviewProps]

	view    viewport.Model
	width   int
	height  int
	oldpath lipgloss.Style
	newpath lipgloss.Style
	names   []map[string]string
}

type PreviewProps struct {
	Names *batch.Names
}

var (
	arrow   = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF5F5F"))
	confirm = lipgloss.NewStyle().Background(lipgloss.Color("#FF5F5F")).Foreground(lipgloss.Color("#262626"))
)

func NewPreview() *Preview {
	w, h := TermSize()
	d := lipgloss.NewStyle().Width(w)
	return &Preview{
		width:   w,
		height:  h,
		oldpath: d,
		newpath: d.Copy().Foreground(lipgloss.Color("#eeeeee")),
		view:    viewport.New(w, h-1),
	}
}

func PreviewRoute(names *batch.Names) router.RouteInitializer {
	return func(router.Params) (reactea.SomeComponent, tea.Cmd) {
		cmpnt := NewPreview()
		return cmpnt, cmpnt.Init(PreviewProps{Names: names})
	}
}

func (c *Preview) Init(props PreviewProps) tea.Cmd {
	c.UpdateProps(props)
	c.names = c.Props().Names.Transform()
	return nil
}

func (c *Preview) Update(msg tea.Msg) tea.Cmd {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch key := msg.String(); key {
		case "y", "enter":
			for _, name := range c.names {
				for og, n := range name {
					os.Rename(og, n)
				}
			}
			fallthrough
		case "n":
			fallthrough
		case "ctrl+c", "q":
			return reactea.Destroy
		}
	}
	c.view, cmd = c.view.Update(msg)
	cmds = append(cmds, cmd)

	return tea.Batch(cmds...)
}

func (c *Preview) Render(w, h int) string {
	var names []string
	for _, name := range c.names {
		var s strings.Builder
		for og, n := range name {
			if og != n {
				s.WriteString(c.oldpath.Render(og))
				s.WriteString(arrow.Render("\n=> "))
				s.WriteString(c.newpath.Render(n))
				names = append(names, s.String())
			}
		}
	}
	c.view.SetContent(strings.Join(names, "\n\n"))

	return lipgloss.JoinVertical(
		lipgloss.Left,
		c.view.View(),
		confirm.Render("rename? (y|enter/no)"),
	)
}

func TermSize() (int, int) {
	w, h, _ := term.GetSize(int(os.Stdin.Fd()))
	return w, h
}
