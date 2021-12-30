package tui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/miamisunset/scoutfm/tui/styles"
	"io/fs"
	"io/ioutil"
	"log"
)

type pane struct {
	cursor   int
	path     string
	style    *lipgloss.Style
	dirStyle *lipgloss.Style
	files    []fs.FileInfo
}

func newPane(style styles.Style) *pane {
	return &pane{
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
				filename = p.dirStyle.Render(f.Name())
			} else {
				filename = f.Name()
			}

			l += fmt.Sprintf("%s %s\n", cursor, filename)
		}
	}

	return p.style.Render(l)
}

func (p *pane) readDir() {
	if files, err := ioutil.ReadDir(p.path); err != nil {
		log.Fatalln(err)
	} else {
		p.files = files
	}
}

func (p pane) sendSelectedFile() tea.Msg {
	return selectedFileMsg{
		name: p.files[p.cursor].Name(),
	}
}
