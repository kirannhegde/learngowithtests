package dependency

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	bytes := bytes.Buffer{}
	Greet(&bytes, "Chris")

	got := bytes.String()
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
