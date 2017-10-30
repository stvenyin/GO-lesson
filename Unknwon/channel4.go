package main

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Go(&wg)
	}
	wg.Wait()

}

func Go(wg *sync.WaitGroup) {
	fmt.Printf("Go Go Go!")
	wg.Done()
}
