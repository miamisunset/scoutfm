package tui

// Pane messages
type readDirMsg struct{ dir string }

func (m readDirMsg) getDir() string {
	return m.dir
}

type selectedFileMsg struct{ name string }

func (s selectedFileMsg) getName() string {
	return s.name
}
