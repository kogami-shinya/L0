package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		// Выполнение работы в горутине
		for {
			select {
			default:
				// Работа горутины

			case <-ctx.Done():
				// Получен сигнал остановки
				fmt.Println("Received stop signal")
				return
			}
		}
	}()

	// Пауза в основном потоке
	time.Sleep(2 * time.Second)

	// Отправка сигнала остановки
	cancel()

	// Ожидание завершения горутины
	time.Sleep(1 * time.Second)

	fmt.Println("Program completed")
}
