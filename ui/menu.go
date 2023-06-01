package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/londek/reactea"
	"github.com/londek/reactea/router"
	"github.com/ohzqq/rename/cfg"
)

type MenuProps = int

const (
	horizontal = iota + 1
	vertical
)

//go:generate stringer -type MenuEntry
type MenuEntry int

const (
	Num MenuEntry = iota
	Case
	Replace
	Misc
	View
	Menu
)

func (m MenuEntry) Key() string {
	switch m {
	case Num:
		return "f1"
	case Case:
		return "f2"
	case Replace:
		return "f3"
	case Misc:
		return "f4"
	case View:
		return "f12"
	case Menu:
		return "esc"
	}
	return ""
}

func initMenu(dir int) router.RouteInitializer {
	return func(router.Params) (reactea.SomeComponent, tea.Cmd) {
		cmpnt := reactea.Componentify[int](MenuRenderer)
		return cmpnt, cmpnt.Init(dir)
	}
}

func MiscForm() *Form {
	inputs := make([]*Input, 3)

	inputs[0] = NewInput(cfg.SetPrefix)
	inputs[0].Prompt = "prefix: "

	inputs[1] = NewInput(cfg.SetSuffix)
	inputs[1].Prompt = "suffix: "

	inputs[2] = NewInput(cfg.Sanitize)
	inputs[2].Prompt = "sanitize (y/n): "
	inputs[2].Validate = ValidateBool

	return &Form{
		inputs:  inputs,
		focused: 0,
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

var menuKeyStyle = lipgloss.NewStyle().
	Background(lipgloss.Color("#afffaf")).
	Foreground(lipgloss.Color("#262626"))

var menuPrompt = []map[string]string{
	map[string]string{"[F1]": "num"},
	map[string]string{"[F2]": "case"},
	map[string]string{"[F3]": "replace"},
	map[string]string{"[F4]": "misc"},
	map[string]string{"[F12]": "preview"},
	map[string]string{"[esc]": "menu"},
}

func ValidateBool(v string) error {
	if v != "y" {
		return fmt.Errorf("invalid")
	}
	return nil
}
