package gtk

import "github.com/coyim/gotk3adapter/gtki"

type MockMenuShell struct {
	MockContainer
}

func (*MockMenuShell) Append(v1 gtki.MenuItem) {
}
