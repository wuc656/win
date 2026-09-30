package win

import (
	"testing"
)

func TestFAILED(t *testing.T) {
	if FAILED(0) {
		t.Errorf("FAILED(0) should be false")
	}
	if FAILED(1) {
		t.Errorf("FAILED(1) should be false")
	}
	if !FAILED(-1) {
		t.Errorf("FAILED(-1) should be true")
	}
	if FAILED(S_OK) {
		t.Errorf("FAILED(S_OK) should be false")
	}
	if FAILED(S_FALSE) {
		t.Errorf("FAILED(S_FALSE) should be false")
	}

	// Test failure codes using variables to avoid constant overflow errors
	var eFail uint32 = E_FAIL
	if !FAILED(HRESULT(int32(eFail))) {
		t.Errorf("FAILED(E_FAIL) should be true")
	}
	var eUnexpected uint32 = E_UNEXPECTED
	if !FAILED(HRESULT(int32(eUnexpected))) {
		t.Errorf("FAILED(E_UNEXPECTED) should be true")
	}
}

func TestSUCCEEDED(t *testing.T) {
	if !SUCCEEDED(0) {
		t.Errorf("SUCCEEDED(0) should be true")
	}
	if !SUCCEEDED(1) {
		t.Errorf("SUCCEEDED(1) should be true")
	}
	if SUCCEEDED(-1) {
		t.Errorf("SUCCEEDED(-1) should be false")
	}
	if !SUCCEEDED(S_OK) {
		t.Errorf("SUCCEEDED(S_OK) should be true")
	}
	if !SUCCEEDED(S_FALSE) {
		t.Errorf("SUCCEEDED(S_FALSE) should be true")
	}

	// Test failure codes
	var eFail uint32 = E_FAIL
	if SUCCEEDED(HRESULT(int32(eFail))) {
		t.Errorf("SUCCEEDED(E_FAIL) should be false")
	}
	var eUnexpected uint32 = E_UNEXPECTED
	if SUCCEEDED(HRESULT(int32(eUnexpected))) {
		t.Errorf("SUCCEEDED(E_UNEXPECTED) should be false")
	}
}
