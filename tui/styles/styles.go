package styles

import "github.com/charmbracelet/lipgloss"

type Style struct {
	app    lipgloss.Style
	header lipgloss.Style
	title  lipgloss.Style
	cwd    lipgloss.Style
	clock  lipgloss.Style
}

func DefaultStyles() *Style {
	s := new(Style)

	s.app = lipgloss.NewStyle().
		Margin(1, 1, 1, 1)

	s.header = lipgloss.NewStyle().
		Background(lipgloss.Color("#333"))

	s.title = lipgloss.NewStyle().
		Inherit(s.header).
		Background(lipgloss.Color("#FF5F87")).
		Foreground(lipgloss.Color("#FFFDF5")).
		Padding(0, 1)

	s.cwd = lipgloss.NewStyle().
		Inherit(s.header).
		Padding(0, 1)

	s.clock = lipgloss.NewStyle().
		Inherit(s.header).
		Background(lipgloss.Color("#6124DF")).
		Padding(0, 1).
		Align(lipgloss.Right)

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
