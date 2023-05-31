package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/londek/reactea"
	"github.com/londek/reactea/router"
	"github.com/ohzqq/rename/batch"
)

type App struct {
	reactea.BasicComponent
	reactea.BasicPropfulComponent[reactea.NoProps]

	router reactea.Component[router.Props]

	names *batch.Names
	route string
}

type FormProps struct {
	SetNames func([]map[string]string)
}

func New(names *batch.Names) *App {
	return &App{
		route:  "preview",
		router: router.New(),
		names:  names,
	}
}

func (c *App) Route(r string) *App {
	c.route = r
	return c
}

func (c *App) Init(reactea.NoProps) tea.Cmd {
	routes := map[string]router.RouteInitializer{
		"preview": PreviewRoute(c.names),
		"padding": FormRoute(NewPaddingForm()),
		"replace": FormRoute(FindReplaceForm()),
	}
	routes["default"] = routes[c.route]
	return c.router.Init(routes)
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

func (c *App) SetNames(names *batch.Names) {
	c.names = names
}
