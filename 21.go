// // package main

// // import "fmt"

// // // Generic function with type parameter T
// // func PrintValue[T any](value T) {
// // 	fmt.Println(value)
// // }

// // func main() {
// // 	PrintValue("Hello")
// // 	PrintValue(123)
// // 	PrintValue(true)
// // }

// package main

// import "fmt"

// // Type constraint: فقط int یا float64
// type Number interface {
// 	int | float64
// }

// // تابع جنریک با محدودیت Number
// func DoubleValue[T Number](value T) T {
// 	return value * 2
// }

// func main() {
// 	fmt.Println(DoubleValue(5))   // 10
// 	fmt.Println(DoubleValue(3.5)) // 7.0
// 	// fmt.Println(DoubleValue("Hi")) // خطا میده چون string عدد نیست
// }
