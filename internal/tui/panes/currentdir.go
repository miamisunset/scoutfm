package panes

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	fz "github.com/miamisunset/scoutfm/internal/fs"
)

const (
	cwdPaneMaxWidth = 0.30 // percentage of screen
)

type SelectedFileMsg struct {
	Name string
}

type CwdPane struct {
	commonPane

	Cursor int

	Files        []fs.FileInfo
	SelectedFile int
}

func NewCwdPane(width, height int) CwdPane {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal()
	}

	w := int(float64(width) * cwdPaneMaxWidth)

	files := fz.ReadDir(cwd)
	fz.SortByName(files, true)
	fz.SortDirFirst(files)

	return CwdPane{
		commonPane:   newCommonPane(w),
		Cursor:       0,
		Files:        files,
		SelectedFile: 0,
	}
}

func (p *CwdPane) Update(msg tea.Msg) (*CwdPane, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		case "up", "k":
			if p.Cursor > 0 {
				p.Cursor--
			}

		case "down", "j":
			if p.Cursor < len(p.Files)-1 {
				p.Cursor++
			}
		}
	}

	return p, nil
}

func (p *CwdPane) View() string {
	b := strings.Builder{}
	b.WriteString(p.fileBrowser())
	return p.styles.App.Render(b.String())
}

func (p CwdPane) fileBrowser() string {

	var fileList string

	for i, file := range p.Files {
		cursor := " "

		filename := file.Name()

		if p.Cursor == i {
			cursor = ">"

			sb := strings.Builder{}
			sb.WriteString(cursor)
			sb.WriteRune(' ')
			sb.WriteString(filename)

			selected := p.styles.App.
				Foreground(lipgloss.Color("#14F9D5")).
				Render(sb.String())

			fileList += fmt.Sprintf("%s\n", selected)
		} else {
			if file.IsDir() {
				filename = p.styles.App.
					Foreground(lipgloss.Color("#F25D94")).
					Render(file.Name())
			}

			fileList += fmt.Sprintf("%s %s\n", cursor, filename)
		}
	}

	return p.styles.App.
		BorderStyle(p.styles.FileBrowserBorder).
		BorderForeground(p.styles.BorderColor).
		Width(p.Width).
		Height(p.Height).
		Render(fileList)
}
