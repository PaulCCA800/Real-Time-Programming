package main

import(
	"fmt"
	"sync"
)

type Server struct{
	i int
	mu sync.Mutex
}

func (s *Server) increment (done chan bool){
	for j := 0; j < 1000000; j++ {
		s.mu.Lock()
		s.i++
		s.mu.Unlock()
	}
	done <- true
}

func (s *Server) decrement (done chan bool){
	for j := 0; j < 1000000; j++ {
		s.mu.Lock()
		s.i--
		s.mu.Unlock()
	}
	done <- true
}

func (s *Server) read () int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.i
}

func main(){
	server := &Server{i: 0}
	done := make(chan bool)

	go server.increment(done)
	go server.decrement(done)

	<-done
	<-done

	value := server.read()
	fmt.Println("the magic number is", value)
}