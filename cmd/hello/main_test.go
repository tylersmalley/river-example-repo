package main

import "testing"

func TestGreeting(t *testing.T) {
	want := "Hello, River!"
	if got := greeting(); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
