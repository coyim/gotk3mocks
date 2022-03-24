package gtk

import "github.com/coyim/gotk3adapter/gtki"

type MockInfoBar struct {
	MockBox
}

func (m *MockInfoBar) GetOrientation() gtki.Orientation {
	return ret[gtki.Orientation](m.Called(), 0)
}

func (m *MockInfoBar) SetOrientation(v1 gtki.Orientation) {
	m.Called(v1)
}

func (m *MockInfoBar) AddActionWidget(v1 gtki.Widget, v2 gtki.ResponseType) {
	m.Called(v1, v2)
}

func (m *MockInfoBar) AddButton(v1 string, v2 gtki.ResponseType) {
	m.Called(v1, v2)
}

func (m *MockInfoBar) SetResponseSensitive(v1 gtki.ResponseType, v2 bool) {
	m.Called(v1, v2)
}

func (m *MockInfoBar) SetDefaultResponse(v1 gtki.ResponseType) {
	m.Called(v1)
}

func (m *MockInfoBar) SetMessageType(v1 gtki.MessageType) {
	m.Called(v1)
}

func (m *MockInfoBar) GetMessageType() gtki.MessageType {
	return ret[gtki.MessageType](m.Called(), 0)
}

func (m *MockInfoBar) GetActionArea() (gtki.Widget, error) {
	args := m.Called()
	return ret[gtki.Widget](args, 0), args.Error(1)
}

func (m *MockInfoBar) GetContentArea() (gtki.Box, error) {
	args := m.Called()
	return ret[gtki.Box](args, 0), args.Error(1)
}

func (m *MockInfoBar) GetShowCloseButton() bool {
	return m.Called().Bool(0)
}

func (m *MockInfoBar) SetShowCloseButton(v1 bool) {
	m.Called(v1)
}
