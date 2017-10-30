package main	
	
	c := make(chan bool, 0)
	fmt.Println("xxx")
	go func() {
		fmt.Print("Go Go GO!\n")
		<-c
	}()
	c <- true