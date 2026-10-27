// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package win

import (
	"math"
	"unsafe"

	"golang.org/x/sys/windows"
)

type integer interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

const (
	S_OK           HRESULT = 0x00000000
	S_FALSE        HRESULT = 0x00000001
	E_UNEXPECTED   HRESULT = -0x7FFF0001 // 0x8000FFFF
	E_NOTIMPL      HRESULT = -0x7FFFBFFF // 0x80004001
	E_OUTOFMEMORY  HRESULT = -0x7FF8FFF2 // 0x8007000E
	E_INVALIDARG   HRESULT = -0x7FF8FFA9 // 0x80070057
	E_NOINTERFACE  HRESULT = -0x7FFFBFFE // 0x80004002
	E_POINTER      HRESULT = -0x7FFFBFFD // 0x80004003
	E_HANDLE       HRESULT = -0x7FF8FFFA // 0x80070006
	E_ABORT        HRESULT = -0x7FFFBFFC // 0x80004004
	E_FAIL         HRESULT = -0x7FFFBFFB // 0x80004005
	E_ACCESSDENIED HRESULT = -0x7FF8FFFB // 0x80070005
	E_PENDING      HRESULT = -0x7FFFFFF6 // 0x8000000A
)

const (
	FALSE = 0
	TRUE  = 1
)

type (
	BOOL    int32
	HRESULT int32
)

func SUCCEEDED(hr HRESULT) bool {
	return hr >= 0
}

func FAILED(hr HRESULT) bool {
	return hr < 0
}

func MAKEWORD(lo, hi byte) uint16 {
	return uint16(uint16(lo) | ((uint16(hi)) << 8))
}

func LOBYTE(w uint16) byte {
	return byte(w)
}

func HIBYTE(w uint16) byte {
	return byte(w >> 8 & 0xff)
}

func MAKELONG(lo, hi uint16) uint32 {
	return uint32(uint32(lo) | ((uint32(hi)) << 16))
}

func LOWORD(dw uint32) uint16 {
	return uint16(dw)
}

func HIWORD(dw uint32) uint16 {
	return uint16(dw >> 16 & 0xffff)
}

func UTF16PtrToString(s *uint16) string {
	return windows.UTF16PtrToString(s)
}

func MAKEINTRESOURCE[ID integer](id ID) *uint16 {
	if id < 0 || int64(id) > math.MaxUint16 {
		panic("resource id out of uint16 range")
	}
	return (*uint16)(unsafe.Pointer(uintptr(id)))
}

func MAKEINTATOM(atom ATOM) *uint16 {
	return MAKEINTRESOURCE(atom)
}

func BoolToBOOL(value bool) BOOL {
	if value {
		return 1
	}

	return 0
}
