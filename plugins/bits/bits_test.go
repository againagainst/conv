package main

import "testing"

func TestBitsBitsToBytes(t *testing.T) {
	got, _ := BytesConverter("8", "bit", "byte")
	want := "1"

	if got != want {
		t.Errorf("Result = %q, want %q", got, want)
	}
}

func TestBitsBitsToKb(t *testing.T) {
	got, _ := BytesConverter("9000", "bit", "kb")
	want := "1.0986328125"

	if got != want {
		t.Errorf("Result = %q, want %q", got, want)
	}
}

func TestBitsBitsToTb(t *testing.T) {
	got, _ := BytesConverter("4", "bit", "tb")
	want := "4.54747350886464e-13"

	if got != want {
		t.Errorf("Result = %q, want %q", got, want)
	}
}

func TestBitsTbToBytes(t *testing.T) {
	got, _ := BytesConverter("2", "tb", "byte")
	want := "2199023255552"

	if got != want {
		t.Errorf("Result = %q, want %q", got, want)
	}
}

func TestBitsMbToKb(t *testing.T) {
	got, _ := BytesConverter("0.5", "mb", "kb")
	want := "512"

	if got != want {
		t.Errorf("Result = %q, want %q", got, want)
	}
}

func TestBitsTbToTb(t *testing.T) {
	got, _ := BytesConverter("1.2345", "tb", "tb")
	want := "1.2345"

	if got != want {
		t.Errorf("Result = %q, want %q", got, want)
	}
}
