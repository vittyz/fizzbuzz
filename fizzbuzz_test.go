package fizzbuzz

import "testing"

func TestInputOne(t *testing.T) {
	result := fizzbuzz(1)
	if result != 1 {
		t.Error("Incorrect")
	}
}

func TestInputTwoExpectTwo(t *testing.T) {
	result := fizzbuzz(2)
	if result != 2 {
		t.Error("2 not equal 2")
	}
}
