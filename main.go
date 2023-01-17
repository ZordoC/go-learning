package main

import "fmt"

func main() {

	fmt.Println("Main file.")
	explicitTypeConversion()
	sliceLiterals()
	makeSlice()
	slicingSlices()
	stringRunes()
	mappingMaps()
	structuredStruct()
	shadowingShadows()
	loopingLoops()
	fmt.Println(" ----- Functions -----")
	// The _ is a blank identifier, it is used to ignore the return value, the compiler will not complain.
	q, r, _ := divide(10, 2)
	fmt.Println(q, r)
	fmt.Println(divide(10, 2))
	fmt.Println(divide(100, 3))
	fmt.Println(divide(100, 0))
	fmt.Println(addTo(100, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(addTo(10, 3, 5, 6))
	calculator()
	anonymousFunction()
	functionParameters()
}
