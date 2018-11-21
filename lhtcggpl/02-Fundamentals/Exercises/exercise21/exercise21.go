/*
Hands-on exercise #1

Write a program that prints a number in decimal, binary, and hex
*/
package main

import "fmt"

func main() {
	n := 42

	fmt.Printf("%d\t%b\t%#x", n, n, n)
}
