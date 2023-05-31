package ui

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
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
		return strings.Join(menuPrompt, " | ")
	case vertical:
		return strings.Join(menuPrompt, "\n")
	default:
		return ""
	}
}

var menuPrompt = []string{
	"[F1] Case",
	"[F2] Padding",
	"[F3] Replace",
	"[F4] Prefix",
	"[F5] Suffix",
	"[esc] Menu",
}
