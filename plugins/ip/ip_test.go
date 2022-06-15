package main

import "testing"

func TestIpv4ToHexBasic1(t *testing.T) {
	got, _ := Ipv4ToHex("127.0.0.1")
	want := "0x7F000001"
	if got != want {
		t.Errorf("Result = %q, want %q", got, want)
	}
}

func TestIpv4ToHexBasic2(t *testing.T) {
	got, _ := Ipv4ToHex("149.76.12.4")
	want := "0x954C0C04"
	if got != want {
		t.Errorf("Result = %q, want %q", got, want)
	}
}

func TestIpv4ToHexError1(t *testing.T) {
	_, got := Ipv4ToHex("kaaaa boom")
	if got == nil {
		t.Errorf("Result must be an error")
	}
}

func TestIpv4ToHexError2(t *testing.T) {
	_, got := Ipv4ToHex("123.abc.100.100")
	if got == nil {
		t.Errorf("Result must be an error")
	}
}
