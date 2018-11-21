/*
Building on the code from the previous example
1. at the package level scope, using the “var” keyword, create a
   VARIABLE with the IDENTIFIER “y”. The variable should be of the
   UNDERLYING TYPE of your custom TYPE “x”
2. in func main
	a. this should already be done
		i. print out the value of the variable “x”
		ii. print out the type of the variable “x”
		iii. assign your own VALUE to the VARIABLE “x” using the “=”
             OPERATOR
		iv. print out the value of the variable “x”
	b. now do this
		i. now use CONVERSION to convert the TYPE of the VALUE stored
           in “x” to the UNDERLYING TYPE then use the “=” operator to
           ASSIGN that value to “y” print out the value stored in “y”
 */
package main

import "fmt"

type NewType int
var x NewType
var y int

func main() {
	fmt.Printf("%v\t%T\n", x, x)

	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
}
