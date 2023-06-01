package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
		"View":    PreviewRoute(c.names),
		"Num":     FormRoute(NewPaddingForm()),
		"Replace": FormRoute(FindReplaceForm()),
		"Case":    FormRoute(CaseForm()),
		"Misc":    FormRoute(MiscForm()),
		"Menu":    initMenu(vertical),
	}
	routes["default"] = routes["Num"]
	return c.router.Init(routes)
}

func (c *App) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		// ctrl+c support
		switch msg.String() {
		case "f1":
			reactea.SetCurrentRoute("Num")
		case "f2":
			reactea.SetCurrentRoute("Case")
		case "f3":
			reactea.SetCurrentRoute("Replace")
		case "f4":
			reactea.SetCurrentRoute("Misc")
		case "f12":
			reactea.SetCurrentRoute("View")
		case "esc", "?":
			reactea.SetCurrentRoute("Menu")
		case "ctrl+c":
			return reactea.Destroy
		}
	}

	return c.router.Update(msg)
}

func (c *App) Render(w, h int) string {
	var views []string
	if r := reactea.CurrentRoute(); r != "Menu" {
		views = append(views, MenuRenderer(horizontal, w, h))
	}
	views = append(views, c.router.Render(w, h-1))
	return lipgloss.JoinVertical(lipgloss.Left, views...)
}

func (c *App) SetNames(names *batch.Names) {
	c.names = names
}

type UpdateRouteMsg struct {
	Route string
}

func UpdateRoute(r string) tea.Cmd {
	return func() tea.Msg {
		return UpdateRouteMsg{Route: r}
	}
}
