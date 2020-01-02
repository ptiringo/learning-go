package main

import "fmt"

func main() {
	const days = 28
	const distance = 56000000 // km

	const hours = days * 24

	fmt.Printf("%v km/時", distance / hours)
}
