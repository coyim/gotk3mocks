package gtk

type MockListBoxRow struct {
	MockBin
}

func (*MockListBoxRow) GetIndex() int {
	return 0
}
