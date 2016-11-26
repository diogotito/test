package main

import "fmt"

func main() {
	fruits := []string{"banana", "kiwi", "apple", "pineapple"}
	fmt.Printf("%v\n", fruits)

	for _, num := range fruits {
		fmt.Println(num)
	}
}
