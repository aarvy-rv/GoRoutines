package main

import (
	"fmt"
	"time"
)

func greetOne(done chan bool) {
	fmt.Println("Greet 1 function called!")
	done <- true
}

func slowGreet(done chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Slow Greet invoked!")
	done <- true
}

func greetTwo(done chan bool) {
	fmt.Println("Greet 2 invoked!")
	done <- true
}

func greetThree(done chan bool) {
	fmt.Println("Greet 3 invoked!")
	done <- true
}

func main() {

	//done := make(chan bool) // We can use one channel for multiple go routines, however we need to read the channels
	// for all the number of times go routines is invokek i.e. 4 times in this example to avoid the kind of race condition

	dones := make([]chan bool, 4)

	dones[0] = make(chan bool)
	go greetOne(dones[0])

	dones[1] = make(chan bool)
	go slowGreet(dones[1])

	dones[2] = make(chan bool)
	go greetTwo(dones[2])

	dones[3] = make(chan bool)
	go greetThree(dones[3])

	for _, done := range dones {
		<-done
	}

}
