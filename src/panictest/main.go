package main 

import (
	"time"
	"fmt"
)

func crash() {
	defer fmt.Println("crash defer")
	panic("kaboom");
}

func main() {
	defer fmt.Println("main defer")
	fmt.Println("before crash()")

	go crash()
	
	fmt.Println("after crash()")
	time.Sleep(3)
	fmt.Println("after sleep")
}
