
package main

import (
	"fmt"
)

func main() {
	c := make(chan string)
	go func() { c <- "hello" } ()
	msg := <-c
	fmt.Println(msg)
}