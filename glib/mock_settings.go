package glib

import "github.com/coyim/gotk3adapter/glibi"

type MockSettings struct {
	MockObject
}

func (m *MockSettings) IsWritable(v1 string) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockSettings) Delay() {
	m.Called()
}

func (m *MockSettings) Apply() {
	m.Called()
}

func (m *MockSettings) Revert() {
	m.Called()
}

func (m *MockSettings) GetHasUnapplied() bool {
	return 	m.Called().Bool(0)
}

func (m *MockSettings) GetChild(v1 string) glibi.Settings {
	return ret[glibi.Settings](m.Called(v1), 0)
}

func (m *MockSettings) Reset(v1 string) {
	m.Called(v1)
}

func (m *MockSettings) ListChildren() []string {
	return ret[[]string](m.Called(), 0)
}

func (m *MockSettings) GetBoolean(v1 string) bool {
	return m.Called(v1).Bool(0)
}

func (m *MockSettings) SetBoolean(v1 string, v2 bool) bool {
	return m.Called(v1, v2).Bool(0)
}

func (m *MockSettings) GetInt(v1 string) int {
	return m.Called(v1).Int(0)
}

func (m *MockSettings) SetInt(v1 string, v2 int) bool {
	return m.Called(v1, v2).Bool(0)
}

func (m *MockSettings) GetUInt(v1 string) uint {
	return uint(m.Called(v1).Int(0))
}

func (m *MockSettings) SetUInt(v1 string, v2 uint) bool {
	return m.Called(v1, v2).Bool(0)
}

func (m *MockSettings) GetDouble(v1 string) float64 {
	return ret[float64](m.Called(v1), 0)
}

func (m *MockSettings) SetDouble(v1 string, v2 float64) bool {
	return m.Called(v1, v2).Bool(0)
}

func (m *MockSettings) GetString(v1 string) string {
	return m.Called(v1).String(0)
}

func (m *MockSettings) SetString(v1 string, v2 string) bool {
	return m.Called(v1, v2).Bool(0)
}

func (m *MockSettings) GetEnum(v1 string) int {
	return m.Called(v1).Int(0)
}

func (m *MockSettings) SetEnum(v1 string, v2 int) bool {
	return m.Called(v1, v2).Bool(0)
}

func (m *MockSettings) GetFlags(v1 string) uint {
	return uint(m.Called(v1).Int(0))
}

func (m *MockSettings) SetFlags(v1 string, v2 uint) bool {
	return m.Called(v1, v2).Bool(0)
}
