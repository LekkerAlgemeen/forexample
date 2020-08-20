package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	y := Years(2)
	if y != 14 {
		t.Error("Expected", 14, "got", y)
	}
}

func ExampleYears() {
	fmt.Println(Years(2))
	// Output:
	// 14
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(2)
	}
}

func TestYearsTwo(t *testing.T) {
	y := YearsTwo(2)
	if y != 14 {
		t.Error("Expected", 14, "got", y)
	}
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(2))
	// Output:
	// 14
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(2)
	}
}
