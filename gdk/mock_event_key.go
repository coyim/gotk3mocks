package gdk_mock

type MockEventKey struct {
	MockEvent
}

func (m *MockEventKey) KeyVal() uint {
	return ret[uint](m.Called(), 0)
}

func (m *MockEventKey) State() uint {
	return ret[uint](m.Called(), 0)
}
