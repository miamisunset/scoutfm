package tui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/miamisunset/scoutfm/tui/styles"
)

type footer struct {
	width  int
	height int

	style *lipgloss.Style
}

func newFooter(width int, style styles.Style) *footer {
	return &footer{
		width: width,
		style: style.GetFooter(),
	}
}

func (f footer) view() string {
	return f.style.Width(f.width).Render("status")
}

func (f *footer) setWidth(w int) {
	f.width = w
}
