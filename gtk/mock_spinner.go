package gtk

import "github.com/coyim/gotk3adapter/gtki"

type MockSpinner struct {
	MockWidget
}

func (m *MockSpinner) Start() {
	m.Called()
}

func (m *MockSpinner) Stop() {
	m.Called()
}

func (m *Mock) SpinnerNew() (gtki.Spinner, error) {
	args := m.Called()
	return ret[gtki.Spinner](args, 0), args.Error(1)
}
