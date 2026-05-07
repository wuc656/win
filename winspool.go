// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

package win

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

// EnumPrinters flags
const (
	PRINTER_ENUM_DEFAULT     = 0x00000001
	PRINTER_ENUM_LOCAL       = 0x00000002
	PRINTER_ENUM_CONNECTIONS = 0x00000004
	PRINTER_ENUM_FAVORITE    = 0x00000004
	PRINTER_ENUM_NAME        = 0x00000008
	PRINTER_ENUM_REMOTE      = 0x00000010
	PRINTER_ENUM_SHARED      = 0x00000020
	PRINTER_ENUM_NETWORK     = 0x00000040
)

type PRINTER_INFO_4 struct {
	PPrinterName *uint16
	PServerName  *uint16
	Attributes   uint32
}

var (
	// Library
	libwinspool *windows.LazyDLL

	// Functions
	deviceCapabilities *windows.LazyProc
	documentProperties *windows.LazyProc
	enumPrinters       *windows.LazyProc
	getDefaultPrinter  *windows.LazyProc
)

func init() {
	// Library
	libwinspool = windows.NewLazySystemDLL("winspool.drv")

	// Functions
	deviceCapabilities = libwinspool.NewProc("DeviceCapabilitiesW")
	documentProperties = libwinspool.NewProc("DocumentPropertiesW")
	enumPrinters = libwinspool.NewProc("EnumPrintersW")
	getDefaultPrinter = libwinspool.NewProc("GetDefaultPrinterW")
}

func DeviceCapabilities(pDevice, pPort *uint16, fwCapability uint16, pOutput *uint16, pDevMode *DEVMODE) uint32 {
	ret, _, _ := syscall.SyscallN(deviceCapabilities.Addr(),
		uintptr(unsafe.Pointer(pDevice)),
		uintptr(unsafe.Pointer(pPort)),
		uintptr(fwCapability),
		uintptr(unsafe.Pointer(pOutput)),
		uintptr(unsafe.Pointer(pDevMode)))

	return uint32(ret)
}

func DocumentProperties(hWnd HWND, hPrinter HANDLE, pDeviceName *uint16, pDevModeOutput, pDevModeInput *DEVMODE, fMode uint32) int32 {
	ret, _, _ := syscall.SyscallN(documentProperties.Addr(),
		uintptr(hWnd),
		uintptr(hPrinter),
		uintptr(unsafe.Pointer(pDeviceName)),
		uintptr(unsafe.Pointer(pDevModeOutput)),
		uintptr(unsafe.Pointer(pDevModeInput)),
		uintptr(fMode))

	return int32(ret)
}

func EnumPrinters(flags uint32, name *uint16, level uint32, pPrinterEnum *byte, cbBuf uint32, pcbNeeded, pcReturned *uint32) bool {
	ret, _, _ := syscall.SyscallN(enumPrinters.Addr(),
		uintptr(flags),
		uintptr(unsafe.Pointer(name)),
		uintptr(level),
		uintptr(unsafe.Pointer(pPrinterEnum)),
		uintptr(cbBuf),
		uintptr(unsafe.Pointer(pcbNeeded)),
		uintptr(unsafe.Pointer(pcReturned)))

	return ret != 0
}

func GetDefaultPrinter(pszBuffer *uint16, pcchBuffer *uint32) bool {
	ret, _, _ := syscall.SyscallN(getDefaultPrinter.Addr(),
		uintptr(unsafe.Pointer(pszBuffer)),
		uintptr(unsafe.Pointer(pcchBuffer)))

	return ret != 0
}
