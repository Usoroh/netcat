package main

import (
	"log"
	"net"
	"os"
)

func main() {

	args := os.Args[1:]
	port := ":" + args[0]

	s := newServer()
	go s.run()

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("unable to start server ", err.Error())
	}
	defer listener.Close()

	log.Printf("Started server on port %s", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Unable to accept connection", err.Error())
			continue
		}

		go s.newClient(conn)
	}
}

// func foo(c chan int, someValue int) {
// 	c <- someValue * 5
// }

// func main() {
// 	fooVal := make(chan int)

// 	go foo(fooVal, 5)
// 	go foo(fooVal, 3)

// 	v1 := <-fooVal
// 	v2 := <-fooVal

// 	fmt.Println(v1, v2)
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var wg sync.WaitGroup

// func cleanup() {
// 	defer wg.Done()
// 	if r := recover(); r != nil {
// 		fmt.Println("Recovered in cleanup ", r)
// 	}
// }

// func say(s string) {

// 	defer cleanup()
// 	for i := 0; i < 3; i++ {
// 		fmt.Println(s)
// 		if i == 2 {
// 			panic("Oh dear, a 2")
// 		}
// 	}
// }

// func main() {
// 	wg.Add(1)
// 	go say("Hey")
// 	wg.Add(1)
// 	go say("There")
// 	wg.Wait()
// }
