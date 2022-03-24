package gtk

import "github.com/coyim/gotk3adapter/gtki"

type MockNotebook struct {
	MockContainer
}

func (m *MockNotebook) NextPage() {
	m.Called()
}

func (m *MockNotebook) PrevPage() {
	m.Called()
}

func (m *MockNotebook) GetCurrentPage() int {
	return m.Called().Int(0)
}

func (m *MockNotebook) GetNPages() int {
	return m.Called().Int(0)
}

func (m *MockNotebook) SetCurrentPage(v1 int) {
	m.Called(v1)
}

func (m *MockNotebook) SetShowTabs(v1 bool) {
	m.Called(v1)
}

func (m *MockNotebook) AppendPage(v1, v2 gtki.Widget) int {
	return m.Called(v1, v2).Int(0)
}

func (m *MockNotebook) GetNthPage(v1 int) (gtki.Widget, error) {
	args := m.Called(v1)
	return ret[gtki.Widget](args, 0), args.Error(1)
}

func (m *MockNotebook) SetTabLabelText(v1 gtki.Widget, v2 string) {
	m.Called(v1, v2)
}
