package main

import (
	"fmt"
	"time"
)

func greetOne() {
	fmt.Println("Greet 1 function called!")
}

func slowGreet(done chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Slow Greet invoked!")
	done <- true
}

func greetTwo() {
	fmt.Println("Greet 2 invoked!")
}

func greetThree() {
	fmt.Println("Greet 3 invoked!")
}

func main() {

	//	go greetOne()
	done := make(chan bool)
	go slowGreet(done)
	<-done
	//greetTwo()
	//	go greetThree()

}
