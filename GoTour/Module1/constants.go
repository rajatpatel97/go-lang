package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "Hakuna Matata"
	fmt.Println("hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
