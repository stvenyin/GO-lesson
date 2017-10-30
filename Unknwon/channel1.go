
package main

import fmt

func main() {
	c := make(chan bool)
	fmt.Println("xxx")
	go func() {
		fmt.Print("Go Go GO!\n")
		c <- true

	}()
	<-c
}
