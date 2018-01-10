package fizzbuzz

import "testing"

func TestInputOne(t *testing.T) {
	result := fizzbuzz(1)
	if result != 1 {
		t.Error("Incorrect")
	}
}
