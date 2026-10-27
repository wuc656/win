// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build windows && 386

package win

import (
	"encoding/binary"
	"unsafe"
)

const (
	sizeTASKDIALOGCONFIG  = 0x60
	sizeTASKDIALOG_BUTTON = 8
)

type packPtr uint32

var packByteOrder = binary.LittleEndian

func (u *TASKDIALOGCONFIG_UNPACKED) packPlatform(b *[sizeTASKDIALOGCONFIG]byte) {
	packByteOrder.PutUint32(b[0:4], uint32(sizeTASKDIALOGCONFIG))
	packByteOrder.PutUint32(b[4:8], uint32(u.HWNDParent))
	packByteOrder.PutUint32(b[8:12], uint32(u.HInstance))
	packByteOrder.PutUint32(b[12:16], uint32(u.Flags))
	packByteOrder.PutUint32(b[16:20], uint32(u.CommonButtons))
	packByteOrder.PutUint32(b[20:24], uint32(uintptr(unsafe.Pointer(u.WindowTitle))))
	packByteOrder.PutUint32(b[24:28], uint32(u.MainIcon))
	packByteOrder.PutUint32(b[28:32], uint32(uintptr(unsafe.Pointer(u.MainInstruction))))
	packByteOrder.PutUint32(b[32:36], uint32(uintptr(unsafe.Pointer(u.Content))))
	packByteOrder.PutUint32(b[36:40], uint32(u.CButtons))
	packByteOrder.PutUint32(b[40:44], uint32(uintptr(unsafe.Pointer(u.PButtons))))
	packByteOrder.PutUint32(b[44:48], uint32(u.DefaultButton))
	packByteOrder.PutUint32(b[48:52], uint32(u.CRadioButtons))
	packByteOrder.PutUint32(b[52:56], uint32(uintptr(unsafe.Pointer(u.PRadioButtons))))
	packByteOrder.PutUint32(b[56:60], uint32(u.DefaultRadioButton))
	packByteOrder.PutUint32(b[60:64], uint32(uintptr(unsafe.Pointer(u.VerificationText))))
	packByteOrder.PutUint32(b[64:68], uint32(uintptr(unsafe.Pointer(u.ExpandedInformation))))
	packByteOrder.PutUint32(b[68:72], uint32(uintptr(unsafe.Pointer(u.ExpandedControlText))))
	packByteOrder.PutUint32(b[72:76], uint32(uintptr(unsafe.Pointer(u.CollapsedControlText))))
	packByteOrder.PutUint32(b[76:80], uint32(u.FooterIcon))
	packByteOrder.PutUint32(b[80:84], uint32(uintptr(unsafe.Pointer(u.Footer))))
	packByteOrder.PutUint32(b[84:88], uint32(u.Callback))
	packByteOrder.PutUint32(b[88:92], uint32(u.CallbackData))
	packByteOrder.PutUint32(b[92:96], uint32(u.Width))
}
