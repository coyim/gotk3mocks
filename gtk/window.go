package gtk

import (
	"github.com/coyim/gotk3adapter/gdki"
	"github.com/coyim/gotk3adapter/gtki"
)

type MockWindow struct {
	MockBin
}

func (m *MockWindow) AddAccelGroup(v1 gtki.AccelGroup) {
	m.Called(v1)
}

func (m *MockWindow) GetTitle() string {
	return m.Called().String(0)
}

func (m *MockWindow) HasToplevelFocus() bool {
	return m.Called().Bool(0)
}

func (m *MockWindow) Fullscreen() {
	m.Called()
}

func (m *MockWindow) Unfullscreen() {
	m.Called()
}

func (m *MockWindow) IsActive() bool {
	return m.Called().Bool(0)
}

func (m *MockWindow) Resize(v1, v2 int) {
	m.Called(v1, v2)
}

func (m *MockWindow) SetApplication(v1 gtki.Application) {
	m.Called(v1)
}

func (m *MockWindow) SetIcon(v1 gdki.Pixbuf) {
	m.Called(v1)
}

func (m *MockWindow) SetTitle(v1 string) {
	m.Called(v1)
}

func (m *MockWindow) SetTitlebar(v1 gtki.Widget) {
	m.Called(v1)
}

func (m *MockWindow) SetTransientFor(v1 gtki.Window) {
	m.Called(v1)
}

func (m *MockWindow) GetTransientFor() (gtki.Window, error) {
	args := m.Called()
	return ret[gtki.Window](args, 0), args.Error(1)
}

func (m *MockWindow) SetUrgencyHint(v1 bool) {
	m.Called(v1)
}

func (m *MockWindow) Present() {
	m.Called()
}

func (m *MockWindow) Iconify() {
	m.Called()
}

func (m *MockWindow) Deiconify() {
	m.Called()
}

func (m *MockWindow) Maximize() {
	m.Called()
}

func (m *MockWindow) Unmaximize() {
	m.Called()
}

func (m *MockWindow) AddMnemonic(v1 uint, v2 gtki.Widget) {
	m.Called(v1, v2)
}

func (m *MockWindow) RemoveMnemonic(v1 uint, v2 gtki.Widget) {
	m.Called(v1, v2)
}

func (m *MockWindow) ActivateMnemonic(v1 uint, v2 gdki.ModifierType) bool {
	return m.Called(v1, v2).Bool(0)
}

func (m *MockWindow) GetMnemonicModifier() gdki.ModifierType {
	return ret[gdki.ModifierType](m.Called(), 0)
}

func (m *MockWindow) SetMnemonicModifier(v1 gdki.ModifierType) {
	m.Called(v1)
}

func (m *MockWindow) SetDecorated(v1 bool) {
	m.Called(v1)
}

func (m *MockWindow) GetSize() (int, int) {
	args := m.Called()
	return args.Int(0), args.Int(1)
}

func (m *MockWindow) GetPosition() (int, int) {
	args := m.Called()
	return args.Int(0), args.Int(1)
}

func (m *MockWindow) Move(v1, v2 int) {
	m.Called(v1, v2)
}
