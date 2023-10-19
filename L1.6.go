package main

//Реализовать все возможные способы остановки выполнения горутины.
import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool) // Канал для сигнала остановки

	go func() {
		// Выполнение работы в горутине
		for {
			select {
			default:
				// Работа горутины

			case <-stop:
				// Получен сигнал остановки
				fmt.Println("Received stop signal")
				return
			}
		}
	}()

	// Пауза в основном потоке
	time.Sleep(2 * time.Second)

	// Отправка сигнала остановки
	stop <- true

	// Ожидание завершения горутины
	time.Sleep(1 * time.Second)

	fmt.Println("Program completed")
}
