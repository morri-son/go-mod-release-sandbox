package cli

import "testing"

func TestHello(t *testing.T) {
	if Hello() == "" {
		t.Fatal("empty")
	}
}
