package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func divide(numerator int, denominator int) (int, int, error) {
	if denominator == 0 {
		// By convetion error is always the last return value.
		return 0, 0, errors.New("Cannot divide by zero")
	}
	quotient := numerator / denominator
	remainder := numerator % denominator
	// Multiple return values, when successful error is nil.
	return quotient, remainder, nil
}

// optional parameters, just avoided it.
type MyFunctOps struct {
	FirsName string
	LastName string
	Age      int
}

func MyFunct(opts MyFunctOps) error {
	// do something
	return nil
}

// variadic parameters
func addTo(n int, numbers ...int) int {
	total := n
	for _, number := range numbers {
		total += number
	}
	return total
}

func add(n int, m int) int {
	return n + m
}

func sub(n int, m int) int {
	return n - m
}

func mul(n int, m int) int {
	return n * m
}

func div(n int, m int) int {
	return n / m
}

func calculator() {
	fmt.Println("\n ------ Calculator ------")

	// declaring types for functions.
	type opFuncType func(int, int) int

	var opMap = map[string]opFuncType{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}

	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "plus", "3"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Println("Invalid expression:", expression)
			continue
		}

		p1, err := strconv.Atoi(string(expression[0]))
		if err != nil {
			fmt.Println(err)
			continue
		}

		op := string(expression[1])
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Println("Invalid operator:", op)
			continue
		}
		p2, err := strconv.Atoi(string(expression[2]))
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := opFunc(p1, p2)

		fmt.Println(expression, "=", result)
	}
}

func anonymousFunction() {
	fmt.Println("\n ------ Anonymous Function ------")
	// Anonymous functions are functions without a name.
	// They are useful when you want to define a function inline without having to name it.
	// They are also useful when you want to define a function that you will only use once.
	// Anonymous functions are often used as arguments to other functions.
	// Anonymous functions are also called closures because they can "close" over the variables in the function that they are defined in.
	// The syntax for defining an anonymous function is similar to the syntax for defining a named function.
	// The only difference is that you don't give it a name.
	// The following example defines an anonymous function and assigns it to the variable add.
	// The function takes two int parameters and returns an int.
	// The function adds the two parameters and returns the result.
	add := func(a int, b int) int {
		return a + b
	}
	// The add variable is now a function that can be called like any other function.
	// The following code calls the add function and prints the result.
	fmt.Println(add(1, 1))

	for i := 0; i < 5; i++ {
		// The following code defines an anonymous function inside a for loop.
		// The function prints the value of the i variable.
		// The function is called immediately after it is defined.
		func(j int) {
			fmt.Println(j)
		}(i)
	}

}

func functionParameters() {
	// Functions as parameters.
	fmt.Println("\n ------ Function Parameters ------")
	type Person struct {
		firstName string
		lastName  string
		age       int
	}

	people := []Person{
		{"John", "Doe", 25},
		{"Jane", "Doe", 23},
		{"Jack", "Doe", 21},
	}
	fmt.Println("Original ->", people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].age < people[j].age
	})

	fmt.Println("Ascending order ->", people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].age > people[j].age
	})

	fmt.Println("Descending order -> ", people)
}

func catCommand() {
	// cat command
	fmt.Println("\n ------ Cat Command ------")
	filePath := "test.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	data := make([]byte, 2048)

	for {
		count, err := file.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
	}
}
