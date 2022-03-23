package gtk_mock

import "github.com/coyim/gotk3adapter/gtki"
import "github.com/stretchr/testify/mock"

type MockTextIter struct {
	mock.Mock
}

func (m *MockTextIter) BackwardChar() bool {
	return m.Called().Bool(0)
}

func (m *MockTextIter) BackwardChars(v1 int) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockTextIter) BackwardCursorPosition() bool {
	return m.Called().Bool(0)
}

func (m *MockTextIter) BackwardCursorPositions(v1 int) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockTextIter) BackwardLine() bool {
	return m.Called().Bool(0)
}

func (m *MockTextIter) BackwardLines(v1 int) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockTextIter) BackwardToTagToggle(v1 gtki.TextTag) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockTextIter) BackwardVisibleCursorPosition() bool {
	return m.Called().Bool(0)
}

func (m *MockTextIter) BackwardVisibleCursorPositions(v1 int) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockTextIter) BackwardVisibleLine() bool {
	return m.Called().Bool(0)
}

func (m *MockTextIter) BackwardVisibleLines(v1 int) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockTextIter) CanInsert(v1 bool) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockTextIter) Compare(v1 gtki.TextIter) int {
	return m.Called(v1).Int(0)
}

func (m *MockTextIter) Editable(v1 bool) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockTextIter) EndsLine() bool {
	return m.Called().Bool(0)
}

func (m *MockTextIter) EndsSentence() bool {
	return m.Called().Bool(0)
}

func (m *MockTextIter) EndsTag(v1 gtki.TextTag) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockTextIter) EndsWord() bool {
	return m.Called().Bool(0)
}

func (m *MockTextIter) Equal(v1 gtki.TextIter) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockTextIter) ForwardChar() bool {
	return m.Called().Bool(0)
}

func (m *MockTextIter) ForwardChars(v1 int) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockTextIter) ForwardCursorPosition() bool {
	return m.Called().Bool(0)
}

func (m *MockTextIter) ForwardCursorPositions(v1 int) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockTextIter) ForwardLine() bool {
	return m.Called().Bool(0)
}

func (m *MockTextIter) ForwardLines(v1 int) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockTextIter) ForwardSentenceEnd() bool {
	return m.Called().Bool(0)
}

func (m *MockTextIter) ForwardSentenceEnds(v1 int) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockTextIter) ForwardToEnd() {
	m.Called()
}

func (m *MockTextIter) ForwardToLineEnd() bool {
	return m.Called().Bool(0)
}

func (m *MockTextIter) ForwardToTagToggle(v1 gtki.TextTag) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockTextIter) ForwardVisibleCursorPosition() bool {
	return m.Called().Bool(0)
}

func (m *MockTextIter) ForwardVisibleCursorPositions(v1 int) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockTextIter) ForwardVisibleLine() bool {
	return m.Called().Bool(0)
}

func (m *MockTextIter) ForwardVisibleLines(v1 int) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockTextIter) ForwardVisibleWordEnd() bool {
	return m.Called().Bool(0)
}

func (m *MockTextIter) ForwardVisibleWordEnds(v1 int) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockTextIter) ForwardWordEnd() bool {
	return m.Called().Bool(0)
}

func (m *MockTextIter) ForwardWordEnds(v1 int) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockTextIter) GetBuffer() gtki.TextBuffer {
	return ret[gtki.TextBuffer](m.Called(), 0)
}

func (m *MockTextIter) GetBytesInLine() int {
	return m.Called().Int(0)
}

func (m *MockTextIter) GetChar() rune {
	return ret[rune](m.Called(), 0)
}

func (m *MockTextIter) GetCharsInLine() int {
	return m.Called().Int(0)
}

func (m *MockTextIter) GetLine() int {
	return m.Called().Int(0)
}

func (m *MockTextIter) GetLineIndex() int {
	return m.Called().Int(0)
}

func (m *MockTextIter) GetLineOffset() int {
	return m.Called().Int(0)
}

func (m *MockTextIter) GetOffset() int {
	return m.Called().Int(0)
}

func (m *MockTextIter) GetSlice(v1 gtki.TextIter) string {
	return m.Called(v1).String(0)
}

func (m *MockTextIter) GetText(v1 gtki.TextIter) string {
	return m.Called(v1).String(0)
}

func (m *MockTextIter) GetVisibleLineIndex() int {
	return m.Called().Int(0)
}

func (m *MockTextIter) GetVisibleLineOffset() int {
	return m.Called().Int(0)
}

func (m *MockTextIter) GetVisibleSlice(v1 gtki.TextIter) string {
	return m.Called(v1).String(0)
}

func (m *MockTextIter) GetVisibleText(v1 gtki.TextIter) string {
	return m.Called(v1).String(0)
}

func (m *MockTextIter) HasTag(v1 gtki.TextTag) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockTextIter) InRange(v1 gtki.TextIter, v2 gtki.TextIter) bool {
	return m.Called(v1, v2).Bool(0)
}

func (m *MockTextIter) InsideSentence() bool {
	return m.Called().Bool(0)
}

func (m *MockTextIter) InsideWord() bool {
	return m.Called().Bool(0)
}

func (m *MockTextIter) IsCursorPosition() bool {
	return m.Called().Bool(0)
}

func (m *MockTextIter) IsEnd() bool {
	return m.Called().Bool(0)
}

func (m *MockTextIter) IsStart() bool {
	return m.Called().Bool(0)
}

func (m *MockTextIter) SetLine(v1 int) {
	m.Called(v1)
}

func (m *MockTextIter) SetLineIndex(v1 int) {
	m.Called(v1)
}

func (m *MockTextIter) SetLineOffset(v1 int) {
	m.Called(v1)
}

func (m *MockTextIter) SetOffset(v1 int) {
	m.Called(v1)
}

func (m *MockTextIter) SetVisibleLineIndex(v1 int) {
	m.Called(v1)
}

func (m *MockTextIter) SetVisibleLineOffset(v1 int) {
	m.Called(v1)
}

func (m *MockTextIter) StartsLine() bool {
	return m.Called().Bool(0)
}

func (m *MockTextIter) StartsSentence() bool {
	return m.Called().Bool(0)
}

func (m *MockTextIter) StartsWord() bool {
	return m.Called().Bool(0)
}

func (m *MockTextIter) TogglesTag(v1 gtki.TextTag) bool {
	return m.Called(v1).Bool(0)
}
