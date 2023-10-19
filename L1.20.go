package main

//Разработать программу, которая переворачивает слова в строке. Пример: «snow dog sun — sun dog snow».
import (
	"fmt"
	"strings"
)

func reverseWords(input string) string {
	words := strings.Fields(input)

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}

func main() {
	input := "snow dog sun"
	reversed := reverseWords(input)
	fmt.Println(reversed) // Вывод: "sun dog snow"
}
