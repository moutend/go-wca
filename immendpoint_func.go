// +build !windows

package wca

import (
	"github.com/go-ole/go-ole"
)

func mmeGetDataFlow(mme *IMMEndpoint, eDataFlow *uint32) (err error) {
	return ole.NewError(ole.E_NOTIMPL)
}
