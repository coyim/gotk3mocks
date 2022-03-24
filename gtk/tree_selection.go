package gtk

import (
	"github.com/coyim/gotk3mocks/glib"
	"github.com/coyim/gotk3adapter/gtki"
)

type MockTreeSelection struct {
	glib.MockObject
}

func (m *MockTreeSelection) SelectIter(v1 gtki.TreeIter) {
	m.Called(v1)
}

func (m *MockTreeSelection) UnselectPath(v1 gtki.TreePath) {
	m.Called(v1)
}

func (m *MockTreeSelection) GetSelected() (gtki.TreeModel, gtki.TreeIter, bool) {
	args := m.Called()
	return ret[gtki.TreeModel](args, 0), ret[gtki.TreeIter](args, 1), args.Bool(2)
}

func (m *MockTreeSelection) GetSelectedRows(v1 gtki.TreeModel) []gtki.TreePath {
	return ret[[]gtki.TreePath](m.Called(v1), 0)
}
