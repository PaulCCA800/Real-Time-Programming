

package main

import(
	"fmt"
)

var i = 0

func increment(incchan, done chan int){
	for j := 0; j < 1000000; j++{
		incchan <- 1
	}
	done <- 1
}

func decrement(decchan, done chan int){
	for j := 0; j < 1000000; j++{
		decchan <- 1
	}
	done <- 1
}

func Server(incchan, decchan, readchan, reply chan int){
	for{
		select{
		case <- incchan:
			i++
		case <- decchan:
			i--
		case <- readchan:
			reply <- i
		}
	}
}

func main() {
	incchan := make(chan int)
	decchan := make(chan int)
	readchan := make(chan int)
	reply := make(chan int)
	done := make(chan int)

	go Server(incchan, decchan, readchan, reply)

	go increment(incchan, done)
	go decrement(decchan, done)

	<- done
	<- done
	readchan <- 1

	fmt.Println("The magic number is", <- reply)
}