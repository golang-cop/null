package Null

import "testing"

func TestNewReturnsNonNilInterface(t *testing.T) {
	got := New()
	if got == nil {
		t.Fatal("New() returned nil; want a non-nil null object")
	}
}

func TestNewReturnsDataPointer(t *testing.T) {
	got, ok := New().(*data)
	if !ok {
		t.Fatalf("New() returned %T; want *data", New())
	}
	if got == nil {
		t.Fatal("New() returned a nil *data pointer")
	}
}

func TestIsNull(t *testing.T) {
	d := data{}
	if !d.IsNull() {
		t.Error("data.IsNull() = false; want true")
	}
}

func TestIsNullThroughInterface(t *testing.T) {
	if !New().IsNull() {
		t.Error("New().IsNull() = false; want true")
	}
}
