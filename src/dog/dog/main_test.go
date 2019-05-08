package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	value := 5
	x := Years(value)
	if x != 50 {
		t.Error("Expected", value, "Got", x)
	}

}

func ExampleYears() {
	fmt.Println(Years(10))
	// Output:
	// 100
}

func BenchmarkYears(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Years(10)
	}

}
