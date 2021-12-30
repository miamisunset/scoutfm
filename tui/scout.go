package tui

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/term"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/miamisunset/scoutfm/tui/styles"
)

type scout struct {
	styles  *styles.Style
	width   int
	height  int
	header  *header
	cwdPane *pane
}

func NewScout(cwd string) *scout {
	w, h, err := term.GetSize(int(os.Stdout.Fd()))

	if err != nil {
		fmt.Println("Unable to initialize terminal...")
		os.Exit(1)
	}

	s := styles.DefaultStyles()

	widgetWidth := w - s.GetAppStyle().GetHorizontalMargins()

	return &scout{
		styles:  s,
		width:   w,
		height:  h,
		header:  newHeader(widgetWidth, *s),
		cwdPane: newPane(widgetWidth, *s),
	}
}

func (s scout) Init() tea.Cmd {
	return tea.Batch(tick(), s.setupCmd)
}

func (s scout) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	// When terminal size changes
	case tea.WindowSizeMsg:
		s.width = msg.Width
		s.height = msg.Height

		w := msg.Width - s.styles.GetAppStyle().GetHorizontalMargins()

		s.header.setWidth(w)
		s.cwdPane.setWidth(w)

	case tickMsg:
		return s, tick()

	// Shortcuts
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return s, tea.Quit
		}
	}

	_, cmd := s.cwdPane.update(msg)
	cmds = append(cmds, cmd)
	_, cmd = s.header.update(msg)
	cmds = append(cmds, cmd)

	return s, tea.Batch(cmds...)
}

func (s scout) View() string {
	layout := lipgloss.JoinVertical(
		lipgloss.Top,
		s.header.view(),
		s.cwdPane.view(),
	)

	return s.styles.GetAppStyle().Render(layout)
}

func (s *scout) setupCmd() tea.Msg {
	return readDirMsg{dir: "."}
}

type tickMsg time.Time

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
