// package main

// import (
// 	"fmt"
// 	"time"
// )

// // func sayHello() {
// // 	fmt.Println("Hello from goroutine")
// // }

// // func main() {
// // 	go sayHello()               // اجرا در goroutine جدا
// // 	time.Sleep(1 * time.Second) // صبر کن تا goroutine تموم شه
// // 	fmt.Println("Main function done")
// // }

// func printNumber(n int) {
// 	fmt.Println(n)
// }

// func main() {
// 	go printNumber(10)
// 	go printNumber(20)
// 	time.Sleep(1 * time.Second)
// }
