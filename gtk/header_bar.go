package gtk

type MockHeaderBar struct {
	MockContainer
}

func (*MockHeaderBar) SetSubtitle(v1 string) {
}

func (*MockHeaderBar) SetShowCloseButton(v1 bool) {
}

func (*MockHeaderBar) GetShowCloseButton() bool {
	return false
}
