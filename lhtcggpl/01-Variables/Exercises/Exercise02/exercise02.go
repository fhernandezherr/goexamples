// Hands-on exercise #2
// 1. Use var to DECLARE  three variables. The variables should have
//    package level scope. Do not assign VALUES to the variables. Use
//    the following IDENTIFIERS for the variables and make sure the
//    variables store values of the following type.
//        a. identifier "x" type int
//        b. identifier "y" type string
//        c. identifier "z" type bool
// 2. in func main
//        a. print out the values for each identifier
//        b. the compiler assigned values to the variables. What are
//           these values called?

package main

import "fmt"

var x int
var y string
var z bool

func main() {
	// The compiler assigned the ZERO VALUES of each type
	fmt.Println(x, y, z)
}
