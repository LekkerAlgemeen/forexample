package mymath

import (
	"fmt"
	"testing"
)

type test struct {
	data  []int
	value float64
}

func BenchmarkCenteredAvg(b *testing.B) {
	d := []int{1, 4, 6, 8, 100}
	for i := 0; i < b.N; i++ {
		CenteredAvg(d)
	}
}

func ExampleCenteredAvg() {
	xi := []int{1, 2, 2, 3}
	fmt.Println(CenteredAvg(xi))
	// Output:
	// 2
}

func TestCenteredAvg(t *testing.T) {
	a := test{
		data:  []int{1, 2, 3, 4, 5},
		value: 3,
	}
	b := test{
		data:  []int{0, 0, 0, 0},
		value: 0,
	}
	c := test{
		data:  []int{10, 20, 35000000},
		value: 20,
	}
	xt := []test{a, b, c}

	for _, v := range xt {
		ca := CenteredAvg(v.data)
		if ca != v.value {
			t.Error("Expected", v.value, "got", ca)
		}
	}
}
