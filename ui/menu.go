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

//go:generate stringer -type MenuEntry
type MenuEntry int

const (
	Name MenuEntry = iota
	Num
	Case
	Replace
	View
	Menu
)

var menuEntries = []MenuEntry{
	Name,
	Num,
	Case,
	Replace,
	View,
	Menu,
}

func (m MenuEntry) Key() string {
	switch m {
	case Name:
		return "f1"
	case Num:
		return "f2"
	case Case:
		return "f3"
	case Replace:
		return "f4"
	case View:
		return "f12"
	case Menu:
		return "esc"
	}
	return "?"
}

func (m MenuEntry) Render() string {
	cur := reactea.CurrentRoute()
	key := "[" + m.Key() + "]"
	switch {
	case cur == m.String():
		fallthrough
	case cur == "" && m.String() == initialRoute:
		fallthrough
	case cur == "default" && m.String() == initialRoute:
		return menuKeyActiveStyle.Render(key + m.String())
	default:
		return menuKeyInactiveStyle.Render(key) + m.String()
	}
}

func initMenu(dir int) router.RouteInitializer {
	return func(router.Params) (reactea.SomeComponent, tea.Cmd) {
		cmpnt := reactea.Componentify[int](MenuRenderer)
		return cmpnt, cmpnt.Init(dir)
	}
}

func MenuRenderer(props MenuProps, w, h int) string {
	var menu []string
	for _, ent := range menuEntries {
		menu = append(menu, ent.Render())
	}
	switch props {
	case horizontal:
		return strings.Join(menu, "")
	case vertical:
		return strings.Join(menu, "\n")
	default:
		return ""
	}
}

var menuKeyActiveStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#afffaf")).
	Foreground(lipgloss.Color("#262626"))

var menuKeyInactiveStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#afffaf"))
