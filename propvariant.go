package wca

import "github.com/go-ole/go-ole"

type PROPVARIANT struct {
	ole.VARIANT
}

func (v PROPVARIANT) String() string {
	return pvString(v.Val)
}
