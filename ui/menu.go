package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/londek/reactea"
	"github.com/londek/reactea/router"
)

type MenuProps = int

const (
	horizontal = iota + 1
	vertical
)

func initMenu(dir int) router.RouteInitializer {
	return func(router.Params) (reactea.SomeComponent, tea.Cmd) {
		cmpnt := reactea.Componentify[int](MenuRenderer)
		return cmpnt, cmpnt.Init(dir)
	}
}

func MenuRenderer(props MenuProps, w, h int) string {
	switch props {
	case horizontal:
		return lipgloss.NewStyle().
			Background(lipgloss.Color("#afffaf")).
			Foreground(lipgloss.Color("#262626")).
			Render(strings.Join(menuPrompt, "|"))
	case vertical:
		return strings.Join(menuPrompt, "\n")
	default:
		return ""
	}
}

var menuPrompt = []string{
	"[F1]Sanitize",
	"[F2]Case",
	"[F3]Padding",
	"[F4]Replace",
	"[F5]Prefix",
	"[F6]Suffix",
	"[F12]Preview",
	"[esc]Menu",
}
