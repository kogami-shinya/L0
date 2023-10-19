package main

import (
	"fmt"
	"time"
)

func main() {
	interrupt := make(chan struct{}) // Канал для прерывания работы

	go func() {
		// Выполнение работы в горутине
		for {
			select {
			default:
				// Работа горутины

			case <-interrupt:
				// Получен сигнал прерывания
				fmt.Println("Received interrupt signal")
				return
			}
		}
	}()

	// Пауза в основном потоке
	time.Sleep(2 * time.Second)

	// Отправка сигнала прерывания
	close(interrupt)

	// Ожидание завершения горутины
	time.Sleep(1 * time.Second)

	fmt.Println("Program completed")
}
