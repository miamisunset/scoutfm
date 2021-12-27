package styles

import (
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	App                  lipgloss.Style
	FileBrowserBorder    lipgloss.Border
	PreviewBrowserBorder lipgloss.Border

	Header      lipgloss.Style
	Clock       lipgloss.Style
	CurrentPath lipgloss.Style

	CwdFileBrowser lipgloss.Style

	BorderColor lipgloss.Color
}

func DefaultStyles() *Styles {
	s := new(Styles)

	s.App = lipgloss.NewStyle()

	s.FileBrowserBorder = lipgloss.RoundedBorder()

	s.PreviewBrowserBorder = lipgloss.RoundedBorder()

	s.CurrentPath = lipgloss.NewStyle().
		Background(lipgloss.Color("#353533")).
		Padding(0, 2)

	s.Header = lipgloss.NewStyle().
		Inherit(s.CurrentPath).
		Foreground(lipgloss.Color("#FFFDF5")).
		Background(lipgloss.Color("#FF5F87")).
		Padding(0, 1).
		MarginLeft(2)

	s.Clock = lipgloss.NewStyle().
		Inherit(s.CurrentPath).
		Background(lipgloss.Color("#6124DF")).
		Padding(0, 2).
		MarginRight(3)

	s.CwdFileBrowser = lipgloss.NewStyle().
		Padding(0, 1)

	s.BorderColor = "#353533"

	return s
}
