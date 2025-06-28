// package main

// import "fmt"

// type Box[T any] struct {
// 	value T
// }

// func (b Box[T]) GetValue() T {
// 	return b.value
// }

// func main() {
// 	intBox := Box[int]{value: 42}
// 	stringBox := Box[string]{value: "Hello, Generics"}

// 	fmt.Println(intBox.GetValue())    // 42
// 	fmt.Println(stringBox.GetValue()) // Hello, Generics
// }
