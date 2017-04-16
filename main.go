package main

import "fmt"

func main() {
    var buffer int = 0
    var iterations int = 10
    fmt.Println("Hello World")
    for i := 0; i < iterations; i++ {
        buffer = produce(buffer)
    }
    for i := 0; i < iterations; i++ {
        buffer = consume(buffer)
    }
}

func produce(buf int) int {
    buf++
    fmt.Println(buf)
    return buf
}

func consume(buf int) int {
    buf--
    fmt.Println(buf)
    return buf
}