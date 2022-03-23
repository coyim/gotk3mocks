package gtk

import "github.com/coyim/gotk3adapter/gtki"

type MockButton struct {
	MockBin
}

func (*MockButton) SetImage(v1 gtki.Widget) {
}

func (*MockButton) GetLabel() (string, error) {
	return "", nil
}

func (*MockButton) SetLabel(v1 string) {
}
