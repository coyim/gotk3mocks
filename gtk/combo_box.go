package gtk

import "github.com/coyim/gotk3adapter/gtki"

type MockComboBox struct {
	MockBin
}

func (m *MockComboBox) GetActiveIter() (gtki.TreeIter, error) {
	args := m.Called()
	return ret[gtki.TreeIter](args, 0), args.Error(1)
}

func (m *MockComboBox) GetActiveID() string {
	return m.Called().String(0)
}

func (m *MockComboBox) GetActive() int {
	return m.Called().Int(0)
}

func (m *MockComboBox) SetActive(v1 int) {
	m.Called(v1)
}

func (m *MockComboBox) SetModel(v1 gtki.TreeModel) {
	m.Called(v1)
}

func (m *MockComboBox) AddAttribute(v1 gtki.CellRenderer, v2 string, v3 int) {
	m.Called(v1, v2, v3)
}

func (m *MockComboBox) PackStart(v1 gtki.CellRenderer, v2 bool) {
	m.Called(v1, v2)
}

func (m *MockComboBox) SetIDColumn(v1 int) {
	m.Called(v1)
}

func (m *MockComboBox) SetEntryTextColumn(v1 int) {
	m.Called(v1)
}

func (m *MockComboBox) GetToggleButton() (gtki.Button, error) {
	args := m.Called()
	return ret[gtki.Button](args, 0), args.Error(1)
}
