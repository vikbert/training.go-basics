package main

import "fmt"

func main() {
	fmt.Println("Pointers:")

	var (
		n1 = 88
		n2 = 2
	)

	miniSort(&n1, &n2)
	fmt.Println(n1, n2) // Ausgabe: 2 88
}

func miniSort(a, b *int) {
	if *a > *b {
		*a, *b = *b, *a
	}
}
