package main

import (
	"fmt"
	"sync"
)

//Globals
var buffer int = 0
var iterations int = 100000 //Increase the iterations to check for data races.
//Channels allow the main thread to identify the end of a child thread.
var produceChannel chan int = make(chan int) 
var consumeChannel chan int = make(chan int)
//Mutex to prevent Data races
var mutex = &sync.Mutex{}

func main() {
	//Start goroutines
	go threadProduce()
	go threadConsume()
	for{
		//Waiting for child threads to finish
		pC := <-produceChannel
		cC := <-consumeChannel
		if pC != 0 && cC != 0 {
			break
		}
	}
}

//Thread Creation functions
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
//Helper functions can be implemented inside the thread logic
func produce () {
	mutex.Lock()
	buffer++
	mutex.Unlock()
	fmt.Println(buffer)
}
func consume () {
	//Waiting for buffer to be greater than 0
	for{
		//Use of mutexes to prevent data races
		mutex.Lock()
		if buffer > 0 {
			mutex.Unlock()
			break
		}
		mutex.Unlock() //Make sure to unlock the mutex in order for the producer to work
	}
	mutex.Lock()
	buffer--
	mutex.Unlock()
	fmt.Println(buffer)
}