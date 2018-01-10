package fizzbuzz

import "strconv"

func fizzbuzz(a int) string {
	mod3equal0 := (a % 3) == 0
	mod5equal0 := (a % 5) == 0

	if mod3equal0 && mod5equal0 {
		return "FizzBuzz"
	}

	if mod3equal0 {
		return "Fizz"
	}

	if mod5equal0 {
		return "Buzz"
	}

	return strconv.Itoa(a)
}
