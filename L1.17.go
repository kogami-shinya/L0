package main

//Реализовать бинарный поиск встроенными методами языка.
import (
	"fmt"
	"sort"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

	index := sort.Search(len(numbers), func(i int) bool {
		return numbers[i] >= 6
	})

	if index < len(numbers) && numbers[index] == 6 {
		fmt.Println("Число 6 найдено на позиции", index)
	} else {
		fmt.Println("Число 6 не найдено")
	}
}
