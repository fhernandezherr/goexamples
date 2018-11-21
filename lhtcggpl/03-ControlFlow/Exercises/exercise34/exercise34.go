/*
Hands-on exercise #4
Create a for loop using this syntax
	for { }
Have it print out the years you have been alive.
*/

package main

import "fmt"

func main() {

	bd := 1973
	for {

		if bd > 2018 {
			break
		}

		fmt.Println(bd)

		bd++
	}

}
