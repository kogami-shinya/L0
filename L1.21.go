package main

//Реализовать паттерн «адаптер» на любом примере.
import "fmt"

type AnimalSystem interface {
	GetAnimals() []Animal
	AddNewAnimal(newAnimal Animal)
}

type AnimalCatalog interface {
	GetAllAnimals() []Animal
	AddAnimal(animal Animal)
}

// Определяем структуру животного
type Animal struct {
	Name string
	Age  int
}

// Реализуем структуру системы работы с каталогом животных
type AnimalCatalogImpl struct {
	Animals []Animal
}

// Реализуем метод GetAllAnimals для системы работы с каталогом животных
func (ac *AnimalCatalogImpl) GetAllAnimals() []Animal {
	return ac.Animals
}

// Реализуем метод AddAnimal для системы работы с каталогом животных
func (ac *AnimalCatalogImpl) AddAnimal(animal Animal) {
	ac.Animals = append(ac.Animals, animal)
}

// Создаем адаптер для системы работы с каталогом животных
type AnimalSystemAdapter struct {
	Catalog AnimalCatalog
}

// Реализуем метод GetAnimals для адаптера
func (asa *AnimalSystemAdapter) GetAnimals() []Animal {
	return asa.Catalog.GetAllAnimals()
}

// Реализуем метод AddNewAnimal для адаптера
func (asa *AnimalSystemAdapter) AddNewAnimal(newAnimal Animal) {
	asa.Catalog.AddAnimal(newAnimal)
}

func main() {
	// Создаем экземпляр системы работы с каталогом животных
	catalog := &AnimalCatalogImpl{
		Animals: []Animal{
			{Name: "Кошка", Age: 2},
			{Name: "Собака", Age: 4},
		},
	}

	// Создаем адаптер
	adapter := &AnimalSystemAdapter{
		Catalog: catalog,
	}

	// Используем адаптер для получения и добавления животных
	animals := adapter.GetAnimals()
	fmt.Println("Животные в системе:", animals)

	newAnimal := Animal{Name: "Хомяк", Age: 1}
	adapter.AddNewAnimal(newAnimal)
	animals = adapter.GetAnimals()
	fmt.Println("Животные в системе после добавления:", animals)
}
