// Helper functions that will be used from day to day.
package helpers

import (
	"os"
)

// Check for errors and panic if needed.
func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// Open filename and return its contents as a single string.
func Read(filename string) string {
	/* Thanks - https://gobyexample.com/reading-files  */
	contents, err := os.ReadFile(filename)
	Check(err)
	return string(contents)
}

// greatest common divisor (GCD) via Euclidean algorithm
// cred: https://siongui.github.io/2017/06/03/go-find-lcm-by-GCD/
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
// cred: https://siongui.github.io/2017/06/03/go-find-LCM-by-gcd/
func LCM(integers ...int) int {
	var a, b int
	a, integers = integers[0], integers[1:]
	b, integers = integers[0], integers[1:]
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
