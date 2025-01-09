package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg sync.WaitGroup
var numbers = make(chan int, 10)
var squareNum = make(chan int, 10)

func main() {
	res := make([]int, 0, 10)

	wg.Add(1)
	go GenNumSlice()
	wg.Add(1)
	go SquareNum()

	wg.Wait()
	close(squareNum)
	for num := range squareNum {
		res = append(res, num)
	}
	for _, val := range res {
		fmt.Printf("%d ", val)
	}
}

func GenNumSlice() {
	defer wg.Done()

	for i := 0; i < 10; i++ {
		r := rand.Intn(100)
		numbers <- r
	}
	close(numbers)
}

func SquareNum() {
	defer wg.Done()
	for num := range numbers {
		squareNum <- num * num
	}
}
