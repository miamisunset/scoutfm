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

// Pane messages
type readDirMsg struct{ dir string }

func (m readDirMsg) getDir() string {
	return m.dir
}

type pane struct {
	cursor int
	style  *lipgloss.Style
	path   string
	files  []fs.FileInfo
}

func newPane(style styles.Style) *pane {
	return &pane{
		style: style.GetPane(),
	}
}

func (p *pane) update(msg tea.Msg) (*pane, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case readDirMsg:
		p.path = msg.getDir()
		p.readDir()

	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if p.cursor > 0 {
				p.cursor--
			}
		case "down", "j":
			if p.cursor < len(p.files)-1 {
				p.cursor++
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

			l += fmt.Sprintf("%s %s\n", cursor, f.Name())
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
