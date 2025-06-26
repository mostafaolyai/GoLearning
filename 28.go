package main

import (
	"fmt"
	"reflect"
)

// func main() {
// 	x := 42

// 	fmt.Println("Type:", reflect.TypeOf(x))   // int
// 	fmt.Println("Value:", reflect.ValueOf(x)) // 42
// }

// type User struct {
// 	Name string
// 	Age  int
// }

// func main() {
// 	u := User{"Reza", 30}
// 	val := reflect.ValueOf(u)

// 	for i := 0; i < val.NumField(); i++ {
// 		fmt.Printf("Field %d: %v\n", i, val.Field(i))
// 	}
// }

func main() {
	x := 10
	v := reflect.ValueOf(&x).Elem() // به مقدار اشاره‌گر دسترسی پیدا کن

	fmt.Println("Before:", x)
	v.SetInt(99)
	fmt.Println("After:", x)
}
