package styles

import "github.com/charmbracelet/lipgloss"

type Style struct {
	app    lipgloss.Style
	header lipgloss.Style
	title  lipgloss.Style
	cwd    lipgloss.Style
	clock  lipgloss.Style
	whoAmI lipgloss.Style
	pane   lipgloss.Style
	dir    lipgloss.Style
}

func DefaultStyles() *Style {
	s := new(Style)

	s.app = lipgloss.NewStyle().
		Margin(1, 2, 1, 2)

	s.header = lipgloss.NewStyle().
		Background(lipgloss.Color("#333"))

	s.title = lipgloss.NewStyle().
		Inherit(s.header).
		Background(lipgloss.Color("#FF5F87")).
		Foreground(lipgloss.Color("#FFFDF5")).
		Padding(0, 1)

	s.cwd = lipgloss.NewStyle().
		Inherit(s.header).
		Foreground(lipgloss.Color("#1ed3ec"))

	s.clock = lipgloss.NewStyle().
		Inherit(s.header).
		Background(lipgloss.Color("#6124DF")).
		Foreground(lipgloss.Color("#fff")).
		Padding(0, 1).
		Align(lipgloss.Right)

	s.whoAmI = lipgloss.NewStyle().
		Inherit(s.header).
		Foreground(lipgloss.Color("#ffe261")).
		Padding(0, 1)

	s.pane = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#666"))

	s.dir = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF5F87"))

	return s
}

func (s Style) GetAppStyle() *lipgloss.Style {
	return &s.app
}

func (s Style) GetHeader() *lipgloss.Style {
	return &s.header
}

func (s Style) GetCwd() *lipgloss.Style {
	return &s.cwd
}

func (s Style) GetTitle() *lipgloss.Style {
	return &s.title
}

func (s Style) GetClock() *lipgloss.Style {
	return &s.clock
}

func (s Style) GetWhoAmI() *lipgloss.Style {
	return &s.whoAmI
}

func (s Style) GetPane() *lipgloss.Style {
	return &s.pane
}

func (s Style) GetDir() *lipgloss.Style {
	return &s.dir
}
