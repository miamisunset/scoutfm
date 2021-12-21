package styles

import (
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	App               lipgloss.Style
	FileBrowserBorder lipgloss.Border

	Header lipgloss.Style

	BorderColor lipgloss.Color
}

func DefaultStyles() *Styles {
	s := new(Styles)

	s.App = lipgloss.NewStyle()

	s.FileBrowserBorder = lipgloss.RoundedBorder()

	s.Header = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFDF5")).
		Background(lipgloss.Color("#FF5F87")).
		Padding(0, 1).
		MarginRight(1)

	s.BorderColor = "#874BFD"

	return s
}
