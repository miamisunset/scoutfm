package tui

import (
	"fmt"
	"os"

	"golang.org/x/term"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/miamisunset/scoutfm/tui/styles"
)

type scout struct {
	styles *styles.Style
	width  int
	height int
	header *header
}

func NewScout(cwd string) *scout {
	w, h, err := term.GetSize(int(os.Stdout.Fd()))

	if err != nil {
		fmt.Println("Unable to initialize terminal...")
		os.Exit(1)
	}

	styles := styles.DefaultStyles()

	return &scout{
		styles: styles,
		width:  w,
		height: h,
		header: newHeader(w-styles.GetAppStyle().GetHorizontalMargins(), *styles),
	}
}

func (s scout) Init() tea.Cmd {
	return nil
}

func (s scout) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	// When terminal size changes
	case tea.WindowSizeMsg:
		s.width = msg.Width
		s.height = msg.Height
		s.header.setWidth(msg.Width - s.styles.GetAppStyle().GetHorizontalMargins())

	// Shortcuts
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return s, tea.Quit
		}
	}

	return s, nil
}

func (s scout) View() string {
	return s.styles.GetAppStyle().Render(s.header.view())
}
