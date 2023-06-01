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
		View.String():    PreviewRoute(c.names),
		Num.String():     FormRoute(NewPaddingForm()...),
		Replace.String(): FormRoute(FindReplaceForm()...),
		Case.String():    FormRoute(CaseForm()...),
		Name.String():    FormRoute(NameForm()...),
		Menu.String():    initMenu(vertical),
	}
	routes["default"] = routes[Name.String()]
	return c.router.Init(routes)
}

func (c *App) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch key := msg.String(); key {
		case "?":
			reactea.SetCurrentRoute(Menu.String())
		case "ctrl+c":
			return reactea.Destroy
		default:
			for _, ent := range menuEntries {
				if key == ent.Key() {
					reactea.SetCurrentRoute(ent.String())
				}
			}
		}
	}

	return c.router.Update(msg)
}

func (c *App) Render(w, h int) string {
	var views []string
	if r := reactea.CurrentRoute(); r != Menu.String() {
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
