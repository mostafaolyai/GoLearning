// package main

// import "fmt"

// func main() {
// 	var anything interface{} = 5
// 	describe(anything)

// 	anything = "hello"
// 	describe(anything)

// 	anything = true
// 	describe(anything)

// 	// str := anything.(string)
// 	// fmt.Println(str)

// 	// str, ok := anything.(int)
// 	// if ok {
// 	// 	fmt.Println("Value is:", str)
// 	// } else {
// 	// 	fmt.Println("Not a string")
// 	// }
// }
// func describe(i interface{}) {
// 	switch v := i.(type) {
// 	case string:
// 		fmt.Println("It's a string:", v)
// 	case int:
// 		fmt.Println("It's an int:", v)
// 	default:
// 		fmt.Println("Unknown type")
// 	}
// }
