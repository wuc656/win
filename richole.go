// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

package win

import (
	"syscall"
	"unsafe"
)

type REOBJECT struct {
	cbStruct uint32          // Size of structure
	cp       int32           // Character position of object
	clsid    CLSID           // Class ID of object
	poleobj  *IOleObject     // OLE object interface
	pstg     *IStorage       // Associated storage interface
	polesite *IOleClientSite // Associated client site interface
	sizel    SIZE            // Size of object (may be 0,0)
	dvaspect uint32          // Display aspect to use
	dwFlags  uint32          // Object status flags
	dwUser   uint32          // Dword for user's use
}

type IRichEditOleVtbl struct {
	IUnknownVtbl
	GetClientSite        uintptr
	GetObjectCount       uintptr
	GetLinkCount         uintptr
	GetObject            uintptr
	InsertObject         uintptr
	ConvertObject        uintptr
	ActivateAs           uintptr
	SetHostNames         uintptr
	SetLinkAvailable     uintptr
	SetDvaspect          uintptr
	HandsOffStorage      uintptr
	SaveCompleted        uintptr
	InPlaceDeactivate    uintptr
	ContextSensitiveHelp uintptr
	GetClipboardData     uintptr
	ImportDataObject     uintptr
}

type IRichEditOle struct {
	LpVtbl *IRichEditOleVtbl
}

func (obj *IRichEditOle) QueryInterface(riid REFIID, ppvObject *unsafe.Pointer) HRESULT {
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.QueryInterface,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(riid)),
		uintptr(unsafe.Pointer(ppvObject)))
	return HRESULT(ret)
}

func (obj *IRichEditOle) AddRef() uint32 {
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.AddRef,
		uintptr(unsafe.Pointer(obj)))
	return uint32(ret)
}

func (obj *IRichEditOle) Release() uint32 {
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.Release,
		uintptr(unsafe.Pointer(obj)))
	return uint32(ret)
}

func (obj *IRichEditOle) GetClientSite(lplpolesite **IOleClientSite) HRESULT {
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.GetClientSite,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(lplpolesite)))
	return HRESULT(ret)
}

func (obj *IRichEditOle) GetObjectCount() int32 {
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.GetObjectCount,
		uintptr(unsafe.Pointer(obj)))
	return int32(ret)
}

func (obj *IRichEditOle) GetLinkCount() int32 {
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.GetLinkCount,
		uintptr(unsafe.Pointer(obj)))
	return int32(ret)
}

func (obj *IRichEditOle) GetObject(iob int32, lpreobject *REOBJECT, dwFlags uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.GetObject,
		uintptr(unsafe.Pointer(obj)),
		uintptr(iob),
		uintptr(unsafe.Pointer(lpreobject)),
		uintptr(dwFlags))
	return HRESULT(ret)
}

func (obj *IRichEditOle) InsertObject(lpreobject *REOBJECT) HRESULT {
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.InsertObject,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(lpreobject)))
	return HRESULT(ret)
}

func (obj *IRichEditOle) ConvertObject(iob int32, rclsidNew REFCLSID, lpstrUserTypeNew *byte) HRESULT {
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.ConvertObject,
		uintptr(unsafe.Pointer(obj)),
		uintptr(iob),
		uintptr(unsafe.Pointer(rclsidNew)),
		uintptr(unsafe.Pointer(lpstrUserTypeNew)))
	return HRESULT(ret)
}

func (obj *IRichEditOle) ActivateAs(rclsid REFCLSID, rclsidAs REFCLSID) HRESULT {
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.ActivateAs,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(rclsid)),
		uintptr(unsafe.Pointer(rclsidAs)))
	return HRESULT(ret)
}

func (obj *IRichEditOle) SetHostNames(lpstrContainerApp *byte, lpstrContainerObj *byte) HRESULT {
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.SetHostNames,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(lpstrContainerApp)),
		uintptr(unsafe.Pointer(lpstrContainerObj)))
	return HRESULT(ret)
}

func (obj *IRichEditOle) SetLinkAvailable(iob int32, fAvailable BOOL) HRESULT {
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.SetLinkAvailable,
		uintptr(unsafe.Pointer(obj)),
		uintptr(iob),
		uintptr(fAvailable))
	return HRESULT(ret)
}

func (obj *IRichEditOle) SetDvaspect(iob int32, dvaspect uint32) HRESULT {
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.SetDvaspect,
		uintptr(unsafe.Pointer(obj)),
		uintptr(iob),
		uintptr(dvaspect))
	return HRESULT(ret)
}

func (obj *IRichEditOle) HandsOffStorage(iob int32) HRESULT {
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.HandsOffStorage,
		uintptr(unsafe.Pointer(obj)),
		uintptr(iob))
	return HRESULT(ret)
}

func (obj *IRichEditOle) SaveCompleted(iob int32, lpstg *IStorage) HRESULT {
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.SaveCompleted,
		uintptr(unsafe.Pointer(obj)),
		uintptr(iob),
		uintptr(unsafe.Pointer(lpstg)))
	return HRESULT(ret)
}

func (obj *IRichEditOle) InPlaceDeactivate() HRESULT {
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.InPlaceDeactivate,
		uintptr(unsafe.Pointer(obj)))
	return HRESULT(ret)
}

func (obj *IRichEditOle) ContextSensitiveHelp(fEnterMode BOOL) HRESULT {
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.ContextSensitiveHelp,
		uintptr(unsafe.Pointer(obj)),
		uintptr(fEnterMode))
	return HRESULT(ret)
}

func (obj *IRichEditOle) GetClipboardData(lpchrg *CHARRANGE, reco uint32, lplpdataobj **IDataObject) HRESULT {
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.GetClipboardData,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(lpchrg)),
		uintptr(reco),
		uintptr(unsafe.Pointer(lplpdataobj)))
	return HRESULT(ret)
}

func (obj *IRichEditOle) ImportDataObject(lpdataobj *IDataObject, cf CLIPFORMAT, hMetaPict HGLOBAL) HRESULT {
	ret, _, _ := syscall.SyscallN(obj.LpVtbl.ImportDataObject,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(lpdataobj)),
		uintptr(cf),
		uintptr(hMetaPict))
	return HRESULT(ret)
}
