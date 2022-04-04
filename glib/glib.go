package glib

import (
	"github.com/coyim/gotk3adapter/glibi"
	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

type getter interface {
	Get(int) interface{}
}

func ret[T any](args getter, index int) T {
	ret, _ := args.Get(index).(T)
	return ret
}

func (m *Mock) IdleAdd(f interface{}) glibi.SourceHandle {
	args := m.Called(f)
	return ret[glibi.SourceHandle](args, 0)
}

func (m *Mock) TimeoutAdd(milliseconds uint, f interface{}) glibi.SourceHandle {
	args := m.Called(milliseconds, f)
	return ret[glibi.SourceHandle](args, 0)
}

func (m *Mock) TimeoutSecondsAdd(milliseconds uint, f interface{}) glibi.SourceHandle {
	args := m.Called(milliseconds, f)
	return ret[glibi.SourceHandle](args, 0)
}

func (m *Mock) InitI18n(domain string, dir string) {
	m.Called(domain, dir)
}

func (m *Mock) Local(vx string) string {
	return m.Called(vx).String(0)
}

func (m *Mock) MainDepth() int {
	return m.Called().Int(0)
}

func (m *Mock) SignalNew(s string) (glibi.Signal, error) {
	args := m.Called(s)
	return ret[glibi.Signal](args, 0), args.Error(1)
}

func (m *Mock) SettingsNew(v string) glibi.Settings {
	args := m.Called(v)
	return ret[glibi.Settings](args, 0)
}

func (m *Mock) SettingsNewWithPath(v1 string, v2 string) glibi.Settings {
	args := m.Called(v1, v2)
	return ret[glibi.Settings](args, 0)
}

func (m *Mock) SettingsNewWithBackend(v1 string, v2 glibi.SettingsBackend) glibi.Settings {
	args := m.Called(v1, v2)
	return ret[glibi.Settings](args, 0)
}

func (m *Mock) SettingsNewWithBackendAndPath(v1 string, v2 glibi.SettingsBackend, v3 string) glibi.Settings {
	args := m.Called(v1, v2, v3)
	return ret[glibi.Settings](args, 0)
}

func (m *Mock) SettingsNewFull(v1 glibi.SettingsSchema, v2 glibi.SettingsBackend, v3 string) glibi.Settings {
	args := m.Called(v1, v2, v3)
	return ret[glibi.Settings](args, 0)
}

func (m *Mock) SettingsSync() {
	m.Called()
}

func (m *Mock) SettingsBackendGetDefault() glibi.SettingsBackend {
	return ret[glibi.SettingsBackend](m.Called(), 0)
}

func (m *Mock) KeyfileSettingsBackendNew(v1 string, v2 string, v3 string) glibi.SettingsBackend {
	return ret[glibi.SettingsBackend](m.Called(v1, v2, v3), 0)
}

func (m *Mock) MemorySettingsBackendNew() glibi.SettingsBackend {
	return ret[glibi.SettingsBackend](m.Called(), 0)
}

func (m *Mock) NullSettingsBackendNew() glibi.SettingsBackend {
	return ret[glibi.SettingsBackend](m.Called(), 0)
}

func (m *Mock) SettingsSchemaSourceGetDefault() glibi.SettingsSchemaSource {
	return ret[glibi.SettingsSchemaSource](m.Called(), 0)
}

func (m *Mock) SettingsSchemaSourceNewFromDirectory(v1 string, v2 glibi.SettingsSchemaSource, v3 bool) glibi.SettingsSchemaSource {
	return ret[glibi.SettingsSchemaSource](m.Called(v1, v2, v3), 0)
}

func (m *Mock) MenuNew() glibi.Menu {
	return ret[glibi.Menu](m.Called(), 0)
}

func (m *Mock) MenuItemNew(v1, v2 string) glibi.MenuItem {
	return ret[glibi.MenuItem](m.Called(v1, v2), 0)
}

func (m *Mock) MenuItemNewSection(v1 string, v2 glibi.MenuModel) glibi.MenuItem {
	return ret[glibi.MenuItem](m.Called(v1, v2), 0)
}

func (m *Mock) MenuItemNewSubmenu(v1 string, v2 glibi.MenuModel) glibi.MenuItem {
	return ret[glibi.MenuItem](m.Called(v1, v2), 0)
}

func (m *Mock) MenuItemNewFromModel(v1 glibi.MenuModel, v2 int) glibi.MenuItem {
	return ret[glibi.MenuItem](m.Called(v1, v2), 0)
}

func (m *Mock) ActionNameIsValid(v1 string) bool {
	return m.Called(v1).Bool(0)
}

func (m *Mock) SimpleActionNew(v1 string, v2 glibi.VariantType) glibi.SimpleAction {
	return ret[glibi.SimpleAction](m.Called(v1, v2), 0)
}

func (m *Mock) SimpleActionNewStateful(v1 string, v2 glibi.VariantType, v3 glibi.Variant) glibi.SimpleAction {
	return ret[glibi.SimpleAction](m.Called(v1, v2, v3), 0)
}

func (m *Mock) PropertyActionNew(v1 string, v2 glibi.Object, v3 string) glibi.PropertyAction {
	return ret[glibi.PropertyAction](m.Called(v1, v2, v3), 0)
}

func (m *Mock) SetFinalizerStrategy(v1 func(func())) {
	m.Called(v1)
}

func (m *Mock) MarkupEscapeText(v1 string) string {
	return m.Called(v1).String(0)
}
