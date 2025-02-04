package bytemap_test

import (
	"fmt"

	"github.com/carlmjohnson/bytemap"
)

func ExampleRange() {
	ascii := bytemap.Range(0, 127)
	fmt.Println(ascii.Contains("Hello, world"))
	fmt.Println(ascii.Contains("Hello, 🌎"))

	upper := bytemap.Range('A', 'Z')
	nonupper := upper.Invert()
	fmt.Println(nonupper.Contains("hello, world!"))

	// Output:
	// true
	// false
	// true
}

func ExampleUnion() {
	upper := bytemap.Range('A', 'Z')
	lower := bytemap.Range('a', 'z')
	alpha := bytemap.Union(upper, lower)
	fmt.Println(alpha.Contains("CamelCase"))
	fmt.Println(alpha.Contains("snake_case"))

	// Output:
	// true
	// false
}
