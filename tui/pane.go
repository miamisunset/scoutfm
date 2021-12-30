package tui

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"github.com/miamisunset/scoutfm/tui/styles"
)

type pane struct {
	style *lipgloss.Style
	files []fs.FileInfo
	dir   string
}

type readDirMsg struct{ dir string }

func newPane(style styles.Style) *pane {
	return &pane{
		style: style.GetPane(),
	}
}

func (p pane) update(msg tea.Msg) (pane, tea.Cmd) {
	switch msg.(type) {
	case readDirMsg:
		p.readDir()
	}

	return p, nil
}

func (p pane) view() string {
	var l string

	if p.files != nil {
		for _, file := range p.files {
			l += fmt.Sprintf("%s", file)
		}
	} else {
		l += "Empty"
	}

	return p.style.Render(l)
}

func (p *pane) readDir() {
	if files, err := ioutil.ReadDir(p.dir); err != nil {
		log.Fatal(err)
	} else {
		p.files = files
	}
}
