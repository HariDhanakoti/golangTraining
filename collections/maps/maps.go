package main

import "fmt"

func main() {
	m := map[int]string{1: "One", 2: "Two", 3: "Three"}
	fmt.Println(m)

	fmt.Println(m[2])

	m[3] = "Thrity three"
	fmt.Println(m[3])

	fmt.Println(m)

	delete(m, 3)

	fmt.Println(m)
}