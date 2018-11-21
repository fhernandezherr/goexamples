/*
Hands-on exercise #3
Create a for loop using this syntax
for condition { }
Have it print out the years you have been alive.
*/

package main

import "fmt"

func main() {

	bd := 1973
	for bd <= 2018 {
		fmt.Println(bd)
		bd++
	}

}
