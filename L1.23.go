package main

//Удалить i-ый элемент из слайса.
import "fmt"

func removeElement(s []int, i int) []int {
	copy(s[i:], s[i+1:]) // Сдвигаем элементы после удаленного элемента на одну позицию влево
	s = s[:len(s)-1]     // Уменьшаем размер слайса на 1
	return s
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	index := 2

	fmt.Println(slice)

	newSlice := removeElement(slice, index)
	fmt.Println(newSlice)
}
