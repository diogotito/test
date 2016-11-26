package main

import "fmt"

func main() {
	fmt.Println("--------------------------------------------------")
	fmt.Println("Cool new features!")
	fmt.Println("--------------------------------------------------")

	recursionFTW(12)
}

func recursionFTW(numero int) {
	if numero != 4 && numero != 2 && numero != 10 {
		fmt.Println("i =", numero)
	}

	if numero == 0 {
		return
	}

	recursionFTW(numero - 1)
}
