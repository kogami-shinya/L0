package main

//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{5, 3, 8, 2, 1, 6, 4, 7}

	sort.Ints(numbers)

	fmt.Println(numbers) // Вывод: [1 2 3 4 5 6 7 8]
}
