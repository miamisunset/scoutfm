package styles

import (
	"github.com/charmbracelet/lipgloss"
)

type Styles struct {
	App               lipgloss.Style
	FileBrowserBorder lipgloss.Border
	BorderColor       lipgloss.Color
}

func DefaultStyles() *Styles {
	s := new(Styles)

	s.App = lipgloss.NewStyle()

	s.FileBrowserBorder = lipgloss.RoundedBorder()
	s.BorderColor = "#874BFD"

	return s
}
