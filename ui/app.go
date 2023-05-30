package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/londek/reactea"
	"github.com/londek/reactea/router"
)

type App struct {
	reactea.BasicComponent
	reactea.BasicPropfulComponent[reactea.NoProps]

	router reactea.Component[router.Props]

	names []map[string]string
}

func New(names []map[string]string) *App {
	return &App{
		router: router.New(),
		names:  names,
	}
}

func (c *App) Init(reactea.NoProps) tea.Cmd {
	return c.router.Init(map[string]router.RouteInitializer{
		"default": func(router.Params) (reactea.SomeComponent, tea.Cmd) {
			cmpnt := NewPreview()
			props := PreviewProps{
				Names: c.names,
			}
			return cmpnt, cmpnt.Init(props)
		},
	})
}

func (c *App) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// ctrl+c support
		if msg.String() == "ctrl+c" {
			return reactea.Destroy
		}
	}

	return c.router.Update(msg)
}

func (c *App) Render(width, height int) string {
	return c.router.Render(width, height)
}
