package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup              // Создаем переменную WaitGroup для ожидания завершения всех горутин перед выполнением горутины main
var numbers = make(chan int, 10)   // Создаем буферизированный канал для передачи в него случайных чисел из первой горутины GenNumSlice
var squareNum = make(chan int, 10) // Создаем второй буферизированны канал для передачи в него измененных чисел из второй горутины SquareNum

func main() {
	res := make([]int, 0, 10) // Создаем слайс для помещения туда рузультатов работы горутин

	wg.Add(1)        // Создаем счетчик первой горутины
	go GenNumSlice() // Запускаем первую горутину
	wg.Add(1)        // Создаем счетчик второй горутины
	go SquareNum()   // Запускаем вторую горутину

	wg.Wait()        // Ждем, пока горутины завершат свою работу
	close(squareNum) // Закрываем второй канал
	// проходимся циклом по второму каналу и добавляем из него каждое значение в слайс
	for num := range squareNum {
		res = append(res, num)
	}
	// проходимся по слайсу и выводим в консоль каждое его значение
	for _, val := range res {
		fmt.Printf("%d ", val)
	}
}

// Функция GenNumSlice передает в канал numbers 10 случайных чисел
func GenNumSlice() {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		r := rand.Intn(100)
		numbers <- r
	}
	close(numbers)
}

// Функция SquareNum проходится по каналу numbers, возводит каждое число из канала во вторую степень и записывает измененные числа в канал squareNum
func SquareNum() {
	defer wg.Done()
	for num := range numbers {
		squareNum <- num * num
	}
}
