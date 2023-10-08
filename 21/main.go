package main

import "fmt"

type Dog struct{}

func (d *Dog) SitDown(command string) {
	if command == "сидеть" {
		fmt.Println("*Собака села*")
	}
}

type Cat struct{}

func (c *Cat) RunToOwner(command string, food string) {
	if command == "кис-кис" && food == "вискас" {
		fmt.Println("*Кот бежит со всех ног*")
	}
}

type PetAdapter interface {
	PetAction()
}

type DogAdapter struct {
	*Dog
}

func (adapter *DogAdapter) PetAction() {
	adapter.SitDown("сидеть") // собаке нужно просто сказать команду
}

// конструктор адаптера для собаки
func NewDogAdapter(dog *Dog) PetAdapter {
	return &DogAdapter{dog}
}

type CatAdapter struct {
	*Cat
}

func (adapter *CatAdapter) PetAction() {
	// адаптер автоматически зовет кошку isCall = true
	adapter.RunToOwner("кис-кис", "вискас") //кота надо позвать и иметь в наличии лакомство - иначе не придет
}

// конструктор адаптера для кота
func NewCatAdapter(cat *Cat) PetAdapter {
	return &CatAdapter{cat}
}

func main() {
	pets := [2]PetAdapter{ // через интерфейс адаптера по очереди вызываем действия питомцев
		NewDogAdapter(&Dog{}),
		NewCatAdapter(&Cat{}),
	}

	for _, pet := range pets {
		pet.PetAction()
	}
	// скрестили действия питомцев, можем использовать 2 класса как что-то однотипное
}
