package gio

import (
	"github.com/coyim/gotk3adapter/gioi"
)

func init() {
	gioi.AssertGio(&Mock{})
	gioi.AssertResource(&MockResource{})
}
