/*
	Hands-on exercise #6

	Using iota, create 4 constants for the NEXT 4 years. Print the constant values.
*/

package main

import "fmt"

const (
	year1 = 2018 + iota
	year2
	year3
	year4
)

func main() {

	fmt.Println(year1, year2, year3, year4)
}
