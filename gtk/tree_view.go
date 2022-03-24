package gtk

import "github.com/coyim/gotk3adapter/gtki"

type MockTreeView struct {
	MockContainer
}

func (m *MockTreeView) RowExpanded(v1 gtki.TreePath) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockTreeView) ExpandRow(v1 gtki.TreePath, v2 bool) bool {
	return m.Called(v1, v2).Bool(0)
}

func (m *MockTreeView) CollapseRow(v1 gtki.TreePath) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockTreeView) ExpandAll() {
	m.Called()
}

func (m *MockTreeView) GetCursor() (gtki.TreePath, gtki.TreeViewColumn) {
	args := m.Called()
	return ret[gtki.TreePath](args, 0), ret[gtki.TreeViewColumn](args, 1)
}

func (m *MockTreeView) GetSelection() (gtki.TreeSelection, error) {
	args := m.Called()
	return ret[gtki.TreeSelection](args, 0), args.Error(1)
}

func (m *MockTreeView) GetPathAtPos(v1 int, v2 int) (gtki.TreePath, gtki.TreeViewColumn, int, int, bool) {
	args := m.Called(v1, v2)
	return ret[gtki.TreePath](args, 0), ret[gtki.TreeViewColumn](args, 1), args.Int(2), args.Int(3), args.Bool(4)
}

func (m *MockTreeView) SetEnableSearch(v1 bool) {
	m.Called(v1)
}

func (m *MockTreeView) GetEnableSearch() bool {
	return m.Called().Bool(0)
}

func (m *MockTreeView) SetSearchColumn(v1 int) {
	m.Called(v1)
}

func (m *MockTreeView) GetSearchColumn() int {
	return m.Called().Int(0)
}

func (m *MockTreeView) SetSearchEntry(v1 gtki.Entry) {
	m.Called(v1)
}

func (m *MockTreeView) GetSearchEntry() gtki.Entry {
	return ret[gtki.Entry](m.Called(), 0)
}

func (m *MockTreeView) SetSearchEqualSubstringMatch() {
	m.Called()
}

func (m *MockTreeView) GetModel() (gtki.TreeModel, error) {
	args := m.Called()
	return ret[gtki.TreeModel](args, 0), args.Error(1)
}

func (m *MockTreeView) SetModel(v1 gtki.TreeModel) {
	m.Called(v1)
}

func (m *MockTreeView) SetCursorOnCell(v1 gtki.TreePath, v2 gtki.TreeViewColumn, v3 gtki.CellRenderer, v4 bool) {
	m.Called(v1, v2, v3, v4)
}
