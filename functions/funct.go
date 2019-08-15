package main

import "fmt"

func main() { 
	fmt.Println("Hello function")
	helloFunction()
	paramFunction("John", 45)
}

func helloFunction() {
	fmt.Println("Hello function called....")
}

func paramFunction(name string, age int) {
	fmt.Printf("Hello %s age %d\n", name, age)
}