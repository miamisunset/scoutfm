package tui

import (
	"fmt"
	"io/fs"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"gitlab.com/Synthwave/scoutfm/internal/tui/styles"

	fz "gitlab.com/Synthwave/scoutfm/internal/fs"
)

type scout struct {
	styles       *styles.Styles
	cursor       int
	files        []fs.FileInfo
	selectedFile map[int]struct{}
	termWidth    int
	termHeight   int
	cwd          string // current working directory
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

		case "up", "k":
			if s.cursor > 0 {
				s.cursor--
			}
		case "down", "j":
			if s.cursor < len(s.files)-1 {
				s.cursor++
			}
		}

	}

	return s, nil
}

func (s scout) headerView() string {
	header := s.styles.Header.Render("PATH")
	cwd := s.styles.App.Render(s.cwd)

	headerBar := lipgloss.JoinHorizontal(
		lipgloss.Top,
		header,
		cwd,
	)

	return headerBar
}

func (s scout) fileBrowser() string {

	var fileList string

	for i, file := range s.files {
		cursor := " "

		filename := file.Name()

		if s.cursor == i {
			cursor = ">"

			sb := strings.Builder{}
			sb.WriteString(cursor)
			sb.WriteRune(' ')
			sb.WriteString(filename)

			selected := s.styles.App.
				Foreground(lipgloss.Color("#14F9D5")).
				Render(sb.String())

			fileList += fmt.Sprintf("%s\n", selected)
		} else {
			if file.IsDir() {
				filename = s.styles.App.
					Foreground(lipgloss.Color("#F25D94")).
					Render(file.Name())
			}

			fileList += fmt.Sprintf("%s %s\n", cursor, filename)
		}
	}

	return s.styles.App.
		BorderStyle(s.styles.FileBrowserBorder).
		BorderForeground(s.styles.BorderColor).
		Width(s.termWidth).
		Height(s.termHeight).
		Render(fileList)
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
		files:      fz.ReadDir(cwd),
	}
}
