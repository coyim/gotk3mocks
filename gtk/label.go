package gtk

import "github.com/coyim/gotk3adapter/pangoi"

type MockLabel struct {
	MockWidget
}

func (m *MockLabel) GetLabel() string {
	return m.Called().String(0)
}

func (m *MockLabel) SetLabel(v1 string) {
	m.Called(v1)
}

func (m *MockLabel) SetText(v1 string) {
	m.Called(v1)
}

func (m *MockLabel) SetMarkup(v1 string) {
	m.Called(v1)
}

func (m *MockLabel) SetSelectable(v1 bool) {
	m.Called(v1)
}

func (m *MockLabel) GetMnemonicKeyval() uint {
	return ret[uint](m.Called(), 0)
}

func (m *MockLabel) GetAttributes() (pangoi.AttrList, error) {
	args := m.Called()
	return ret[pangoi.AttrList](args, 0), args.Error(1)
}

func (m *MockLabel) SetAttributes(v1 pangoi.AttrList) {
	m.Called(v1)
}
