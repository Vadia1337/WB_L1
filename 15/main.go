package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now() // пример плохой реализации, время выполнения такого кода 60-70 ms
	for i := 10000; i > 0; i-- {
		someFunc1()
	}
	fmt.Println(time.Since(start))

	start = time.Now() // время выполнения кода 48 - 52 ms
	for i := 10000; i > 0; i-- {
		someFunc()
	}
	fmt.Println(time.Since(start))
}

//К каким негативным последствиям может привести данный фрагмент кода,
//и как это исправить? Приведите корректный пример реализации.

//Первое что попадается на глаза - объявлена глобальная переменная. В данном случае использование глоб. переменной
// не оправдано, по причине отсутствия необходимости в доступе к переменной из вне.
// судя по названию функции createHugeString возвращаться будет строка, а далее хотим обрезать такую строку
// до 100 символов проблема заключается в том, что длинная строка занимает много памяти, операции с большими строками
//мало того, что медленные, но и при обрезке срез ссылается на строку, что излишне использует память, ссылка на область
// памяти не дает сборщику очистить данную область.

var justString string

func someFunc1() {
	v := createHugeString1(1 << 10)
	justString = v[:100] // операция потребляет излишнее время
}

func createHugeString1(lenOfString int) string {
	str := make([]byte, lenOfString)
	j := 0
	for i := lenOfString; i > 0; i-- {
		j += copy(str[j:], "L")
	}
	return string(str)
}

func someFunc() {
	createHugeString(1 << 10)
}

func createHugeString(lenOfString int) string {
	str := make([]byte, lenOfString)
	j := 0
	for i := lenOfString; i > 0; i-- {
		j += copy(str[j:], "L")
	}

	var strSlice []byte //создаем слайс для подстроки

	copy(strSlice, str[:100]) // копируем значение в слайс

	return string(strSlice)
}
