package pkg

import (
	"context"
	"fmt"
	"time"
)

func Context() {
	// var wg sync.WaitGroup
	// ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	// defer cancel()

	// wg.Add(3)
	// go subTask(ctx, &wg, "a")
	// go subTask(ctx, &wg, "b")
	// go subTask(ctx, &wg, "c")
	// wg.Wait()

	// ---------------------------------------
	// var wg sync.WaitGroup
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	v, err := criticalTask(ctx)
	// 	if err != nil {
	// 		fmt.Println("critical task canceled due to: ", err)
	// 		cancel()
	// 		return
	// 	}
	// 	fmt.Println("success", v)
	// }()

	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	v, err := normalTask(ctx)
	// 	if err != nil {
	// 		fmt.Println("normal task canceled due to: ", err)
	// 		return
	// 	}
	// 	fmt.Println("success", v)
	// }()

	// wg.Wait()

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add((20 * time.Millisecond)))
	defer cancel()
	ch := subTask(ctx)
	v, ok := <-ch
	if ok {
		fmt.Println(v)
	}
	fmt.Println("finish")
}

func subTask(ctx context.Context) <-chan string {
	ch := make(chan string)

	go func() {
		defer close(ch)
		deadline, ok := ctx.Deadline()
		if ok {
			if deadline.Sub(time.Now().Add(30*time.Millisecond)) < 0 {
				fmt.Println("impossible to meet deadline")
				return
			}
		}
		time.Sleep(30 * time.Millisecond)
		ch <- "hello"
	}()
	return ch
}

// func criticalTask(ctx context.Context) (string, error) {
// 	ctx, cancel := context.WithTimeout(ctx, 1200*time.Millisecond)
// 	defer cancel()

// 	t := time.NewTicker(1000 * time.Millisecond)
// 	select {
// 	case <-ctx.Done():
// 		return "", ctx.Err()
// 	case <-t.C:
// 		t.Stop()
// 	}
// 	return "A", nil
// }

// func normalTask(ctx context.Context) (string, error) {
// 	t := time.NewTicker(3000 * time.Millisecond)
// 	select {
// 	case <-ctx.Done():
// 		return "", ctx.Err()
// 	case <-t.C:
// 		t.Stop()
// 	}
// 	return "B", nil
// }

// func subTask(ctx context.Context, wg *sync.WaitGroup, id string) {
// 	defer wg.Done()
// 	t := time.NewTicker(500 * time.Millisecond)

// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println(ctx.Err())
// 			return
// 		case <-t.C:
// 			t.Stop()
// 			fmt.Println(id)
// 			return
// 		}
// 	}
// }
