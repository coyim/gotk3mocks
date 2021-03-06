package glib

import "github.com/coyim/gotk3adapter/glibi"

func init() {
	glibi.AssertGlib(&Mock{})
	glibi.AssertApplication(&MockApplication{})
	glibi.AssertObject(&MockObject{})
	glibi.AssertSettings(&MockSettings{})
	glibi.AssertSettingsBackend(&MockSettingsBackend{})
	glibi.AssertSettingsSchema(&MockSettingsSchema{})
	glibi.AssertSettingsSchemaSource(&MockSettingsSchemaSource{})
	glibi.AssertSignal(&MockSignal{})
	glibi.AssertValue(&MockValue{})
	glibi.AssertMenuModel(&MockMenuModel{})
	glibi.AssertMenu(&MockMenu{})
	glibi.AssertMenuItem(&MockMenuItem{})
}
