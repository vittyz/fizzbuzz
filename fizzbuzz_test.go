package fizzbuzz

import "testing"

func TestInputOne(t *testing.T) {
	result := fizzbuzz(1)
	if result != "1" {
		t.Error("Incorrect")
	}
}

func TestInputTwoExpectTwo(t *testing.T) {
	result := fizzbuzz(2)
	if result != "2" {
		t.Error("2 not equal 2")
	}
}

func TestInputThreeExpectFizz(t *testing.T) {
	result := fizzbuzz(3)
	if result != "Fizz" {
		t.Error("3 is not Fizz")
	}
}

func TestInput4Expect4(t *testing.T) {
	result := fizzbuzz(4)
	if result != "4" {
		t.Error("4 is not 4")
	}
}

func TestInput5ExpectBuzz(t *testing.T) {
	result := fizzbuzz(5)
	if result != "Buzz" {
		t.Error("5 is not Buzz")
	}
}

func TestInput6ExpectFizz(t *testing.T) {
	result := fizzbuzz(6)
	if result != "Fizz" {
		t.Error("6 is not Fizz")
	}
}

func TestInput7Expect7(t *testing.T) {
	result := fizzbuzz(7)
	if result != "7" {
		t.Error("7 is not 7")
	}
}

func TestInput8Expect8(t *testing.T) {
	result := fizzbuzz(8)
	if result != "8" {
		t.Error("8 is not 8")
	}
}

func TestInput9ExpectFizz(t *testing.T) {
	result := fizzbuzz(9)
	if result != "Fizz" {
		t.Error("9 is not Fizz")
	}
}

func TestInput10ExpectBuzz(t *testing.T) {
	result := fizzbuzz(10)
	if result != "Buzz" {
		t.Error("10 is not Buzz")
	}
}
