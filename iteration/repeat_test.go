package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 10)
	expected := "aaaaaaaaaa"
	if expected != repeated {
		t.Errorf("expected %q, but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
	output := Repeat("b", 10)
	fmt.Println(output)
	// Output: bbbbbbbbbb
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}
