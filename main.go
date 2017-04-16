package main

import "fmt"

func main() {
    var buffer int = 0
    var iterations int = 10
	produce := func (buf int) int {
		buf++
		fmt.Println(buf)
		return buf
	}
	consume := func (buf int) int {
		buf--
		fmt.Println(buf)
		return buf
	}
	for i := 0; i < iterations; i++ {
		buffer = produce(buffer)
	}
	for i := 0; i < iterations; i++ {
		buffer = consume(buffer)
	}
}

