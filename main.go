package main

import (
	"fmt"
	"sync"
)

//Globals
var buffer int = 0
var iterations int = 100000
var produceChannel chan int = make(chan int)
var consumeChannel chan int = make(chan int)
var mutex = &sync.Mutex{}

func main() {
	go threadProduce()
	go threadConsume()
	for{
		pC := <-produceChannel
		cC := <-consumeChannel
		if pC != 0 && cC != 0 {
			break
		}
	}
}

func threadProduce() {
	for i:= 0; i < iterations; i++ {
		produce()
	}
	produceChannel <- 1
}
func threadConsume() {
	for i:= 0; i < iterations; i++ {
		consume()
	}
	consumeChannel <- 1
}
func produce () {
	mutex.Lock()
	buffer++
	mutex.Unlock()
	fmt.Println(buffer)
}
func consume () {
	mutex.Lock()
	buffer--
	mutex.Unlock()
	fmt.Println(buffer)
}