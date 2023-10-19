package main

//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов с использованием конкурентных вычислений.
import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	var sum int
	var wg sync.WaitGroup
	var mux sync.Mutex

	for _, num := range numbers {
		wg.Add(1)
		go func(value int) {
			defer wg.Done()
			square := value * value
			// Защищаем доступ к переменной sum с помощью мьютекса
			// для предотвращения гонки данных
			mux.Lock()
			sum += square
			mux.Unlock()
		}(num)
	}

	wg.Wait()

	fmt.Println("Сумма квадратов чисел:", sum)
}
