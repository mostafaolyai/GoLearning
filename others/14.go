// package main

// import "fmt"

// // func counter() func() int {
// // 	count := 0
// // 	return func() int {
// // 		count++
// // 		return count
// // 	}
// // }

// //	func makeValidator(min int, max int) func(int) bool {
// //		return func(value int) bool {
// //			return value >= min && value <= max
// //		}
// //	}
// //
// //	func makeMultiplier(x int) func(int) int {
// //		return func(y int) int {
// //			return x * y
// //		}
// //	}
// func adder(a int) func(int) int {
// 	return func(b int) int {
// 		return a + b
// 	}
// }

// func main() {
// 	// c := counter()

// 	// fmt.Println(c())
// 	// fmt.Println(c())
// 	// fmt.Println(c())

// 	// c1 := counter()

// 	// fmt.Println(c1())
// 	// fmt.Println(c1())

// 	// ageValidator := makeValidator(18, 35)

// 	// fmt.Println(ageValidator(20))
// 	// fmt.Println(ageValidator(40))

// 	// mult2 := makeMultiplier(2)
// 	// mult3 := makeMultiplier(3)

// 	// fmt.Println(mult2(5))
// 	// fmt.Println(mult3(5))

// 	add5 := adder(5)
// 	fmt.Println(add5(10))
// }
