package main

import "fmt"

const c int = 64

func explicitTypeConversion() {
	fmt.Println("\n --- Explicit Type conversions ---\n ")
	var x int = 10
	var y float64 = 10.2
	var res float64 = float64(x) + y
	fmt.Printf("Result %v \n", res)
	fmt.Println(c)

}

func sliceLiterals() {
	fmt.Println("\n --- Slice Literals ---\n ")
	var s []int
	s = append(s, 10, 11)
	fmt.Printf("Value: %v Length: %v \n", s, len(s))

	s2 := []int{20, 30, 40}
	fmt.Printf("Concatenating two slices \n s : %v \n s2 : %v \n", s, s2)
	s = append(s, s2...)
	fmt.Printf("Value: %v \n Length: %v \n Capacity %v \n", s, len(s), cap(s))

}

func makeSlice() {
	fmt.Println("\n --- make Slices ---\n ")
	x := make([]int, 5)
	x = append(x, 10)
	fmt.Println("Value of X:", x)

	y := make([]int, 0, 10) // size, capacity
	y = append(y, 10, 2, 3)
	fmt.Println("Value of Y:", y)

}

func slicingSlices() {

	fmt.Println("\n --- Slicing Slices ---")
	y := make([]int, 0, 10) // size, capacity
	y = append(y, 10, 2, 3, 4)
	fmt.Println("Value of Y:", y)
	w := y[:2]
	fmt.Println("Slicing the slice : ", w) // Slice shares memory between original and slice.
	w[0] = 33
	fmt.Println("Original after mutating slice: ", y)
	fmt.Println("Using copy on slice")
	z := make([]int, 2)
	copy(z, y)
	fmt.Println("New slice z :", z)
	z[0] = 3
	fmt.Println("Modified z's 0th element :", z)
	fmt.Println("Original y :", y)

}

func stringRunes() {

	fmt.Println("\n ------ Strings and runes ------")
	s := "Hello World!"
	fmt.Printf("Slicing %v\n", s)
	s2 := s[:5]
	fmt.Printf("Sliced %v\n", s2)
	fmt.Println("\n Converting int string")
	var x int = 65
	var y = string(x)
	fmt.Printf("Original -> Value: %v Type: %T \n", x, x)
	fmt.Printf("Converted -> Value: %v Type: %T \n", y, y)
	fmt.Println("\nConverting strings to slices")
	var s3 string = "Hello 😊"
	fmt.Printf("Original %v\n", s3)
	var bs []byte = []byte(s3)
	var rs []rune = []rune(s3)
	fmt.Printf("Byte %v\n", bs)
	fmt.Printf("Rune %v\n", rs)
}

func mappingMaps() {
	fmt.Println("\n ------ Maps ------")
	// map literal
	totalWins := map[string]int{"Jose": 2, "Anna": 1}
	fmt.Printf("Literal Map:  %v\n",totalWins)
	fmt.Println(totalWins["Jose"])
	fmt.Println(totalWins["Anna"])
	// No default size
	fmt.Println("\nReading and writting\n")
	ages := map[string]int{}
	// Adding new keys:
	ages["Jose"] = 26
	ages["Anna"] = 28
	ages["Rossi"] = 8
	// Reading keys
	fmt.Println("Jose's age:", ages["Jose"])
	fmt.Println("Anna's age:", ages["Anna"])
	fmt.Println("Rossi's age:", ages["Rossi"])
	// Testing when no key exists
	fmt.Println("No key's age:", ages["Nokey"])
	fmt.Println("\nThe comma ok Idiom\n")
	// Used to know if the 0 returned it's actually a value and not a non-existing key
	a := map[string]int{"Hello": 2, "World":0}
	v, ok := a["Hello"]
	fmt.Println(v, ok)
	v, ok = a["hello"]
	fmt.Println(v, ok)
	v, ok = a["World"]
	fmt.Println(v, ok)
	fmt.Println("\nDeleting from maps\n")
	fmt.Println("Original Map:", a)
	delete(a, "Hello") // doesn't return a value
	fmt.Println("After deleting key:", a)
	fmt.Println("\nUsing maps as sets.\n")
	set := map[int]string{}
	// Add some elements to the set
	set[1] = ""
	set[2] = ""
	set[3] = ""
	// Check if an element is in the set
	if _, ok := set[2]; ok {
		fmt.Println(2, "is in the set")
	} else {
		fmt.Println(2, "is not in the set")
	}

	// Print the set
	fmt.Println("Set:", set)
}

func structuredStruct() {
	fmt.Println("\n ------ Structs ------")
	type person struct {
		name string
		age int
		pet string
	}
	// declaring and assigning
	var jose person
	jose.name = "jose"
	jose.age = 28
	jose.pet = "Rossi"
	fmt.Println("Jose", jose)
	// creating a literal
	anna := person{"Anna", 28, "Jose"} // use for small strcuts
	fmt.Println("Anna", anna)
	// creating literal like a map
	joao := person{name: "joao", pet: "farrusco"} // preffered
	fmt.Println("Joao", joao)
	// Anonymous struct
	pet := struct{name string
				  kind string
				}{
				  name: "Rossi",
				  kind: "Dog-Labrador",
				  }
	fmt.Println("Anonymous struct", pet)


}

func shadowingShadows() {
	fmt.Println("\n ------ Shadowing ------")
	x := 10
	fmt.Println("Outside block", x)
	if x > 5{
		x:=5
		fmt.Println("Inside if block", x)
	}
	fmt.Println("Outside block", x)

}

func loopingLoops() {
	fmt.Println("\n ------ Loops ------")
	// Loop through an array using a for loop
	fmt.Println("For loop while style:")
	i := 1
	for i < 10 {
		fmt.Println("Loop")
		i = i * 2
	}

	// Loop through a slice using a for loop with a range clause
	fmt.Println("\nFor loop with range clause:")
	slice := []int{6, 7, 8, 9, 10}
	for _, x := range slice {
		fmt.Printf("Value: %d\n", x)
	}

	// Loop indefinitely using a for loop with no condition
	fmt.Println("\nFor loop with no condition:")
	for {
		fmt.Println("Infinite loop!")
		break
	}

	// Loop through a slice using a for loop with a condition
	fmt.Println("\nFor loop C style:")
	for i := 0; i < len(slice); i++ {
		fmt.Printf("Index: %d Value: %d\n", i, slice[i])
	}
}