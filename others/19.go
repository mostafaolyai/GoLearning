// package main

// import "fmt"

// // func main() {
// // 	// defer fmt.Println("World")
// // 	// fmt.Println("Hello")

// // 	panic("Something went wrong!")
// // 	fmt.Println("This will not be printed")
// // }

// func safe() {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("Recovered from:", r)
// 		}
// 	}()
// 	panic("Crash")
// }
// func test() {
// 	defer fmt.Println("1")
// 	defer fmt.Println("2")
// 	defer fmt.Println("3")
// }
// func main() {
// 	safe()
// 	fmt.Println("After recovery")
// }
