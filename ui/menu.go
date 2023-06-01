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

var menuPrompt = []string{
	"[F1]Misc",
	"[F2]Case",
	"[F3]Num",
	"[F4]Replace",
	"[F12]Preview",
	"[esc]Menu",
}

func ValidateBool(v string) error {
	if v != "y" {
		return fmt.Errorf("invalide")
	}
	return nil
}
