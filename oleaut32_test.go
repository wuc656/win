// Copyright 2024 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

package win

import (
	"testing"
)

func TestVariantMustMethods(t *testing.T) {
	// Initialize a variant with a type that will intentionally mismatch in all Must* methods.
	// We use VT_EMPTY (0) as the mismatched type.
	v := &VARIANT{Vt: VT_EMPTY}

	// Test MustLong
	// Long() checks for VT_I4
	if val := v.MustLong(); val != 0 {
		t.Errorf("MustLong() on invalid variant returned %v, expected 0", val)
	}

	// Test MustULong
	// ULong() checks for VT_UI4
	if val := v.MustULong(); val != 0 {
		t.Errorf("MustULong() on invalid variant returned %v, expected 0", val)
	}

	// Test MustBool
	// Bool() checks for VT_BOOL
	if val := v.MustBool(); val != VARIANT_FALSE {
		t.Errorf("MustBool() on invalid variant returned %v, expected %v", val, VARIANT_FALSE)
	}

	// Test MustBSTR
	// BSTR() checks for VT_BSTR
	if val := v.MustBSTR(); val != nil {
		t.Errorf("MustBSTR() on invalid variant returned %v, expected nil", val)
	}

	// Test MustPDispatch
	// PDispatch() checks for VT_DISPATCH
	if val := v.MustPDispatch(); val != nil {
		t.Errorf("MustPDispatch() on invalid variant returned %v, expected nil", val)
	}

	// Test MustPVariant
	// PVariant() checks for VT_BYREF|VT_VARIANT
	if val := v.MustPVariant(); val != nil {
		t.Errorf("MustPVariant() on invalid variant returned %v, expected nil", val)
	}

	// Test MustPBool
	// PBool() checks for VT_BYREF|VT_BOOL
	if val := v.MustPBool(); val != nil {
		t.Errorf("MustPBool() on invalid variant returned %v, expected nil", val)
	}

	// Test MustPPDispatch
	// PPDispatch() checks for VT_BYREF|VT_DISPATCH
	if val := v.MustPPDispatch(); val != nil {
		t.Errorf("MustPPDispatch() on invalid variant returned %v, expected nil", val)
	}

	// Test MustPSafeArray
	// PSafeArray() checks for VT_ARRAY bit
	if val := v.MustPSafeArray(); val != nil {
		t.Errorf("MustPSafeArray() on invalid variant returned %v, expected nil", val)
	}
}
