package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool) // Канал для сигнала остановки

	go func(stop chan bool) {
		// Выполнение работы в горутине
		for {
			// Работа горутины

			// Проверка сигнала остановки
			select {
			default:
				// Продолжение работы

			case <-stop:
				// Получен сигнал остановки
				fmt.Println("Received stop signal")
				return
			}
		}
	}(stop)

	// Пауза в основном потоке
	time.Sleep(2 * time.Second)

	// Отправка сигнала остановки
	stop <- true

	// Ожидание завершения горутины
	time.Sleep(1 * time.Second)

	fmt.Println("Program completed")
}
