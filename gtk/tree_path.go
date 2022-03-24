package gtk

import "github.com/stretchr/testify/mock"

type MockTreePath struct {
	mock.Mock
}

func (m *MockTreePath) GetDepth() int {
	return m.Called().Int(0)
}

func (m *MockTreePath) String() string {
	return m.Called().String(0)
}
