package panes

import (
	tea "github.com/charmbracelet/bubbletea"
	"io/ioutil"
	"path"
)

type Preview struct {
	commonPane

	err error
}

func NewPreview(width int) Preview {
	w := int(float64(width)*(1.0-cwdPaneMaxWidth)) - 3

	return Preview{
		commonPane: newCommonPane(w, false),
	}
}

func (p *Preview) Update(msg tea.Msg) (*Preview, tea.Cmd) {

	switch msg := msg.(type) {
	case fileMsg:
		p.currentDir = msg.cwd
		p.selectedFile = msg.file
		p.readFile()
	}

	return p, nil
}

func (p *Preview) readFile() {
	if p.selectedFile.IsDir() {
		files, err := ioutil.ReadDir(path.Join(p.currentDir, p.selectedFile.Name()))

		if err != nil {
			p.err = err
			return
		}

		p.Files = files
	} else {
		p.Files = nil
	}
}
