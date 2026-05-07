// Copyright 2011 The win Authors. All rights reserved.
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

// for second parameter of WglSwapLayerBuffers
const (
	WGL_SWAP_MAIN_PLANE = (1 << 0)
	WGL_SWAP_OVERLAY1   = (1 << 1)
	WGL_SWAP_OVERLAY2   = (1 << 2)
	WGL_SWAP_OVERLAY3   = (1 << 3)
	WGL_SWAP_OVERLAY4   = (1 << 4)
	WGL_SWAP_OVERLAY5   = (1 << 5)
	WGL_SWAP_OVERLAY6   = (1 << 6)
	WGL_SWAP_OVERLAY7   = (1 << 7)
	WGL_SWAP_OVERLAY8   = (1 << 8)
	WGL_SWAP_OVERLAY9   = (1 << 9)
	WGL_SWAP_OVERLAY10  = (1 << 10)
	WGL_SWAP_OVERLAY11  = (1 << 11)
	WGL_SWAP_OVERLAY12  = (1 << 12)
	WGL_SWAP_OVERLAY13  = (1 << 13)
	WGL_SWAP_OVERLAY14  = (1 << 14)
	WGL_SWAP_OVERLAY15  = (1 << 15)
	WGL_SWAP_UNDERLAY1  = (1 << 16)
	WGL_SWAP_UNDERLAY2  = (1 << 17)
	WGL_SWAP_UNDERLAY3  = (1 << 18)
	WGL_SWAP_UNDERLAY4  = (1 << 19)
	WGL_SWAP_UNDERLAY5  = (1 << 20)
	WGL_SWAP_UNDERLAY6  = (1 << 21)
	WGL_SWAP_UNDERLAY7  = (1 << 22)
	WGL_SWAP_UNDERLAY8  = (1 << 23)
	WGL_SWAP_UNDERLAY9  = (1 << 24)
	WGL_SWAP_UNDERLAY10 = (1 << 25)
	WGL_SWAP_UNDERLAY11 = (1 << 26)
	WGL_SWAP_UNDERLAY12 = (1 << 27)
	WGL_SWAP_UNDERLAY13 = (1 << 28)
	WGL_SWAP_UNDERLAY14 = (1 << 29)
	WGL_SWAP_UNDERLAY15 = (1 << 30)
)

type (
	HGLRC HANDLE
)

type LAYERPLANEDESCRIPTOR struct {
	NSize           uint16
	NVersion        uint16
	DwFlags         uint32
	IPixelType      uint8
	CColorBits      uint8
	CRedBits        uint8
	CRedShift       uint8
	CGreenBits      uint8
	CGreenShift     uint8
	CBlueBits       uint8
	CBlueShift      uint8
	CAlphaBits      uint8
	CAlphaShift     uint8
	CAccumBits      uint8
	CAccumRedBits   uint8
	CAccumGreenBits uint8
	CAccumBlueBits  uint8
	CAccumAlphaBits uint8
	CDepthBits      uint8
	CStencilBits    uint8
	CAuxBuffers     uint8
	ILayerType      uint8
	BReserved       uint8
	CrTransparent   COLORREF
}

type POINTFLOAT struct {
	X, Y float32
}

type GLYPHMETRICSFLOAT struct {
	GmfBlackBoxX     float32
	GmfBlackBoxY     float32
	GmfptGlyphOrigin POINTFLOAT
	GmfCellIncX      float32
	GmfCellIncY      float32
}

var (
	// Library
	lib *windows.LazyDLL

	// Functions
	wglCopyContext            *windows.LazyProc
	wglCreateContext          *windows.LazyProc
	wglCreateLayerContext     *windows.LazyProc
	wglDeleteContext          *windows.LazyProc
	wglDescribeLayerPlane     *windows.LazyProc
	wglGetCurrentContext      *windows.LazyProc
	wglGetCurrentDC           *windows.LazyProc
	wglGetLayerPaletteEntries *windows.LazyProc
	wglGetProcAddress         *windows.LazyProc
	wglMakeCurrent            *windows.LazyProc
	wglRealizeLayerPalette    *windows.LazyProc
	wglSetLayerPaletteEntries *windows.LazyProc
	wglShareLists             *windows.LazyProc
	wglSwapLayerBuffers       *windows.LazyProc
	wglUseFontBitmaps         *windows.LazyProc
	wglUseFontOutlines        *windows.LazyProc
)

func init() {
	// Library
	lib = windows.NewLazySystemDLL("opengl32.dll")

	// Functions
	wglCopyContext = lib.NewProc("wglCopyContext")
	wglCreateContext = lib.NewProc("wglCreateContext")
	wglCreateLayerContext = lib.NewProc("wglCreateLayerContext")
	wglDeleteContext = lib.NewProc("wglDeleteContext")
	wglDescribeLayerPlane = lib.NewProc("wglDescribeLayerPlane")
	wglGetCurrentContext = lib.NewProc("wglGetCurrentContext")
	wglGetCurrentDC = lib.NewProc("wglGetCurrentDC")
	wglGetLayerPaletteEntries = lib.NewProc("wglGetLayerPaletteEntries")
	wglGetProcAddress = lib.NewProc("wglGetProcAddress")
	wglMakeCurrent = lib.NewProc("wglMakeCurrent")
	wglRealizeLayerPalette = lib.NewProc("wglRealizeLayerPalette")
	wglSetLayerPaletteEntries = lib.NewProc("wglSetLayerPaletteEntries")
	wglShareLists = lib.NewProc("wglShareLists")
	wglSwapLayerBuffers = lib.NewProc("wglSwapLayerBuffers")
	wglUseFontBitmaps = lib.NewProc("wglUseFontBitmapsW")
	wglUseFontOutlines = lib.NewProc("wglUseFontOutlinesW")
}

func WglCopyContext(hglrcSrc, hglrcDst HGLRC, mask uint) bool {
	ret, _, _ := syscall.SyscallN(wglCopyContext.Addr(),
		uintptr(hglrcSrc),
		uintptr(hglrcDst),
		uintptr(mask))

	return ret != 0
}

func WglCreateContext(hdc HDC) HGLRC {
	ret, _, _ := syscall.SyscallN(wglCreateContext.Addr(),
		uintptr(hdc))

	return HGLRC(ret)
}

func WglCreateLayerContext(hdc HDC, iLayerPlane int) HGLRC {
	ret, _, _ := syscall.SyscallN(wglCreateLayerContext.Addr(),
		uintptr(hdc),
		uintptr(iLayerPlane))

	return HGLRC(ret)
}

func WglDeleteContext(hglrc HGLRC) bool {
	ret, _, _ := syscall.SyscallN(wglDeleteContext.Addr(),
		uintptr(hglrc))

	return ret != 0
}

func WglDescribeLayerPlane(hdc HDC, iPixelFormat, iLayerPlane int, nBytes uint8, plpd *LAYERPLANEDESCRIPTOR) bool {
	ret, _, _ := syscall.SyscallN(wglDescribeLayerPlane.Addr(),
		uintptr(hdc),
		uintptr(iPixelFormat),
		uintptr(iLayerPlane),
		uintptr(nBytes),
		uintptr(unsafe.Pointer(plpd)))

	return ret != 0
}

func WglGetCurrentContext() HGLRC {
	ret, _, _ := syscall.SyscallN(wglGetCurrentContext.Addr())

	return HGLRC(ret)
}

func WglGetCurrentDC() HDC {
	ret, _, _ := syscall.SyscallN(wglGetCurrentDC.Addr())

	return HDC(ret)
}

func WglGetLayerPaletteEntries(hdc HDC, iLayerPlane, iStart, cEntries int, pcr *COLORREF) int {
	ret, _, _ := syscall.SyscallN(wglGetLayerPaletteEntries.Addr(),
		uintptr(hdc),
		uintptr(iLayerPlane),
		uintptr(iStart),
		uintptr(cEntries),
		uintptr(unsafe.Pointer(pcr)))

	return int(ret)
}

func WglGetProcAddress(lpszProc *byte) uintptr {
	ret, _, _ := syscall.SyscallN(wglGetProcAddress.Addr(),
		uintptr(unsafe.Pointer(lpszProc)))

	return uintptr(ret)
}

func WglMakeCurrent(hdc HDC, hglrc HGLRC) bool {
	ret, _, _ := syscall.SyscallN(wglMakeCurrent.Addr(),
		uintptr(hdc),
		uintptr(hglrc))

	return ret != 0
}

func WglRealizeLayerPalette(hdc HDC, iLayerPlane int, bRealize bool) bool {
	ret, _, _ := syscall.SyscallN(wglRealizeLayerPalette.Addr(),
		uintptr(hdc),
		uintptr(iLayerPlane),
		uintptr(BoolToBOOL(bRealize)))

	return ret != 0
}

func WglSetLayerPaletteEntries(hdc HDC, iLayerPlane, iStart, cEntries int, pcr *COLORREF) int {
	ret, _, _ := syscall.SyscallN(wglSetLayerPaletteEntries.Addr(),
		uintptr(hdc),
		uintptr(iLayerPlane),
		uintptr(iStart),
		uintptr(cEntries),
		uintptr(unsafe.Pointer(pcr)))

	return int(ret)
}

func WglShareLists(hglrc1, hglrc2 HGLRC) bool {
	ret, _, _ := syscall.SyscallN(wglShareLists.Addr(),
		uintptr(hglrc1),
		uintptr(hglrc2))

	return ret != 0
}

func WglSwapLayerBuffers(hdc HDC, fuPlanes uint) bool {
	ret, _, _ := syscall.SyscallN(wglSwapLayerBuffers.Addr(),
		uintptr(hdc),
		uintptr(fuPlanes))

	return ret != 0
}

func WglUseFontBitmaps(hdc HDC, first, count, listbase uint32) bool {
	ret, _, _ := syscall.SyscallN(wglUseFontBitmaps.Addr(),
		uintptr(hdc),
		uintptr(first),
		uintptr(count),
		uintptr(listbase))

	return ret != 0
}

func WglUseFontOutlines(hdc HDC, first, count, listbase uint32, deviation, extrusion float32, format int, pgmf *GLYPHMETRICSFLOAT) bool {
	ret, _, _ := syscall.SyscallN(wglUseFontOutlines.Addr(),
		uintptr(hdc),
		uintptr(first),
		uintptr(count),
		uintptr(listbase),
		uintptr(deviation),
		uintptr(extrusion),
		uintptr(format),
		uintptr(unsafe.Pointer(pgmf)))

	return ret != 0
}
