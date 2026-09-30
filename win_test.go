// Copyright 2010 The win Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows
// +build windows

package win

import (
	"testing"
	"unsafe"
)

func TestMAKEINTRESOURCE(t *testing.T) {
	tests := []struct {
		name     string
		id       int
		expected uintptr
	}{
		{
			name:     "valid positive",
			id:       123,
			expected: 123,
		},
		{
			name:     "valid negative (truncates properly)",
			id:       -1,
			expected: 65535, // 0xFFFF
		},
		{
			name:     "out of bounds (large positive)",
			id:       65537, // 0x10001
			expected: 1,
		},
		{
			name:     "out of bounds (large negative)",
			id:       -65535, // -0xFFFF
			expected: 1, // -0xFFFF = ...0001
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ptr := MAKEINTRESOURCE(tt.id)
			actual := uintptr(unsafe.Pointer(ptr))
			if actual != tt.expected {
				t.Errorf("MAKEINTRESOURCE(%d) = %v, want %v", tt.id, actual, tt.expected)
			}
		})
	}
}
