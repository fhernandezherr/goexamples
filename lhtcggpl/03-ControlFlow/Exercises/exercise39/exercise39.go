/*
Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.
*/

package main

import "fmt"

func main() {

	favSport := "tennis"

	switch favSport {
	case "football":
		fmt.Println("I like ", favSport)
	case "tennis":
		fmt.Println("I like ", favSport)
	default:
		fmt.Println("I don't like ", favSport)
	}

}
