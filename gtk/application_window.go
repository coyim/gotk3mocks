package gtk

import "github.com/coyim/gotk3adapter/glibi"

type MockApplicationWindow struct {
	MockWindow
}

func (m *MockApplicationWindow) SetShowMenubar(v1 bool) {
	m.Called(v1)
}

func (m *MockApplicationWindow) GetShowMenubar() bool {
	return m.Called().Bool(0)
}

func (m *MockApplicationWindow) GetID() uint {
	return ret[uint](m.Called(), 0)
}

func (m *MockApplicationWindow) LookupAction(v1 string) glibi.Action {
	return ret[glibi.Action](m.Called(v1), 0)
}

func (m *MockApplicationWindow) AddAction(v1 glibi.Action) {
	m.Called(v1)
}

func (m *MockApplicationWindow) RemoveAction(v1 string) {
	m.Called(v1)
}

func (m *MockApplicationWindow) HasAction(v1 string) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockApplicationWindow) GetActionEnabled(v1 string) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockApplicationWindow) GetActionParameterType(v1 string) glibi.VariantType {
	return ret[glibi.VariantType](m.Called(v1), 0)
}

func (m *MockApplicationWindow) GetActionStateType(v1 string) glibi.VariantType {
	return ret[glibi.VariantType](m.Called(v1), 0)
}

func (m *MockApplicationWindow) GetActionState(v1 string) glibi.Variant {
	return ret[glibi.Variant](m.Called(v1), 0)
}

func (m *MockApplicationWindow) GetActionStateHint(v1 string) glibi.Variant {
	return ret[glibi.Variant](m.Called(v1), 0)
}

func (m *MockApplicationWindow) ChangeActionState(v1 string, v2 glibi.Variant) {
	m.Called(v1, v2)
}

func (m *MockApplicationWindow) Activate(v1 string, v2 glibi.Variant) {
	m.Called(v1, v2)
}
