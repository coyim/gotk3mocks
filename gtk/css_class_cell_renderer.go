package gtk

import "github.com/coyim/gotk3adapter/gtki"

type MockCSSClassCellRenderer struct {
	MockCellRenderer
}

func (*MockCSSClassCellRenderer) SetReal(real gtki.CellRenderer) {
}
