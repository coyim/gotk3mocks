package gtk

import "github.com/coyim/gotk3adapter/gtki"

type MockBox struct {
	MockContainer
	gtki.Orientable
}

func (m *MockBox) PackEnd(v1 gtki.Widget, v2, v3 bool, v4 uint) {
	m.Called(v1, v2, v3, v4)
}

func (m *MockBox) PackStart(v1 gtki.Widget, v2, v3 bool, v4 uint) {
	m.Called(v1, v2, v3, v4)
}

func (m *MockBox) SetChildPacking(v1 gtki.Widget, v2, v3 bool, v4 uint, v5 gtki.PackType) {
	m.Called(v1, v2, v3, v4, v5)
}

func (m *MockBox) GetOrientation() gtki.Orientation {
	return ret[gtki.Orientation](m.Called(), 0)
}

func (m *MockBox) SetOrientation(v1 gtki.Orientation) {
	m.Called(v1)
}

func (m *MockBox) SetCenterWidget(v1 gtki.Widget) {
	m.Called(v1)
}

func (m *MockBox) GetCenterWidget() gtki.Widget {
	return ret[gtki.Widget](m.Called(), 0)
}

func (m *MockBox) SetFocusVAdjustment(v1 gtki.Adjustment) {
	m.Called(v1)
}
