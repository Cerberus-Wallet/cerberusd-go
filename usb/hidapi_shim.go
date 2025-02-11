//go:build linux || freebsd || openbsd
// +build linux freebsd openbsd

// shim for linux and freebsd so that cerberusd.go builds

package usb

import (
	"github.com/Cerberus-Wallet/cerberusd-go/core"
	"github.com/Cerberus-Wallet/cerberusd-go/memorywriter"
)

const HIDUse = false

type HIDAPI struct {
}

func InitHIDAPI(mw *memorywriter.MemoryWriter) (*HIDAPI, error) {
	return &HIDAPI{}, nil
}

func (b *HIDAPI) Enumerate() ([]core.USBInfo, error) {
	panic("not implemented for linux and freebsd")
}

func (b *HIDAPI) Has(path string) bool {
	panic("not implemented for linux and freebsd")
}

func (b *HIDAPI) Connect(path string, debug bool, reset bool) (core.USBDevice, error) {
	return &HID{}, nil
}

type HID struct {
}

func (d *HID) Close(disconnected bool) error {
	panic("not implemented for linux and freebsd")
}

func (d *HID) Write(buf []byte) (int, error) {
	panic("not implemented for linux and freebsd")
}

func (d *HID) Read(buf []byte) (int, error) {
	panic("not implemented for linux and freebsd")
}

func (b *HIDAPI) Close() {
	panic("not implemented for linux and freebsd")
}
