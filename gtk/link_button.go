package gtk

import "github.com/coyim/gotk3adapter/gtki"

type MockLinkButton struct {
	MockBin
}

func (*MockLinkButton) GetUri() string {
	return ""
}

func (*MockLinkButton) SetUri(uri string) {
}

func (*MockLinkButton) SetImage(v gtki.Widget) {
}

func (*MockLinkButton) GetLabel() (string, error) {
	return "", nil
}
