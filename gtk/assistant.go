package gtk

import (
	"github.com/coyim/gotk3adapter/gtki"
)

type MockAssistant struct {
	MockWindow
}

func (m *MockAssistant) Commit() {
	m.Called()
}

func (m *MockAssistant) NextPage() {
	m.Called()
}

func (m *MockAssistant) PreviousPage() {
	m.Called()
}

func (m *MockAssistant) SetCurrentPage(v1 int) {
	m.Called(v1)
}

func (m *MockAssistant) GetCurrentPage() int {
	return m.Called().Int(0)
}

func (m *MockAssistant) GetNthPage(v1 int) (gtki.Widget, error) {
	args := m.Called(v1)
	return ret[gtki.Widget](args, 0), args.Error(1)
}

func (m *MockAssistant) AppendPage(v1 gtki.Widget) int {
	return m.Called(v1).Int(0)
}

func (m *MockAssistant) SetPageType(v1 gtki.Widget, v2 gtki.AssistantPageType) {
	m.Called(v1, v2)
}

func (m *MockAssistant) GetPageType(v1 gtki.Widget) gtki.AssistantPageType {
	return ret[gtki.AssistantPageType](m.Called(v1), 0)
}

func (m *MockAssistant) SetPageComplete(v1 gtki.Widget, v2 bool) {
	m.Called(v1, v2)
}

func (m *MockAssistant) GetPageComplete(v1 gtki.Widget) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockAssistant) AddActionWidget(v1 gtki.Widget) {
	m.Called(v1)
}

func (m *MockAssistant) RemoveActionWidget(v1 gtki.Widget) {
	m.Called(v1)
}

func (m *MockAssistant) UpdateButtonsState() {
	m.Called()
}

func (m *MockAssistant) SetPageTitle(v1 gtki.Widget, v2 string) {
	m.Called(v1, v2)
}

func (m *MockAssistant) GetButtons() []gtki.Button {
	return ret[[]gtki.Button](m.Called(), 0)
}

func (m *MockAssistant) GetButtonSizeGroup() (gtki.SizeGroup, error) {
	args := m.Called()
	return ret[gtki.SizeGroup](args, 0), args.Error(1)
}

func (m *MockAssistant) GetHeaderBar() (gtki.HeaderBar, error) {
	args := m.Called()
	return ret[gtki.HeaderBar](args, 0), args.Error(1)
}

func (m *MockAssistant) GetSidebar() (gtki.Box, error) {
	args := m.Called()
	return ret[gtki.Box](args, 0), args.Error(1)
}

func (m *MockAssistant) GetNotebook() (gtki.Notebook, error) {
	args := m.Called()
	return ret[gtki.Notebook](args, 0), args.Error(1)
}

func (m *MockAssistant) HideBottomActionArea() {
	m.Called()
}
