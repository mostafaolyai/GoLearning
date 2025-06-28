// package main

// import (
// 	"fmt"
// 	"sync"
// )

// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done() // بعد از پایان این goroutine، یکی از کارهای انجام نشده کم میشه
// 	fmt.Println("Worker", id, "started")
// 	// فرضی: اینجا یه کاری انجام میشه
// 	fmt.Println("Worker", id, "finished")
// }

// func main() {
// 	// var wg sync.WaitGroup

// 	// for i := 1; i <= 3; i++ {
// 	// 	wg.Add(1) // برای هر goroutine یک کار جدید اضافه می‌کنیم
// 	// 	go worker(i, &wg)
// 	// }

// 	// wg.Wait() // منتظر می‌مونه تا همه goroutineها Done بشن
// 	// fmt.Println("All workers done")

// 	// var wg sync.WaitGroup

// 	// for i := 1; i <= 5; i++ {
// 	// 	wg.Add(1)
// 	// 	go func() {
// 	// 		fmt.Println("running")
// 	// 		wg.Done()
// 	// 	}()
// 	// }
// 	// wg.Wait()
// }
