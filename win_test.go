// Copyright 2026 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package win

import (
	"bytes"
	"encoding/binary"
	"testing"
	"unsafe"
)

func TestWordLongHelpers(t *testing.T) {
	if got := MAKEWORD(0x34, 0x12); got != 0x1234 {
		t.Fatalf("MAKEWORD(0x34, 0x12) = 0x%04x, want 0x1234", got)
	}
	if got := LOBYTE(0x1234); got != 0x34 {
		t.Fatalf("LOBYTE(0x1234) = 0x%02x, want 0x34", got)
	}
	if got := HIBYTE(0x1234); got != 0x12 {
		t.Fatalf("HIBYTE(0x1234) = 0x%02x, want 0x12", got)
	}

	if got := MAKELONG(0x5678, 0x1234); got != 0x12345678 {
		t.Fatalf("MAKELONG(0x5678, 0x1234) = 0x%08x, want 0x12345678", got)
	}
	if got := LOWORD(0x12345678); got != 0x5678 {
		t.Fatalf("LOWORD(0x12345678) = 0x%04x, want 0x5678", got)
	}
	if got := HIWORD(0x12345678); got != 0x1234 {
		t.Fatalf("HIWORD(0x12345678) = 0x%04x, want 0x1234", got)
	}

	if got := BoolToBOOL(true); got != 1 {
		t.Fatalf("BoolToBOOL(true) = %d, want 1", got)
	}
	if got := BoolToBOOL(false); got != 0 {
		t.Fatalf("BoolToBOOL(false) = %d, want 0", got)
	}

	if !SUCCEEDED(S_OK) || !SUCCEEDED(S_FALSE) {
		t.Fatal("SUCCEEDED should be true for S_OK and S_FALSE")
	}
	if FAILED(S_OK) || !FAILED(HRESULT(-1)) {
		t.Fatal("FAILED should be false for S_OK and true for -1")
	}

	if got := MAKEINTRESOURCE(42); uintptr(unsafe.Pointer(got)) != 42 {
		t.Fatalf("MAKEINTRESOURCE(42) = %p, want 0x2a", got)
	}
	func() {
		defer func() {
			if recover() == nil {
				t.Fatal("MAKEINTRESOURCE(-1) did not panic")
			}
		}()
		MAKEINTRESOURCE(-1)
	}()
	func() {
		defer func() {
			if recover() == nil {
				t.Fatal("MAKEINTRESOURCE(0x10000) did not panic")
			}
		}()
		MAKEINTRESOURCE(0x10000)
	}()
}

func TestBoolToVariantBool(t *testing.T) {
	vt := BoolToVariantBool(true)
	if vt.vt != VT_BOOL || vt.boolVal != VARIANT_TRUE {
		t.Fatalf("BoolToVariantBool(true) = {vt:%d, boolVal:%d}, want {vt:%d, boolVal:%d}", vt.vt, vt.boolVal, VT_BOOL, VARIANT_TRUE)
	}
	if !VariantBoolToBool(vt) {
		t.Fatal("VariantBoolToBool(true) = false, want true")
	}

	vf := BoolToVariantBool(false)
	if vf.vt != VT_BOOL || vf.boolVal != VARIANT_FALSE {
		t.Fatalf("BoolToVariantBool(false) = {vt:%d, boolVal:%d}, want {vt:%d, boolVal:%d}", vf.vt, vf.boolVal, VT_BOOL, VARIANT_FALSE)
	}
	if VariantBoolToBool(vf) {
		t.Fatal("VariantBoolToBool(false) = true, want false")
	}
}

func TestVariantAccessors(t *testing.T) {
	var v VARIANT

	v.SetLong(-12345)
	if got, err := v.Long(); err != nil || got != -12345 {
		t.Fatalf("Long() = (%d, %v), want (-12345, nil)", got, err)
	}

	v.SetULong(0xDEADBEEF)
	if got, err := v.ULong(); err != nil || got != 0xDEADBEEF {
		t.Fatalf("ULong() = (0x%x, %v), want (0xDEADBEEF, nil)", got, err)
	}

	v.SetBool(VARIANT_TRUE)
	if got, err := v.Bool(); err != nil || got != VARIANT_TRUE {
		t.Fatalf("Bool() = (%d, %v), want (%d, nil)", got, err, VARIANT_TRUE)
	}
}

func TestBSTR(t *testing.T) {
	if got := BSTRToString(nil); got != "" {
		t.Fatalf("BSTRToString(nil) = %q, want empty", got)
	}

	for _, tc := range []string{"", "hello", "繁體中文測試"} {
		bstr := SysAllocString(tc)
		if bstr == nil {
			t.Fatalf("SysAllocString(%q) returned nil", tc)
		}
		got := BSTRToString(bstr)
		SysFreeString(bstr)
		if got != tc {
			t.Fatalf("BSTR round-trip(%q) = %q", tc, got)
		}
	}
}

func TestTaskDialogPack(t *testing.T) {
	var dummy uint16 = 'A'
	btn := TASKDIALOG_BUTTON_UNPACKED{
		ButtonID:   42,
		ButtonText: &dummy,
	}
	var buf bytes.Buffer
	if err := btn.Pack(&buf); err != nil {
		t.Fatalf("TASKDIALOG_BUTTON_UNPACKED.Pack() error: %v", err)
	}
	if buf.Len() != sizeTASKDIALOG_BUTTON {
		t.Fatalf("packed button len = %d, want %d", buf.Len(), sizeTASKDIALOG_BUTTON)
	}
	if gotID := int32(binary.LittleEndian.Uint32(buf.Bytes()[0:4])); gotID != 42 {
		t.Fatalf("packed ButtonID = %d, want 42", gotID)
	}

	cfg := TASKDIALOGCONFIG_UNPACKED{
		Flags: TDF_ALLOW_DIALOG_CANCELLATION,
	}
	packed := cfg.Pack()
	raw := (*[sizeTASKDIALOGCONFIG]byte)(unsafe.Pointer(packed))
	if cbSize := binary.LittleEndian.Uint32(raw[0:4]); cbSize != sizeTASKDIALOGCONFIG {
		t.Fatalf("packed cbSize = %d, want %d", cbSize, sizeTASKDIALOGCONFIG)
	}
}
