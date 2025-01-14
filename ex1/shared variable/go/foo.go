// Use `go run foo.go` to run your program

package main

import (
    . "fmt"
    "runtime"
    "time"
)

var i = 0

func incrementing() {
    for j := 0; j < 1000000 ; j++ {
        i++
    }
}

func decrementing() {
    for j := 0; j < 1000000 ; j++{
        i--
    }
}

func main() {
    // What does GOMAXPROCS do? What happens if you set it to 1?
    //GOMAXPROCS sets the maximum number of number of CPU's Go can use. 
    runtime.GOMAXPROCS(2)    

    go incrementing()
    go decrementing()
	
    // TODO: Spawn both functions as goroutines
	
    // We have no direct way to wait for the completion of a goroutine (without additional synchronization of some sort)
    // We will do it properly with channels soon. For now: Sleep.
    time.Sleep(500*time.Millisecond)

    Println("The magic number is:", i)
}
