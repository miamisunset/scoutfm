package styles

import (
	"github.com/charmbracelet/lipgloss"
	"gitlab.com/Synthwave/scoutfm/internal/tui/colors"
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
	s.BorderColor = colors.Green

	return s
}
