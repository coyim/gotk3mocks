package gtk

type MockListBoxRow struct {
	MockBin
}

func (m *MockListBoxRow) GetIndex() int {
	return m.Called().Int(0)
}
