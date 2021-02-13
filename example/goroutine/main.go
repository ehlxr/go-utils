package main

import (
	"time"
)

func main() {
	println("1111")

	// wg := sync.WaitGroup{}
	// wg.Add(1)

	go func() {
		println("begin goroutine.....")
		time.Sleep(time.Second * 2)
		println("end goroutine.....")

		// wg.Done()
	}()

	// wg.Wait()
	println("2222222")
}

// output:
// 1111
// 2222222
