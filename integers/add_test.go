package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	}

	t.Run("returns correct answer when adding two numbers", func(t *testing.T) {
		got := Add(4, 3)
		want := 7

		assertCorrectMessage(t, got, want)
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
