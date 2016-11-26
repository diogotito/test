package main

import "fmt"

func main() {
	words := [...]string{"Gate", "Comet", "Pizza"}

	for _, word := range words {

		// Check whether word has the character 'e'
		found := false
		for _, char := range word {
			if char == 'e' {
				found = true
				break
			}
		}
		if found {
			fmt.Printf("Found one! %c\n", word[0])
		}
	}
}
