package pkg

import (
	"fmt"
	"runtime"
	"sync"
)

func GoRoutine() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("goroutine invoke")
	}()
	wg.Wait()
	fmt.Printf("num of working goroutines: %d\n", runtime.NumGoroutine())
	fmt.Println("main func finish")
}
