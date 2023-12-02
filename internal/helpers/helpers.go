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
