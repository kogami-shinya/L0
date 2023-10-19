package main

//Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
import (
	"fmt"
	"time"
)

func main() {
	// Время выполнения программы (в секундах)
	duration := 5

	// Канал для передачи данных
	dataChan := make(chan int)

	// Запуск отправителя данных
	go sender(dataChan)

	// Запуск получателя данных
	go receiver(dataChan)

	// Ожидание N секунд
	time.Sleep(time.Duration(duration) * time.Second)

	// Завершение программы
	fmt.Println("Program completed")
}

// Функция отправки данных в канал
func sender(dataChan chan<- int) {
	for i := 1; ; i++ {
		// Отправка значения в канал
		dataChan <- i
		time.Sleep(1 * time.Second) // Задержка между отправками
	}
}

// Функция чтения данных из канала
func receiver(dataChan <-chan int) {
	for data := range dataChan {
		// Получение значения из канала и обработка
		fmt.Println("Received:", data)
	}
}
