package fizzbuzz

import "strconv"

func fizzbuzz(a int) string {

	result := strconv.Itoa(a)

	if (a % 3) == 0 {
		result = "Fizz"
	}

	if (a % 5) == 0 {
		result = "Buzz"
	}

	if (a % 15) == 0 {
		result = "FizzBuzz"
	}

	return result
}
