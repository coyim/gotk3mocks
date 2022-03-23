package glib

import (
	"github.com/coyim/gotk3adapter/glibi"
	"github.com/stretchr/testify/mock"
)

type MockSettingsSchema struct {
	mock.Mock
}

func (m *MockSettingsSchema) Ref() glibi.SettingsSchema {
	return ret[glibi.SettingsSchema](m.Called(), 0)
}

func (m *MockSettingsSchema) Unref() {
	m.Called()
}

func (m *MockSettingsSchema) GetID() string {
	return m.Called().String(0)
}

func (m *MockSettingsSchema) GetPath() string {
	return m.Called().String(0)
}

func (m *MockSettingsSchema) HasKey(v1 string) bool {
	return m.Called(v1).Bool(0)
}
