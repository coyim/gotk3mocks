package gdk_mock

type MockEventButton struct {
	MockEvent
}

func (m *MockEventButton) Button() uint {
	return ret[uint](m.Called(), 0)
}

func (m *MockEventButton) Time() uint32 {
	return ret[uint32](m.Called(), 0)
}

func (m *MockEventButton) X() float64 {
	return ret[float64](m.Called(), 0)
}

func (m *MockEventButton) Y() float64 {
	return ret[float64](m.Called(), 0)
}
