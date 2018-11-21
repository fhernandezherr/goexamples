package Exercises

import "fmt"

func main() {
	//exercise01()
	//exercise02()
	//exercise03()
	//exercise04()
	//exercise05()
	//exercise06()
	//exercise07()
	//exercise08()
	//exercise09()
	exercise10()
}

/*
Hands-on exercise #1

- Using a COMPOSITE LITERAL:
	- create an ARRAY which holds 5 VALUES of TYPE int
	- assign VALUES to each index position.
- Range over the array and print the values out.
- Using format printing print out the TYPE of the array

solution: https://play.golang.org/p/tD0SzV3hdf
*/
func exercise01() {
	a := [5]int{1, 2, 3, 4, 5}

	for k, v := range a {
		fmt.Println(k, v)
	}

	fmt.Printf("%T", a)
}

/*
Hands-on exercise #2

- Using a COMPOSITE LITERAL:
	- create a SLICE of TYPE int
	- assign 10 VALUES
- Range over the slice and print the values out.
- Using format printing
	- print out the TYPE of the slice

solution: https://play.golang.org/p/sAQeFB7DIs
*/
func exercise02() {
	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for k, v := range s {
		fmt.Println(k, v)
	}

	fmt.Printf("%T", s)

}

/*
Hands-on exercise #3

Using the code from the previous example, use SLICING to create the following new slices which are then printed:
[42 43 44 45 46]
[47 48 49 50 51]
[44 45 46 47 48]
[43 44 45 46 47]
solution: https://play.golang.org/p/SGfiULXzAB
*/
func exercise03() {
	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(s[:5])
	fmt.Println(s[5:])
	fmt.Println(s[2:7])
	fmt.Println(s[1:6])
}

/*
Hands-on exercise #4

Follow these steps:
	start with this slice
		x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	append to that slice this value
		52
	print out the slice
	in ONE STATEMENT append to that slice these values
		53
		54
		55
	print out the slice
	append to the slice this slice
		y := []int{56, 57, 58, 59, 60}
	print out the slice

solution: https://play.golang.org/p/QUYhtSBaDS
*/
func exercise04() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)

	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)

}

/*
Hands-on exercise #5

To DELETE from a slice, we use APPEND along with SLICING. For this hands-on exercise, follow these steps:
	start with this slice
		x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	use APPEND & SLICING to get these values here which you should ASSIGN to a variable “y” and then print:
		[42, 43, 44, 48, 49, 50, 51]

solution: https://play.golang.org/p/u8zpHLfgSE

*/
func exercise05() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	y := append(x[:3], x[6:]...)
	fmt.Println(y)
}

/*
Hands-on exercise #6
Create a slice to store the names of all of the states in the United States of America. What is the length of your slice? What is the capacity? Print out all of the values, along with their index position in the slice, without using the range clause. Here is a list of the states:

" Alabama", " Alaska", " Arizona", " Arkansas", " California", " Colorado", " Connecticut", " Delaware", " Florida", " Georgia", " Hawaii", " Idaho", " Illinois", " Indiana", " Iowa", " Kansas", " Kentucky", " Louisiana", " Maine", " Maryland", " Massachusetts", " Michigan", " Minnesota", " Mississippi", " Missouri", " Montana", " Nebraska", " Nevada", " New Hampshire", " New Jersey", " New Mexico", " New York", " North Carolina", " North Dakota", " Ohio", " Oklahoma", " Oregon", " Pennsylvania", " Rhode Island", " South Carolina", " South Dakota", " Tennessee", " Texas", " Utah", " Vermont", " Virginia", " Washington", " West Virginia", " Wisconsin", " Wyoming",

solution: https://play.golang.org/p/tRKQDQuQCE
*/
func exercise06() {
	states := make([]string, 50, 50)
	states = []string{
		"Alabama", "Alaska", " Arizona", " Arkansas", " California", " Colorado",
		"Connecticut", " Delaware", " Florida", " Georgia", " Hawaii", " Idaho",
		"Illinois", "Indiana", "Iowa", "Kansas", "Kentucky", "Louisiana", "Maine",
		"Maryland", "Massachusetts", "Michigan", "Minnesota", "Mississippi", "Missouri",
		"Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey", "New Mexico",
		"New York", "North Carolina", "North Dakota", "Ohio", "Oklahoma", "Oregon",
		"Pennsylvania", "Rhode Island", "South Carolina", "South Dakota", "Tennessee",
		"Texas", "Utah", "Vermont", "Virginia", "Washington", "West Virginia",
		"Wisconsin", "Wyoming"}

	fmt.Println(len(states))
	fmt.Println(cap(states))

	for i := 0; i < len(states); i++ {
		fmt.Println(i, states[i])
	}
}

/*
Hands-on exercise #7

Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:

	"James", "Bond", "Shaken, not stirred"
	"Miss", "Moneypenny", "Helloooooo, James."

Range over the records, then range over the data in each record.

solution: https://play.golang.org/p/FMM4c2PhZg
*/
func exercise07() {
	x := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helooooo, James."},
	}
	fmt.Println(x)

	for _, v := range x {
		fmt.Println(v)
		for _, d := range v {
			fmt.Println(d)
		}
	}
}

/*
Hands-on exercise #8
Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of TYPE []string which stores their favorite things. Store three records in your map. Print out all of the values, along with their index position in the slice.

	`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
	`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
	`no_dr`, `Being evil`, `Ice cream`, `Sunsets`

solution: https://play.golang.org/p/nTzSlRa9_A
*/
func exercise08() {
	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}

	for k, v := range m {
		fmt.Println("This is the record for ", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}

/*
Hands-on exercise #9

Using the code from the previous example, add a record to your map. Now print the map out using the “range” loop

solution: https://play.golang.org/p/_CkxAhRrDJ
video: 079
*/
func exercise09() {
	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}

	m["hdez_fran"] = []string{"good", "better", "the best"}

	for k, v := range m {
		fmt.Println("This is the record for ", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}

/*
Hands-on exercise #10

Using the code from the previous example, delete a record from your map. Now print the map out using the “range” loop

solution: https://play.golang.org/p/TYl5EbjoeC
*/
func exercise10() {
	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}

	m["hdez_fran"] = []string{"good", "better", "the best"}

	delete(m, "moneypenny_miss")

	for k, v := range m {
		fmt.Println("This is the record for ", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
