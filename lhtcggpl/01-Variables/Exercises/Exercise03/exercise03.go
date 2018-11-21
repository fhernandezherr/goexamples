/*
Using the code from the previous exercise,
At the package level scope, assign the following values to the three variables
for x assign 42
for y assign “James Bond”
for z assign true
in func main
use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER “s”
print out the value stored by variable “s”
*/
package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {

	// The compiler assigned the ZERO VALUES of each type
	s := fmt.Sprintf("x = %v, y = %v, z = %v", x, y, z)

	fmt.Println(s)
}
