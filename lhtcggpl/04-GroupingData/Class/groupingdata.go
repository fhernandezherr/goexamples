package Class

import "fmt"

func main() {
	//array()
	//slice()
	//sliceForRange()
	//slicingSlice()
	//appendToSlice()
	//deletingFromSlice()
	//sliceMake()
	//multiDimensionalSlice()
	//mapIntroduction()
	//mapAddElementAndRange()
	mapDelete()
}

func array() {
	var x [5]int
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x))
}

func slice() {
	x := []int{4, 5, 6, 8, 42}
	fmt.Println(x)
}

func sliceForRange() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(len(x))
	fmt.Println(x)
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])

	for i, v := range x {
		fmt.Println(i, v)
	}
}

func slicingSlice() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	fmt.Println(x[:])
	fmt.Println(x[1:])
	fmt.Println(x[1:3])

	for i, v := range x {
		fmt.Println(i, v)
	}

	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}
}

// https://golang.org/doc/effective_go.html#append
// https://godoc.org/builtin#append
func appendToSlice() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)

	y := []int{234, 456, 678, 987}
	x = append(x, y...)
	fmt.Println(x)
}

func deletingFromSlice() {
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)

	y := []int{234, 456, 678, 987}
	x = append(x, y...)
	fmt.Println(x)

	// This is the way to delete from a slice
	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}

func sliceMake() {
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x[0] = 42
	x[9] = 999
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3423)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3424)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	x = append(x, 3425)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}

func multiDimensionalSlice() {
	jb := []string{"James", "Bond", "Chocolate", "Martini"}
	fmt.Println(jb)
	mp := []string{"Miss", "Monypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)
}

// https://golang.org/ref/spec#Map_types
// https://golang.org/doc/effective_go.html#maps
//
// MapType     = "map" "[" KeyType "]" ElementType .
// KeyType     = Type .
func mapIntroduction() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)

	fmt.Println(m["James"])

	fmt.Println(m["Barnabas"])

	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("THIS IS THE IF PRINT", v)
	}
}

func mapAddElementAndRange() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)

	fmt.Println(m["James"])

	fmt.Println(m["Barnabas"])

	v, ok := m["Barnabas"]
	fmt.Println(v)
	fmt.Println(ok)

	m["todd"] = 33

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("THIS IS THE IF PRINT", v)
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	xi := []int{4, 5, 7, 8, 9, 42}

	for i, v := range xi {
		fmt.Println(i, v)
	}
}

func mapDelete() {
	m := map[string]int{
		"James":           32,
		"Miss Moneypenny": 27,
	}
	fmt.Println(m)

	delete(m, "James")
	fmt.Println(m)

	delete(m, "Ian Fleming")
	fmt.Println(m)

	if v, ok := m["Miss Moneypenny"]; ok {
		fmt.Println("value:", v)
		delete(m, "Miss Moneypenny")
	}

	fmt.Println(m)
}
