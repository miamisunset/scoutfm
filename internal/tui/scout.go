package tui

import (
	"os"
	"strings"
	"time"

	"golang.org/x/term"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/miamisunset/scoutfm/internal/tui/styles"

	"github.com/miamisunset/scoutfm/internal/tui/panes"
)

type scout struct {
	styles         *styles.Styles
	cursor         int
	cwdFileBrowser panes.Cwd
	preview        panes.Preview
	width          int
	height         int
	cwd            string // current working directory
}

func NewScout(cwd string) *scout {
	w, h, _ := term.GetSize(int(os.Stdout.Fd()))

	return &scout{
		styles:         styles.DefaultStyles(),
		cwdFileBrowser: panes.NewCwdPane(w, h-5),
		preview:        panes.NewPreview(w, h-5),
		cwd:            cwd,
		width:          w,
		height:         h - 3,
	}

}

func (s scout) Init() tea.Cmd {
	return tea.Batch(tick())
}

func (s scout) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		case "q":
			return s, tea.Quit
		}

	case tickMsg:
		return s, tick()
	}

	_, cmd := s.cwdFileBrowser.Update(msg)
	cmds = append(cmds, cmd)

	s.preview.Update(msg)

	return s, tea.Batch(cmds...)
}

func (s scout) headerView() string {
	w := lipgloss.Width

	header := s.styles.Header.Render("CWD")
	clock := s.styles.Clock.Render(time.Now().Format("⏰ 3:04:05 pm"))

	cwd := s.styles.CurrentPath.Width(s.width + 2 - w(header) - w(clock)).
		Render(s.cwd)

	headerBar := lipgloss.JoinHorizontal(
		lipgloss.Top,
		header,
		cwd,
		clock,
	)

	return headerBar
}

func (s scout) View() string {
	cwd := s.cwdFileBrowser.View()
	pp := s.preview.View()

	p := lipgloss.JoinHorizontal(
		lipgloss.Top,
		cwd,
		pp,
	)

	b := strings.Builder{}
	b.WriteRune('\n')
	b.WriteString(s.headerView())
	b.WriteRune('\n')
	b.WriteString(p)
	return s.styles.App.Render(b.String())
}

type tickMsg time.Time

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
