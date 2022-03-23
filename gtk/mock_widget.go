package gtk_mock

import (
	"github.com/coyim/gotk3adapter/gdki"
	"github.com/coyim/gotk3adapter/glibi"
	"github.com/coyim/gotk3adapter/gtki"
	"github.com/coyim/gotk3mocks/glib"
)

type MockWidget struct {
	glib.MockObject
}

func (m *MockWidget) Map() {
	m.Called()
}

func (m *MockWidget) SetHExpand(v1 bool) {
	m.Called(v1)
}

func (m *MockWidget) SetVExpand(v1 bool) {
	m.Called(v1)
}

func (m *MockWidget) SetSensitive(v1 bool) {
	m.Called(v1)
}

func (m *MockWidget) IsSensitive() bool {
	return m.Called().Bool(0)
}

func (m *MockWidget) SetTooltipText(v1 string) {
	m.Called(v1)
}

func (m *MockWidget) SetVisible(v1 bool) {
	m.Called(v1)
}

func (m *MockWidget) IsVisible() bool {
	return m.Called().Bool(0)
}

func (m *MockWidget) SetName(v1 string) {
	m.Called(v1)
}

func (m *MockWidget) SetNoShowAll(v1 bool) {
	m.Called(v1)
}

func (m *MockWidget) SetMarginTop(v1 int) {
	m.Called(v1)
}

func (m *MockWidget) SetMarginBottom(v1 int) {
	m.Called(v1)
}

func (m *MockWidget) SetSizeRequest(v1, v2 int) {
	m.Called(v1, v2)
}

func (m *MockWidget) GetAllocatedHeight() int {
	return m.Called().Int(0)
}

func (m *MockWidget) GetName() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockWidget) GetAllocatedWidth() int {
	return m.Called().Int(0)
}

func (m *MockWidget) GetAllocation() gtki.Allocation {
	return ret[gtki.Allocation](m.Called(), 0)
}

func (m *MockWidget) GetParent() (gtki.Widget, error) {
	args := m.Called()
	return ret[gtki.Widget](args, 0), args.Error(1)
}

func (m *MockWidget) GetParentX() (gtki.Widget, error) {
	args := m.Called()
	return ret[gtki.Widget](args, 0), args.Error(1)
}

func (m *MockWidget) GrabFocus() {
	m.Called()
}

func (m *MockWidget) GrabDefault() {
	m.Called()
}

func (m *MockWidget) HasFocus() bool {
	return m.Called().Bool(0)
}

func (m *MockWidget) Hide() {
	m.Called()
}

func (m *MockWidget) HideOnDelete() {
	m.Called()
}

func (m *MockWidget) SetCanFocus(v1 bool) {
	m.Called(v1)
}

func (m *MockWidget) Show() {
	m.Called()
}

func (m *MockWidget) ShowAll() {
	m.Called()
}

func (m *MockWidget) GetWindow() (gdki.Window, error) {
	args := m.Called()
	return ret[gdki.Window](args, 0), args.Error(1)
}

func (m *MockWidget) GetStyleContext() (gtki.StyleContext, error) {
	args := m.Called()
	return ret[gtki.StyleContext](args, 0), args.Error(1)
}

func (m *MockWidget) SetHAlign(v1 gtki.Align) {
	m.Called(v1)
}

func (m *MockWidget) SetVAlign(v1 gtki.Align) {
	m.Called(v1)
}

func (m *MockWidget) SetOpacity(v1 float64) {
	m.Called(v1)
}

func (m *MockWidget) Destroy() {
	m.Called()
}

func (m *MockWidget) TemplateChild(v1 string) (glibi.Object, error) {
	args := m.Called()
	return ret[glibi.Object](args, 0), args.Error(1)
}
