package main

//Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна. Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.
import (
	"fmt"
	"math"
)

func groupTemperatures(temperatures []float64) map[int][]float64 {
	groups := make(map[int][]float64)

	for _, temp := range temperatures {
		group := int(math.Floor(temp/10.0) * 10)
		groups[group] = append(groups[group], temp)
	}

	return groups
}

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := groupTemperatures(temperatures)

	for group, temps := range groups {
		fmt.Printf("%d: %v\n", group, temps)
	}
}
