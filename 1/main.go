package main

import "fmt"

// Human структура
type Human struct {
}

// Action структура
type Action struct {
	Human //встроили Human в Action
}

func main() {
	act := Action{Human{}} // инит Action
	act.Say("Hi!")         // Вызываем метод
}

// Say метод для распечатки сообщений структуры Human
func (h *Human) Say(msg string) {
	fmt.Println(msg)
}
