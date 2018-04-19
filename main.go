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
	// var a = 1
	// ссылка на адресс, тоже самое что и указатель созданый вручную
	// fmt.Println(&a)

	// fmt.Println(myFunction(33, 22))

	myFunc()
}

// func myFunction(a, b int) (int, int) {
// 	return b, a
// }

// 1й способ создать массив
// func myFunc() {
// 	var a = []string{"hello", "world"}
// 	fmt.Println(a)
// }

// второй способ создать массив
// func myFunc() {
// 	var arr = make([]string, 5)
// 	fmt.Print(arr)
// }

// добавить элементы в массив => append
// func myFunc() {
// 	var arr = make([]uint, 5)
// 	arr = append(arr, 5, 6, 7, 8)
// 	fmt.Println(arr)
// }

// for loop
// func myFunc() {
// 	var arr = []uint{1, 2, 3, 4, 5, 6, 7, 10}
// 	for i := 0; i < len(arr); i++ {
// 		fmt.Println(arr[i])
// 	}
// }

// another iterator
func myFunc() {
	var arr = []uint{1, 2, 3, 4, 5, 6, 7, 10}
	for _, value := range arr {
		fmt.Println(value)
	}
}
