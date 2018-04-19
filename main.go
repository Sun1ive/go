package main

import (
	"fmt"
)

// func main() {
// 	var a = 3
// 	fmt.Println("my first var in go", a)
// }

func main() {
	// указатель на область в памяти
	// a := new(int);
	// выводим в консоль значение из памяти
	// fmt.Println(*a)

	//
	var a = 1
	// ссылка на адресс, тоже самое что и указатель созданый вручную
	fmt.Println(&a)

	fmt.Println(myFunction(33, 22))
}

func myFunction(a, b int) (int, int) {
	return b, a
}
