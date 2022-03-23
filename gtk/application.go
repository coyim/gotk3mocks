package gtk

import (
	"github.com/coyim/gotk3adapter/glibi"
	"github.com/coyim/gotk3adapter/gtki"
	"github.com/coyim/gotk3mocks/glib"
)

type MockApplication struct {
	glib.MockApplication
}

func (m *MockApplication) GetActiveWindow() gtki.Window {
	return ret[gtki.Window](m.Called(), 0)
}

func (m *MockApplication) AddWindow(v1 gtki.Window) {
	m.Called(v1)
}

func (m *MockApplication) RemoveWindow(v1 gtki.Window) {
	m.Called(v1)
}

func (m *MockApplication) PrefersAppMenu() bool {
	return m.Called().Bool(0)
}

func (m *MockApplication) GetAppMenu() glibi.MenuModel {
	return ret[glibi.MenuModel](m.Called(), 0)
}

func (m *MockApplication) SetAppMenu(v1 glibi.MenuModel) {
	m.Called(v1)
}

func (m *MockApplication) GetMenubar() glibi.MenuModel {
	return ret[glibi.MenuModel](m.Called(), 0)
}

func (m *MockApplication) SetMenubar(v1 glibi.MenuModel) {
	m.Called(v1)
}

func (m *MockApplication) LookupAction(v1 string) glibi.Action {
	return ret[glibi.Action](m.Called(v1), 0)
}

func (m *MockApplication) AddAction(v1 glibi.Action) {
	m.Called(v1)
}

func (m *MockApplication) RemoveAction(v1 string) {
	m.Called(v1)
}

func (m *MockApplication) HasAction(v1 string) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockApplication) GetActionEnabled(v1 string) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockApplication) GetActionParameterType(v1 string) glibi.VariantType {
	return ret[glibi.VariantType](m.Called(v1), 0)
}

func (m *MockApplication) GetActionStateType(v1 string) glibi.VariantType {
	return ret[glibi.VariantType](m.Called(v1), 0)
}

func (m *MockApplication) GetActionState(v1 string) glibi.Variant {
	return ret[glibi.Variant](m.Called(v1), 0)
}

func (m *MockApplication) GetActionStateHint(v1 string) glibi.Variant {
	return ret[glibi.Variant](m.Called(v1), 0)
}

func (m *MockApplication) ChangeActionState(v1 string, v2 glibi.Variant) {
	m.Called(v1, v2)
}

func (m *MockApplication) Activate(v1 string, v2 glibi.Variant) {
	m.Called(v1, v2)
}
