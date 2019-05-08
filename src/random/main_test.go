package thoughts

import (
	"fmt"
	"testing"
)

func TestHappy(t *testing.T) {
	value := "Tom"
	s := Happy(value)
	if s != "Happy day to you Tom" {
		t.Error("Expected", "Happy day to you", value, "Got", s)
	}

}

func TestSad(t *testing.T) {
	value := "Tom"
	s := Sad(value)
	if s != "Sad day to you Tom" {
		t.Error("Expected", "Sad day to you", value, "Got", s)
	}

}

func ExampleHappy() {
	fmt.Println(Happy("Tom"))
	// Output:
	// Happy day to you Tom
}

func ExampleSad() {
	fmt.Println(Sad("Tom"))
	// Output:
	// Sad day to you Tom
}

func BenchmarkHappy(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Happy("Tom")
	}
}

func BenchmarkSad(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Sad("Tom")
	}
}
