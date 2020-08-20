package word

import (
	"fmt"
	"testing"
)

func BenchmarkCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Count("test test")
	}
}

func ExampleCount() {
	fmt.Println(Count("test test"))
	// Output:
	// 2
}

func TestCount(t *testing.T) {
	c := Count("test test")
	if c != 2 {
		t.Error("Expected", 2, "got", c)
	}
}

func BenchmarkUseCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount("test test hello")
	}
}

func TestUseCount(t *testing.T) {
	m := UseCount("test test hello")
	if m["test"] != 2 {
		t.Error("Expected", 2, "got", m["test"])
	}
	if m["hello"] != 1 {
		t.Error("Expected", 1, "got", m["hello"])
	}
}
