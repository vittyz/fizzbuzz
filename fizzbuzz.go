package fizzbuzz

import "strconv"

func fizzbuzz(a int) string {

	result := strconv.Itoa(a)

	if a == 3 {
		result = "Fizz"
	}

	return result
}
