package glib

import (
	"github.com/coyim/gotk3adapter/glibi"
	"github.com/stretchr/testify/mock"
)

type MockSettingsSchemaSource struct {
	mock.Mock
}

func (m *MockSettingsSchemaSource) Ref() glibi.SettingsSchemaSource {
	return ret[glibi.SettingsSchemaSource](m.Called(), 0)
}

func (m *MockSettingsSchemaSource) Unref() {
	m.Called()
}

func (m *MockSettingsSchemaSource) Lookup(v1 string, v2 bool) glibi.SettingsSchema {
	return ret[glibi.SettingsSchema](m.Called(v1, v2), 0)
}
