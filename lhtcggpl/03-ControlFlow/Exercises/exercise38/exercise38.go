/*
Hands-on exercise #8
Create a program that uses a switch statement with no switch expression specified.
*/

package main

import "fmt"

func main() {

	switch {
	case x == "Moneypenny":
		fmt.Println(x)
	case x == "James Bond":
		fmt.Println("BONDBONDBOND")
	default:
		fmt.Println("neither")
	}

}
