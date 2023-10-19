package main

//Написать программу, которая конкурентно рассчитает значения квадратов чисел из массива (2,4,6,8,10) и выведет их квадраты в stdout.
import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	result := make([]int, len(numbers))
	var wg sync.WaitGroup

	for i, num := range numbers {
		wg.Add(1)
		go func(index, value int) {
			defer wg.Done()
			result[index] = value * value
		}(i, num)
	}

	wg.Wait()

	for _, res := range result {
		fmt.Println(res)
	}
}
