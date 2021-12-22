package panes

import (
	"io/ioutil"
	"net/http"
	"os"
	"path"

	tea "github.com/charmbracelet/bubbletea"
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
		p.getContentType()
	}
}

func (p *Preview) getContentType() {
	f, _ := os.Open(path.Join(p.currentDir, p.selectedFile.Name()))
	defer f.Close()

	buf := make([]byte, 512)

	_, err := f.Read(buf)
	if err != nil {
		p.err = err
	}

	p.contentType = http.DetectContentType(buf)
}
