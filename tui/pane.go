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

type refreshMsg struct{}

func sendRefreshMsg() tea.Msg {
	return refreshMsg{}
}

type readDirMsg struct{ dir string }

func (m readDirMsg) getDir() string {
	return m.dir
}

func newPane(style styles.Style) *pane {
	return &pane{
		style: style.GetPane(),
	}
}

func (p pane) update(msg tea.Msg) (pane, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case readDirMsg:
		p.readDir(msg.getDir())
		cmd = sendRefreshMsg
	}

	return p, cmd
}

func (p pane) view() string {
	var l string

	println("ok")
	if p.files != nil {
		for _, file := range p.files {
			l += fmt.Sprintf("%s\n", file.Name())
		}
	} else {
		l = "Empty"
	}

	return p.style.Render(l)
}

func (p *pane) readDir(dir string) {
	if files, err := ioutil.ReadDir(dir); err != nil {
		log.Fatal(err)
	} else {
		p.dir = dir
		p.files = files
	}
}
