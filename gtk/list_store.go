package gtk

import (
	"github.com/coyim/gotk3adapter/glibi"
	"github.com/coyim/gotk3adapter/gtki"
	"github.com/coyim/gotk3mocks/glib"
)

type MockListStore struct {
	glib.MockObject
}

func (m *MockListStore) Clear() {
	m.Called()
}

func (m *MockListStore) Append() gtki.TreeIter {
	return ret[gtki.TreeIter](m.Called(), 0)
}

func (m *MockListStore) Remove(v1 gtki.TreeIter) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockListStore) Set2(v1 gtki.TreeIter, v2 []int, v3 []interface{}) error {
	return m.Called(v1, v2, v3).Error(0)
}

func (m *MockListStore) SetValue(v1 gtki.TreeIter, v2 int, v3 interface{}) error {
	return m.Called(v1, v2, v3).Error(0)
}

func (m *MockListStore) GetIter(v1 gtki.TreePath) (gtki.TreeIter, error) {
	args := m.Called(v1)
	return ret[gtki.TreeIter](args, 0), args.Error(1)
}

func (m *MockListStore) GetIterFirst() (gtki.TreeIter, bool) {
	args := m.Called()
	return ret[gtki.TreeIter](args, 0), args.Bool(1)
}

func (m *MockListStore) GetIterFromString(v1 string) (gtki.TreeIter, error) {
	args := m.Called(v1)
	return ret[gtki.TreeIter](args, 0), args.Error(1)
}

func (m *MockListStore) GetPath(v1 gtki.TreeIter) (gtki.TreePath, error) {
	args := m.Called(v1)
	return ret[gtki.TreePath](args, 0), args.Error(1)
}

func (m *MockListStore) GetValue(v1 gtki.TreeIter, v2 int) (glibi.Value, error) {
	args := m.Called(v1, v2)
	return ret[glibi.Value](args, 0), args.Error(1)
}

func (m *MockListStore) IterNext(v1 gtki.TreeIter) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockListStore) GetColumnType(v1 int) glibi.Type {
	return ret[glibi.Type](m.Called(v1), 0)
}
