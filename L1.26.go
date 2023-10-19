package main

//Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой. Например: abcd — true abCdefAaf — false aabcd — false
import (
	"fmt"
	"strings"
)

func isUnique(str string) bool {
	charCount := make(map[rune]int)

	for _, char := range strings.ToLower(str) {
		if charCount[char] > 0 {
			return false
		}
		charCount[char]++
	}

	return true
}

func main() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"

	fmt.Println(isUnique(str1)) // Вывод: true
	fmt.Println(isUnique(str2)) // Вывод: false
	fmt.Println(isUnique(str3)) // Вывод: false
}
