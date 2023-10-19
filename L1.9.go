package main

//Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.
import "fmt"

func multiplyNumbers(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * 2
	}
	close(out)
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// Создаем каналы
	input := make(chan int)
	output := make(chan int)

	// Запускаем горутину для умножения чисел
	go multiplyNumbers(input, output)

	// Записываем числа из массива в канал input
	go func() {
		for _, x := range numbers {
			input <- x
		}
		close(input)
	}()

	// Выводим результаты из канала output в stdout
	for result := range output {
		fmt.Println(result)
	}
}
