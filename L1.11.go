package main

//Реализовать пересечение двух неупорядоченных множеств.
import "fmt"

// Функция для вычисления пересечения двух множеств
func intersection(set1, set2 []int) []int {
	// Создаем map для хранения элементов первого множества
	set1Map := make(map[int]bool)

	// Добавляем элементы первого множества в map
	for _, num := range set1 {
		set1Map[num] = true
	}

	// Инициализируем пустой slice для хранения пересечения
	intersect := []int{}

	// Проверяем каждый элемент второго множества на принадлежность к первому множеству
	for _, num := range set2 {
		if set1Map[num] {
			intersect = append(intersect, num)
		}
	}

	return intersect
}

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{4, 5, 6, 7, 8}

	result := intersection(set1, set2)
	fmt.Println(result)
}
