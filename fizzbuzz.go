package fizzbuzz

import "strconv"

func fizzbuzz(a int) string {

	if a == 3 {
		return "fizz"
	}

	return strconv.Itoa(a)
}
