package tui

import (
	"io/fs"
)

// Pane messages
type readDirMsg struct{ dir string }

func (m readDirMsg) getDir() string {
	return m.dir
}

type selectedFileMsg struct {
	file fs.FileInfo
}

func (s selectedFileMsg) getName() string {
	return s.file.Name()
}
