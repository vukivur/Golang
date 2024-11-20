package main

import "fmt"

func main() {

	str := "Go! Run main Go!"

	fmt.Println("Строка: " + str)
	fmt.Println("Длина строки str:", len(str))

	// Так как строки в Go хранятся в виде байтов
	fmt.Println("Первый байт строки:", str[0])
	fmt.Println("Второй байт строки:", str[1])
	fmt.Println("Первый символ строки:", string(str[0]))
	fmt.Println("Второй символ строки:", string(str[1]))

	/*
		Для хранения символов можно использовать int32/rune.
		Здесь используются одинарные кавычки. Компилятор определяет
		код буквы в unicode и присваивает его переменной symbol.
		То есть мы не храним никакую 'c', а храним лишь число 99.
		Функция string() из переданного в него числа 99 делает строку 'c'.
	*/
	var sumbol int32 = 'K'
	fmt.Println(sumbol)
	fmt.Println(string(sumbol))

	var number int = 10
	number++
	fmt.Println(number)

	var name string
	var age int
	fmt.Println("\nВведите имя:")
	//fmt.Scan(&name)
	fmt.Println("Введите возраст:")
	//fmt.Scan(&age)
	fmt.Println("Имя:", name, "Возраст:", age)
}
