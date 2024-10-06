package main // package declaration

import "fmt"

func main() {
	fmt.Printf("Hello World")

	const earthsGravity = 9.80665
	const funFact = "Hummingbirds' wings can beat up to 200 times a second."
	fmt.Println("Did you know?")
	fmt.Println(funFact)
	// Here's where we print out the gravity:
	fmt.Println(earthsGravity)
}
