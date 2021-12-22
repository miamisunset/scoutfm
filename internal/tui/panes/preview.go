package panes

type Preview struct {
	commonPane
}

func NewPreview(width int) Preview {
	w := int(float64(width)*(1.0-cwdPaneMaxWidth)) - 3

	return Preview{
		commonPane: newCommonPane(w),
	}
}
