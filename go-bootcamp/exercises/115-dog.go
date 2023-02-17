package main

/*
Create a dog package. The dog package should have an exported func “Years” which takes human years and turns them into dog years (1 human year = 7 dog years). Document  your code with comments. Use this code in func main.
	run your program and make sure it works
	run a local server with godoc and look at your documentation.
*/

import (
	"fmt"
)

func main() {
	fmt.Println(dog.DogYears(26))
	fmt.Println(dog.HumanYears(26))
}