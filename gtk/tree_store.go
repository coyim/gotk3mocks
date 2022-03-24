package gtk

import (
	"github.com/coyim/gotk3adapter/glibi"
	"github.com/coyim/gotk3adapter/gtki"
	"github.com/coyim/gotk3mocks/glib"
)

type MockTreeStore struct {
	glib.MockObject
}

func (m *MockTreeStore) GetIter(v1 gtki.TreePath) (gtki.TreeIter, error) {
	args := m.Called(v1)
	return ret[gtki.TreeIter](args, 0), args.Error(1)
}

func (m *MockTreeStore) GetIterFirst() (gtki.TreeIter, bool) {
	args := m.Called()
	return ret[gtki.TreeIter](args, 0), args.Bool(1)
}

func (m *MockTreeStore) GetIterFromString(v1 string) (gtki.TreeIter, error) {
	args := m.Called(v1)
	return ret[gtki.TreeIter](args, 0), args.Error(1)
}

func (m *MockTreeStore) GetPath(v1 gtki.TreeIter) (gtki.TreePath, error) {
	args := m.Called(v1)
	return ret[gtki.TreePath](args, 0), args.Error(1)
}

func (m *MockTreeStore) GetValue(v1 gtki.TreeIter, v2 int) (glibi.Value, error) {
	args := m.Called(v1, v2)
	return ret[glibi.Value](args, 0), args.Error(1)
}

func (m *MockTreeStore) IterNext(v1 gtki.TreeIter) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockTreeStore) Append(v1 gtki.TreeIter) gtki.TreeIter {
	return ret[gtki.TreeIter](m.Called(v1), 0)
}

func (m *MockTreeStore) Clear() {
	m.Called()
}

func (m *MockTreeStore) SetValue(v1 gtki.TreeIter, v2 int, v3 interface{}) error {
	return m.Called(v1, v2, v3).Error(0)
}

func (m *MockTreeStore) Remove(v1 gtki.TreeIter) bool {
	return m.Called(v1).Bool(0)
}
