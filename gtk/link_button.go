package gtk

import "github.com/coyim/gotk3adapter/gtki"

type MockLinkButton struct {
	MockBin
}

func (m *MockLinkButton) GetUri() string {
	return m.Called().String(0)
}

func (m *MockLinkButton) SetUri(v1 string) {
	m.Called(v1)
}

func (m *MockLinkButton) SetImage(v1 gtki.Widget) {
	m.Called(v1)
}

func (m *MockLinkButton) GetLabel() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}
