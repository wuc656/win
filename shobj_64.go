// Copyright 2012 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (windows && amd64) || (windows && arm64)
// +build windows,amd64 windows,arm64

package win

import (
	"syscall"
	"unsafe"
)

func (obj *ITaskbarList3) SetProgressValue(hwnd HWND, current uint32, length uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.SetProgressValue,
		uintptr(unsafe.Pointer(obj)),
		uintptr(hwnd),
		uintptr(current),
		uintptr(length))

	return HRESULT(ret)
}
