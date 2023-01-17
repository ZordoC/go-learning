package main

import "fmt"

func tech_nana() {

	fmt.Println("Your personalized phone book")

	mySlice := []string{"Jose", "Anna", "Good"}

	for index, name := range mySlice {

		fmt.Printf("no %v,  name %v \n", index, name)

		if name == "Anna" {
			fmt.Printf("It's the POKEMON! \n")

		} else {
			fmt.Printf("Did not find wanted Pokemon! \n")

		}

	}

	fmt.Println(`Testing raw string literal "Works"? 
Is this a new line`)

}
