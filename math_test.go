package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("Soma was incorrect, got: %d, want: %d.", total, 30)
	}
}
