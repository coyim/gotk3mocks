package gtk

import (
	"github.com/coyim/gotk3adapter/gtki"
	"github.com/coyim/gotk3mocks/glib"
)

type MockTextBuffer struct {
	glib.MockObject
}

func (m *MockTextBuffer) ApplyTagByName(v1 string, v2, v3 gtki.TextIter) {
	m.Called(v1, v2, v3)
}

func (m *MockTextBuffer) GetCharCount() int {
	return m.Called().Int(0)
}

func (m *MockTextBuffer) GetEndIter() gtki.TextIter {
	return ret[gtki.TextIter](m.Called(), 0)
}

func (m *MockTextBuffer) GetIterAtOffset(v1 int) gtki.TextIter {
	return ret[gtki.TextIter](m.Called(v1), 0)
}

func (m *MockTextBuffer) GetLineCount() int {
	return m.Called().Int(0)
}

func (m *MockTextBuffer) GetStartIter() gtki.TextIter {
	return ret[gtki.TextIter](m.Called(), 0)
}

func (m *MockTextBuffer) Insert(v1 gtki.TextIter, v2 string) {
	m.Called(v1, v2)
}

func (m *MockTextBuffer) InsertAtCursor(v1 string) {
	m.Called(v1)
}

func (m *MockTextBuffer) InsertWithTagByName(v1 gtki.TextIter, v2, v3 string) {
	m.Called(v1, v2, v3)
}

func (m *MockTextBuffer) GetText(v1 gtki.TextIter, v2 gtki.TextIter, v3 bool) string {
	return m.Called(v1, v2, v3).String(0)
}

func (m *MockTextBuffer) SetText(v1 string) {
	m.Called(v1)
}

func (m *MockTextBuffer) Delete(v1 gtki.TextIter, v2 gtki.TextIter) {
	m.Called(v1, v2)
}

func (m *MockTextBuffer) GetBounds() (gtki.TextIter, gtki.TextIter) {
	args := m.Called()
	return ret[gtki.TextIter](args, 0), ret[gtki.TextIter](args, 1)
}

func (m *MockTextBuffer) CreateMark(v1 string, v2 gtki.TextIter, v3 bool) gtki.TextMark {
	return ret[gtki.TextMark](m.Called(v1, v2, v3), 0)
}

func (m *MockTextBuffer) GetIterAtMark(v1 gtki.TextMark) gtki.TextIter {
	return ret[gtki.TextIter](m.Called(v1), 0)
}
