package main

//Реализовать постоянную запись данных в канал (главный поток). Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout. Необходима возможность выбора количества воркеров при старте. Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	// Количество воркеров
	numWorkers := 5

	// Создание каналов для взаимодействия
	dataChan := make(chan string)   // Канал данных
	doneChan := make(chan struct{}) // Канал для сигнала завершения

	// Создание WaitGroup для отслеживания завершения работы воркеров
	var wg sync.WaitGroup

	// Запуск воркеров
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, dataChan, &wg)
	}

	// Запись данных в канал (главный поток)
	go func() {
		defer close(dataChan)

		for {
			// Генерация произвольных данных для записи в канал
			data := generateData()

			// Запись данных в канал
			select {
			case dataChan <- data:
				// Успешная запись данных
			case <-doneChan:
				// Получен сигнал завершения работы
				return
			}

			time.Sleep(1 * time.Second) // Задержка между записями
		}
	}()

	// Обработка сигнала Ctrl+C для завершения работы программы и всех воркеров
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-signalChan
		fmt.Println("Received interrupt signal")
		close(doneChan) // Отправляем сигнал завершения воркерам
	}()

	// Ожидание завершения работы всех воркеров
	wg.Wait()

	fmt.Println("All workers have finished. Program has terminated.")
}

// Генерация произвольных данных
func generateData() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// Воркеры для чтения данных из канала
func worker(id int, dataChan <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for data := range dataChan {
		fmt.Printf("Worker ID: %d, Data: %s\n", id, data)
		// Некоторая обработка данных

		// Здесь вы можете выполнять любую другую работу с данными из канала
		// в соответствии с логикой вашего приложения
	}
}
