package tui

import (
	"gitlab.com/Synthwave/scoutfm/internal/fs"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"gitlab.com/Synthwave/scoutfm/internal/tui/styles"
)

type scout struct {
	styles     *styles.Styles
	termWidth  int
	termHeight int
	cwd        string // current working directory
}

func (s scout) Init() tea.Cmd {
	return nil
}

func (s scout) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return s, tea.Quit
		}
	}

	return s, nil
}

func (s scout) headerView() string {
	return s.cwd
}

func (s scout) fileBrowser() string {

	fb := strings.Builder{}

	files := fs.ReadDir(s.cwd)
	l := len(files) - 1

	for i, f := range files {
		fb.WriteString(f.Name())
		if i < l {
			fb.WriteRune('\n')
		}
	}

	return s.styles.App.
		BorderStyle(s.styles.FileBrowserBorder).
		BorderForeground(s.styles.BorderColor).
		Width(s.termWidth).
		Height(s.termHeight).
		Render(fb.String())
}

func (s scout) View() string {
	b := strings.Builder{}
	b.WriteString(s.headerView())
	b.WriteRune('\n')
	b.WriteString(s.fileBrowser())
	return s.styles.App.Render(b.String())
}

func NewScout(cwd string, termWidth int, termHeight int) *scout {
	return &scout{
		styles:     styles.DefaultStyles(),
		cwd:        cwd,
		termWidth:  termWidth - 2,
		termHeight: termHeight - 3,
	}
}
