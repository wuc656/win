// Copyright (c) Tailscale Inc & AUTHORS
// SPDX-License-Identifier: BSD-3-Clause

//go:build windows && !386

package win

import (
	"encoding/binary"
	"unsafe"
)

const (
	sizeTASKDIALOGCONFIG  = 0xA0
	sizeTASKDIALOG_BUTTON = 12
)

type packPtr uint64

var packByteOrder = binary.LittleEndian

func (u *TASKDIALOGCONFIG_UNPACKED) packPlatform(b *[sizeTASKDIALOGCONFIG]byte) {
	packByteOrder.PutUint32(b[0:4], uint32(sizeTASKDIALOGCONFIG))
	packByteOrder.PutUint64(b[4:12], uint64(u.HWNDParent))
	packByteOrder.PutUint64(b[12:20], uint64(u.HInstance))
	packByteOrder.PutUint32(b[20:24], uint32(u.Flags))
	packByteOrder.PutUint32(b[24:28], uint32(u.CommonButtons))
	packByteOrder.PutUint64(b[28:36], uint64(uintptr(unsafe.Pointer(u.WindowTitle))))
	packByteOrder.PutUint64(b[36:44], uint64(u.MainIcon))
	packByteOrder.PutUint64(b[44:52], uint64(uintptr(unsafe.Pointer(u.MainInstruction))))
	packByteOrder.PutUint64(b[52:60], uint64(uintptr(unsafe.Pointer(u.Content))))
	packByteOrder.PutUint32(b[60:64], uint32(u.CButtons))
	packByteOrder.PutUint64(b[64:72], uint64(uintptr(unsafe.Pointer(u.PButtons))))
	packByteOrder.PutUint32(b[72:76], uint32(u.DefaultButton))
	packByteOrder.PutUint32(b[76:80], uint32(u.CRadioButtons))
	packByteOrder.PutUint64(b[80:88], uint64(uintptr(unsafe.Pointer(u.PRadioButtons))))
	packByteOrder.PutUint32(b[88:92], uint32(u.DefaultRadioButton))
	packByteOrder.PutUint64(b[92:100], uint64(uintptr(unsafe.Pointer(u.VerificationText))))
	packByteOrder.PutUint64(b[100:108], uint64(uintptr(unsafe.Pointer(u.ExpandedInformation))))
	packByteOrder.PutUint64(b[108:116], uint64(uintptr(unsafe.Pointer(u.ExpandedControlText))))
	packByteOrder.PutUint64(b[116:124], uint64(uintptr(unsafe.Pointer(u.CollapsedControlText))))
	packByteOrder.PutUint64(b[124:132], uint64(u.FooterIcon))
	packByteOrder.PutUint64(b[132:140], uint64(uintptr(unsafe.Pointer(u.Footer))))
	packByteOrder.PutUint64(b[140:148], uint64(u.Callback))
	packByteOrder.PutUint64(b[148:156], uint64(u.CallbackData))
	packByteOrder.PutUint32(b[156:160], uint32(u.Width))
}
