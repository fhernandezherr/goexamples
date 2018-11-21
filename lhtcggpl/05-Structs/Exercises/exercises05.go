package main

import "fmt"

func main() {
	//exercise01()
	//exercise02()
	//exercise03()
	exercise04()
}

/*
Hands-on exercise #1

Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
	first name
	last name
	favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.

solution: https://play.golang.org/p/ouRHmH_POg
*/
func exercise01() {

	type person struct {
		first   string
		last    string
		flavors []string
	}

	p := []person{
		{
			first: "James",
			last:  "Bond",
			flavors: []string{
				"Orange",
				"Banana",
				"Vanilla",
			},
		},
		{
			first: "Money",
			last:  "Penny",
			flavors: []string{
				"Lemon",
				"Apple",
			},
		},
	}

	for k, v := range p {
		fmt.Println("Favourite flavors of person ", k, " ", v.first, " ", v.last)
		for i, f := range v.flavors {
			fmt.Println("\tFlavor ", i, ": ", f)
		}
	}

}

/*
Hands-on exercise #2

Take the code from the previous exercise, then store the values of type person in a map with the key of last name. Access each value in the map. Print out the values, ranging over the slice.

solution: https://play.golang.org/p/RvvLyAMvGo

*/
func exercise02() {

	type person struct {
		first   string
		last    string
		flavors []string
	}

	p := map[string]person{
		"Bond": person{
			first: "James",
			last:  "Bond",
			flavors: []string{
				"Orange",
				"Banana",
				"Vanilla",
			},
		},
		"Penny": person{
			first: "Money",
			last:  "Penny",
			flavors: []string{
				"Lemon",
				"Apple",
			},
		},
	}

	for k, v := range p {
		fmt.Println("Favourite flavors of person ", k, " ", v.first, " ", v.last)
		for i, f := range v.flavors {
			fmt.Println("\tFlavor ", i, ": ", f)
		}
	}
}

/*
Hands-on exercise #3

Create a new type: vehicle.
	The underlying type is a struct.
	The fields:
		doors
		color
	Create two new types: truck & sedan.
		The underlying type of each of these new types is a struct.
		Embed the “vehicle” type in both truck & sedan.
		Give truck the field “fourWheel” which will be set to bool.
		Give sedan the field “luxury” which will be set to bool. solution
	Using the vehicle, truck, and sedan structs:
		using a composite literal, create a value of type truck and assign values to the fields;
		using a composite literal, create a value of type sedan and assign values to the fields.
	Print out each of these values.
	Print out a single field from each of these values.

solution: https://play.golang.org/p/PrTtTv_vVO
*/
func exercise03() {

	type vehicle struct {
		doors int
		color string
	}

	type truck struct {
		vehicle
		fourWheel bool
	}

	type sedan struct {
		vehicle
		luxury bool
	}

	// composite literal
	t := truck{
		vehicle: vehicle{
			doors: 4,
			color: "blue",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors: 5,
			color: "black",
		},
		luxury: false,
	}

	fmt.Println(t)
	fmt.Println(s)

	fmt.Println(t.fourWheel)
	fmt.Println(s.luxury)
}

/*
Create and use an anonymous struct.

solution: https://play.golang.org/p/otBHFs-lPp
*/
func exercise04() {

	example := struct{
		color string
		age int
	}{
		color: "blue",
		age: 12,
	}

	fmt.Println(example)

}
