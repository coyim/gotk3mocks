package gtk_mock

import (
	"github.com/coyim/gotk3adapter/gdki"
	"github.com/coyim/gotk3adapter/glibi"
	"github.com/coyim/gotk3adapter/gtki"
	"github.com/stretchr/testify/mock"
)

type Mock struct {
	mock.Mock
}

func ret[T any](args mock.Arguments, index int) T {
	var ret T
	v := args.Get(index)
	if v != nil {
		ret = v.(T)
	}
	return ret
}

func (m *Mock) AboutDialogNew() (gtki.AboutDialog, error) {
	args := m.Called()
	return ret[gtki.AboutDialog](args, 0), args.Error(1)
}

func (m *Mock) AccelGroupNew() (gtki.AccelGroup, error) {
	args := m.Called()
	return ret[gtki.AccelGroup](args, 0), args.Error(1)
}

func (m *Mock) AcceleratorParse(v1 string) (uint, gdki.ModifierType) {
	args := m.Called(v1)
	return ret[uint](args, 0), ret[gdki.ModifierType](args, 1)
}

func (m *Mock) AddProviderForScreen(v1 gdki.Screen, v2 gtki.StyleProvider, v3 uint) {
	m.Called(v1, v2, v3)
}

func (m *Mock) ApplicationNew(v1 string, v2 glibi.ApplicationFlags) (gtki.Application, error) {
	args := m.Called(v1, v2)
	return ret[gtki.Application](args, 0), args.Error(1)
}

func (m *Mock) AssistantNew() (gtki.Assistant, error) {
	args := m.Called()
	return ret[gtki.Assistant](args, 0), args.Error(1)
}

func (m *Mock) BuilderNew() (gtki.Builder, error) {
	args := m.Called()
	return ret[gtki.Builder](args, 0), args.Error(1)
}

func (m *Mock) BuilderNewFromResource(v1 string) (gtki.Builder, error) {
	args := m.Called(v1)
	return ret[gtki.Builder](args, 0), args.Error(1)
}

func (m *Mock) CellRendererTextNew() (gtki.CellRendererText, error) {
	args := m.Called()
	return ret[gtki.CellRendererText](args, 0), args.Error(1)
}

func (m *Mock) ButtonNewWithLabel(v1 string) (gtki.Button, error) {
	args := m.Called(v1)
	return ret[gtki.Button](args, 0), args.Error(1)
}

func (m *Mock) CheckButtonNew() (gtki.CheckButton, error) {
	args := m.Called()
	return ret[gtki.CheckButton](args, 0), args.Error(1)
}

func (m *Mock) CheckButtonNewWithMnemonic(v1 string) (gtki.CheckButton, error) {
	args := m.Called(v1)
	return ret[gtki.CheckButton](args, 0), args.Error(1)
}

func (m *Mock) CheckMenuItemNewWithMnemonic(v1 string) (gtki.CheckMenuItem, error) {
	args := m.Called(v1)
	return ret[gtki.CheckMenuItem](args, 0), args.Error(1)
}

func (m *Mock) CheckVersion(v1, v2, v3 uint) error {
	return m.Called(v1, v2, v3).Error(0)
}

func (m *Mock) ComboBoxNew() (gtki.ComboBox, error) {
	args := m.Called()
	return ret[gtki.ComboBox](args, 0), args.Error(1)
}

func (m *Mock) ComboBoxTextNew() (gtki.ComboBoxText, error) {
	args := m.Called()
	return ret[gtki.ComboBoxText](args, 0), args.Error(1)
}

func (m *Mock) CssProviderNew() (gtki.CssProvider, error) {
	args := m.Called()
	return ret[gtki.CssProvider](args, 0), args.Error(1)
}

func (m *Mock) CssProviderGetDefault() (gtki.CssProvider, error) {
	args := m.Called()
	return ret[gtki.CssProvider](args, 0), args.Error(1)
}

func (m *Mock) CssProviderGetNamed(v1 string, v2 string) (gtki.CssProvider, error) {
	args := m.Called(v1, v2)
	return ret[gtki.CssProvider](args, 0), args.Error(1)
}

func (m *Mock) EntryNew() (gtki.Entry, error) {
	args := m.Called()
	return ret[gtki.Entry](args, 0), args.Error(1)
}

func (m *Mock) EventBoxNew() (gtki.EventBox, error) {
	args := m.Called()
	return ret[gtki.EventBox](args, 0), args.Error(1)
}

func (m *Mock) ButtonBoxNew(v1 gtki.Orientation) (gtki.ButtonBox, error) {
	args := m.Called(v1)
	return ret[gtki.ButtonBox](args, 0), args.Error(1)
}

func (m *Mock) PopoverNew(v1 gtki.Widget) (gtki.Popover, error) {
	args := m.Called(v1)
	return ret[gtki.Popover](args, 0), args.Error(1)
}

func (m *Mock) BoxNew(v1 gtki.Orientation, v2 int) (gtki.Box, error) {
	args := m.Called(v1, v2)
	return ret[gtki.Box](args, 0), args.Error(1)
}

func (m *Mock) FileChooserDialogNewWith2Buttons(v1 string, v2 gtki.Window, v3 gtki.FileChooserAction, v4 string, v5 gtki.ResponseType, v6 string, v7 gtki.ResponseType) (gtki.FileChooserDialog, error) {
	args := m.Called(v1, v2, v3, v4, v5, v6, v7)
	return ret[gtki.FileChooserDialog](args, 0), args.Error(1)
}

func (m *Mock) GetMajorVersion() uint {
	return ret[uint](m.Called(), 0)
}

func (m *Mock) GetMinorVersion() uint {
	return ret[uint](m.Called(), 0)
}

func (m *Mock) GetMicroVersion() uint {
	return ret[uint](m.Called(), 0)
}

func (m *Mock) ImageNewFromFile(v1 string) (gtki.Image, error) {
	args := m.Called(v1)
	return ret[gtki.Image](args, 0), args.Error(1)
}

func (m *Mock) ImageNewFromResource(v1 string) (gtki.Image, error) {
	args := m.Called(v1)
	return ret[gtki.Image](args, 0), args.Error(1)
}

func (m *Mock) ImageNewFromPixbuf(v1 gdki.Pixbuf) (gtki.Image, error) {
	args := m.Called(v1)
	return ret[gtki.Image](args, 0), args.Error(1)
}

func (m *Mock) ImageNewFromIconName(v1 string, v2 gtki.IconSize) (gtki.Image, error) {
	args := m.Called(v1, v2)
	return ret[gtki.Image](args, 0), args.Error(1)
}

func (m *Mock) Init(v1 *[]string) {
	m.Called(v1)
}

func (m *Mock) InfoBarNew() (gtki.InfoBar, error) {
	args := m.Called()
	return ret[gtki.InfoBar](args, 0), args.Error(1)
}

func (m *Mock) LabelNew(v1 string) (gtki.Label, error) {
	args := m.Called(v1)
	return ret[gtki.Label](args, 0), args.Error(1)
}

func (m *Mock) ListStoreNew(v1 ...glibi.Type) (gtki.ListStore, error) {
	args := m.Called(v1)
	return ret[gtki.ListStore](args, 0), args.Error(1)
}

func (m *Mock) TreeStoreNew(v1 ...glibi.Type) (gtki.TreeStore, error) {
	args := m.Called(v1)
	return ret[gtki.TreeStore](args, 0), args.Error(1)
}

func (m *Mock) MenuBarNew() (gtki.MenuBar, error) {
	args := m.Called()
	return ret[gtki.MenuBar](args, 0), args.Error(1)
}

func (m *Mock) MenuItemNew() (gtki.MenuItem, error) {
	args := m.Called()
	return ret[gtki.MenuItem](args, 0), args.Error(1)
}

func (m *Mock) MenuItemNewWithMnemonic(v1 string) (gtki.MenuItem, error) {
	args := m.Called(v1)
	return ret[gtki.MenuItem](args, 0), args.Error(1)
}

func (m *Mock) MenuItemNewWithLabel(v1 string) (gtki.MenuItem, error) {
	args := m.Called(v1)
	return ret[gtki.MenuItem](args, 0), args.Error(1)
}

func (m *Mock) MenuNew() (gtki.Menu, error) {
	args := m.Called()
	return ret[gtki.Menu](args, 0), args.Error(1)
}

func (m *Mock) SearchBarNew() (gtki.SearchBar, error) {
	args := m.Called()
	return ret[gtki.SearchBar](args, 0), args.Error(1)
}

func (m *Mock) SearchEntryNew() (gtki.SearchEntry, error) {
	args := m.Called()
	return ret[gtki.SearchEntry](args, 0), args.Error(1)
}

func (m *Mock) SeparatorMenuItemNew() (gtki.SeparatorMenuItem, error) {
	args := m.Called()
	return ret[gtki.SeparatorMenuItem](args, 0), args.Error(1)
}

func (m *Mock) TextBufferNew(v1 gtki.TextTagTable) (gtki.TextBuffer, error) {
	args := m.Called(v1)
	return ret[gtki.TextBuffer](args, 0), args.Error(1)
}

func (m *Mock) TextTagNew(v1 string) (gtki.TextTag, error) {
	args := m.Called(v1)
	return ret[gtki.TextTag](args, 0), args.Error(1)
}

func (m *Mock) TextTagTableNew() (gtki.TextTagTable, error) {
	args := m.Called()
	return ret[gtki.TextTagTable](args, 0), args.Error(1)
}

func (m *Mock) TextViewNew() (gtki.TextView, error) {
	args := m.Called()
	return ret[gtki.TextView](args, 0), args.Error(1)
}

func (m *Mock) TreePathNew() gtki.TreePath {
	return ret[gtki.TreePath](m.Called(), 0)
}

func (m *Mock) WindowSetDefaultIcon(v1 gdki.Pixbuf) {
	m.Called(v1)
}

func (m *Mock) SettingsGetDefault() (gtki.Settings, error) {
	args := m.Called()
	return ret[gtki.Settings](args, 0), args.Error(1)
}

func (m *Mock) SeparatorNew(v1 gtki.Orientation) (gtki.Separator, error) {
	args := m.Called(v1)
	return ret[gtki.Separator](args, 0), args.Error(1)
}

func (m *Mock) EntryCompletionNew() (gtki.EntryCompletion, error) {
	args := m.Called()
	return ret[gtki.EntryCompletion](args, 0), args.Error(1)
}

func (m *Mock) StatusIconNew() (gtki.StatusIcon, error) {
	args := m.Called()
	return ret[gtki.StatusIcon](args, 0), args.Error(1)
}

func (m *Mock) StatusIconNewFromFile(v1 string) (gtki.StatusIcon, error) {
	args := m.Called(v1)
	return ret[gtki.StatusIcon](args, 0), args.Error(1)
}

func (m *Mock) StatusIconNewFromIconName(v1 string) (gtki.StatusIcon, error) {
	args := m.Called(v1)
	return ret[gtki.StatusIcon](args, 0), args.Error(1)
}

func (m *Mock) StatusIconNewFromPixbuf(v1 gdki.Pixbuf) (gtki.StatusIcon, error) {
	args := m.Called(v1)
	return ret[gtki.StatusIcon](args, 0), args.Error(1)
}

func (m *Mock) GetWidgetBuildableName(v1 gtki.Widget) (string, error) {
	args := m.Called(v1)
	return args.String(0), args.Error(1)
}

func (m *Mock) InfoBarSetRevealed(v1 gtki.InfoBar, v2 bool) {
	m.Called(v1, v2)
}

func (m *Mock) InfoBarGetRevealed(v1 gtki.InfoBar) bool {
	return m.Called(v1).Bool(0)
}
