// package main

// import (
// 	"fmt"
// )

// // var counter = 0

// // func increment() {
// // 	for i := 0; i < 1000; i++ {
// // 		counter++
// // 	}
// // }

// func main() {
// 	// go increment()
// 	// go increment()
// 	// time.Sleep(time.Second)
// 	// fmt.Println("Counter:", counter)

// 	ch := make(chan string)

// 	go func() {
// 		ch <- "Hello from goroutine!" // ارسال داده
// 	}()

// 	msg := <-ch // دریافت داده
// 	fmt.Println(msg)
// }
