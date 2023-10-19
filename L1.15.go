package main

/*К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.
var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}
func main() {
  someFunc()
}*/
import (
	"fmt"
)

var justString string

func createHugeString(size int) string {
	// Возвращает строку заданного размера
	// Здесь приведен простой пример для демонстрации
	return "This is a huge string"
}

func someFunc() {
	v := createHugeString(1 << 10)
	if len(v) >= 100 {
		justString = v[:100]
	} else {
		justString = v
	}
}

func main() {
	someFunc()
	fmt.Println(justString)
}
