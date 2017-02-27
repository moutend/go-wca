// +build amd64

package main

type PROPVARIANT struct {
	VT         uint16   // 2
	wReserved1 uint16   //  4
	wReserved2 uint16   //  6
	wReserved3 uint16   //  8
	Val        byte     // 9
	_          [15]byte // 24
}
