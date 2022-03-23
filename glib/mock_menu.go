package glib

import "github.com/coyim/gotk3adapter/glibi"

type MockMenuModel struct {
	MockObject
}

func (m *MockMenuModel) IsMutable() bool {
	return m.Called().Bool(0)
}

func (m *MockMenuModel) GetNItems() int {
	return m.Called().Int(0)
}

func (m *MockMenuModel) GetItemLink(v1 int, v2 string) glibi.MenuModel {
	return ret[glibi.MenuModel](m.Called(v1, v2), 0)
}

func (m *MockMenuModel) ItemsChanged(v1, v2, v3 int) {
	m.Called(v1, v2, v3)
}

type MockMenu struct {
	MockMenuModel
}

func (m *MockMenu) Freeze() {
	m.Called()
}

func (m *MockMenu) Insert(v1 int, v2, v3 string) {
	m.Called(v1, v2, v3)
}

func (m *MockMenu) Prepend(v1, v2 string) {
	m.Called(v1, v2)
}

func (m *MockMenu) Append(v1, v2 string) {
	m.Called(v1, v2)
}

func (m *MockMenu) InsertItem(v1 int, v2 glibi.MenuItem) {
	m.Called(v1, v2)
}

func (m *MockMenu) AppendItem(v1 glibi.MenuItem) {
	m.Called(v1)
}

func (m *MockMenu) PrependItem(v1 glibi.MenuItem) {
	m.Called(v1)
}

func (m *MockMenu) InsertSection(v1 int, v2 string, v3 glibi.MenuModel) {
	m.Called(v1, v2, v3)
}

func (m *MockMenu) PrependSection(v1 string, v2 glibi.MenuModel) {
	m.Called(v1, v2)
}

func (m *MockMenu) AppendSection(v1 string, v2 glibi.MenuModel) {
	m.Called(v1, v2)
}

func (m *MockMenu) InsertSectionWithoutLabel(v1 int, v2 glibi.MenuModel) {
	m.Called(v1, v2)
}

func (m *MockMenu) PrependSectionWithoutLabel(v1 glibi.MenuModel) {
	m.Called(v1)
}

func (m *MockMenu) AppendSectionWithoutLabel(v1 glibi.MenuModel) {
	m.Called(v1)
}

func (m *MockMenu) InsertSubmenu(v1 int, v2 string, v3 glibi.MenuModel) {
	m.Called(v1, v2, v3)
}

func (m *MockMenu) PrependSubmenu(v1 string, v2 glibi.MenuModel) {
	m.Called(v1, v2)
}

func (m *MockMenu) AppendSubmenu(v1 string, v2 glibi.MenuModel) {
	m.Called(v1, v2)
}

func (m *MockMenu) Remove(v1 int) {
	m.Called(v1)
}

func (m *MockMenu) RemoveAll() {
	m.Called()
}

type MockMenuItem struct {
	MockObject
}

func (m *MockMenuItem) SetLabel(v1 string) {
	m.Called(v1)
}

func (m *MockMenuItem) SetDetailedAction(v1 string) {
	m.Called(v1)
}

func (m *MockMenuItem) SetSection(v1 glibi.MenuModel) {
	m.Called(v1)
}

func (m *MockMenuItem) SetSubmenu(v1 glibi.MenuModel) {
	m.Called(v1)
}

func (m *MockMenuItem) GetLink(v1 string) glibi.MenuModel {
	return ret[glibi.MenuModel](m.Called(v1), 0)
}

func (m *MockMenuItem) SetLink(v1 string, v2 glibi.MenuModel) {
	m.Called(v1, v2)
}
