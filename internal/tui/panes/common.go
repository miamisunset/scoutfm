package panes

import "github.com/miamisunset/scoutfm/internal/tui/styles"

type commonPane struct {
	styles *styles.Styles

	Width  int
	Height int
}

func newCommonPane(width int) commonPane {
	return commonPane{
		styles: styles.DefaultStyles(),
		Width:  width,
		Height: 0,
	}
}
