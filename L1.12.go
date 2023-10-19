package main

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
import "fmt"

func createSet(strings []string) map[string]bool {
	set := make(map[string]bool)

	for _, s := range strings {
		set[s] = true
	}

	return set
}

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	set := createSet(strings)

	fmt.Println("Множество:")
	for key := range set {
		fmt.Println(key)
	}
}
