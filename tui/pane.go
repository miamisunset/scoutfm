package tui

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"sort"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/miamisunset/scoutfm/tui/styles"
)

type pane struct {
	cursor   int
	width    int
	path     string
	style    *lipgloss.Style
	dirStyle *lipgloss.Style
	files    []fs.FileInfo
}

func newPane(width int, style styles.Style) *pane {
	return &pane{
		width:    width - 2,
		style:    style.GetPane(),
		dirStyle: style.GetDir(),
	}
}

func (p *pane) update(msg tea.Msg) (*pane, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case readDirMsg:
		p.path = msg.getDir()
		p.readDir()
		cmd = p.sendSelectedFile

	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if p.cursor > 0 {
				p.cursor--
				cmd = p.sendSelectedFile
			}

		case "down", "j":
			if p.cursor < len(p.files)-1 {
				p.cursor++
				cmd = p.sendSelectedFile
			}
		}
	}

	return p, cmd
}

func (p pane) view() string {
	var l string

	if p.path == "" {
		l = "Empty"
	} else {
		var cursor string

		for i, f := range p.files {
			if p.cursor == i {
				cursor = ">"
			} else {
				cursor = " "
			}

			var filename string

			if f.IsDir() {
				filename = p.dirStyle.Render(fmt.Sprintf("🗀 %s", f.Name()))
			} else {
				filename = fmt.Sprintf("🗋 %s", f.Name())
			}

			l += fmt.Sprintf("%s %s\n", cursor, filename)
		}
	}

	return p.style.Width(p.width).Render(l)
}

func (p *pane) readDir() {
	if files, err := ioutil.ReadDir(p.path); err != nil {
		log.Fatalln(err)
	} else {
		sort.Slice(files, func(i, _ int) bool {
			return files[i].IsDir()
		})

		p.files = files
	}
}

func (p pane) sendSelectedFile() tea.Msg {
	return selectedFileMsg{
		file: p.files[p.cursor],
	}
}

func (p *pane) setWidth(w int) {
	p.width = w - p.style.GetBorderRightSize() - p.style.GetBorderLeftSize()
}
