/*
String types
------------

A string type represents the set of string values. A string value is a (possibly empty) sequence of bytes. Strings are immutable: once created, it is impossible to change the contents of a string. The predeclared string type is string; it is a defined type.

The length of a string s (its size in bytes) can be discovered using the built-in function len. The length is a compile-time constant if the string is a constant. A string's bytes can be accessed by integer indices 0 through len(s)-1. It is illegal to take the address of such an element; if s[i] is the i'th byte of a string, &s[i] is invalid.

https://blog.golang.org/strings
*/
package StringTypes

import "fmt"

func main() {

	s := "Example"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	// The len is the size in bytes
	fmt.Printf("A string of legth %v\n", len(s))

	// Show UTF codification
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}

	fmt.Println("")

	for i, v := range s {
		fmt.Println(i, v)
	}
}
