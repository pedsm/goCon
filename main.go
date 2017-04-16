package main

import "fmt"

//Globals
var buffer int = 0
var iterations int = 10

func main() {
	threadProduce()
	threadConsume()
}

func threadProduce() {
	for i:= 0; i < iterations; i++ {
		produce()
	}
}
func threadConsume() {
	for i:= 0; i < iterations; i++ {
		consume()
	}
}
func produce () {
	buffer++
	fmt.Println(buffer)
}
func consume () {
	buffer--
	fmt.Println(buffer)
}