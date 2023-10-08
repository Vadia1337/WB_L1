package main

import (
	"fmt"
	"math/big"
)

func main() {
	var a, b big.Int

	// 2 в 20 может поместиться и в uint64, но под условие задачи >2 в 20 могут попадать числа более 18 квинтилионов,
	// при таком раскладе даже на этапе пристваивания значения появиться ошибка overflows uint64
	// решением станет использование пакета math/big

	a.Exp(big.NewInt(10), big.NewInt(33), nil) // nil - модуль, получим квинтиллиард
	b.Exp(big.NewInt(10), big.NewInt(33), nil) // nil - модуль, получим квинтиллиард

	fmt.Println("a =", a.String(), "b =", b.String()) // взлянем на них

	fmt.Println("сложение: ", a.Add(&a, &b))

	fmt.Println("вычитание:", a.Sub(&a, &b))

	fmt.Println("умножение:", a.Mul(&a, &b))

	fmt.Println("деление:  ", a.Div(&a, &b))
}
