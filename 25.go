// package main

// import "fmt"

// const (
// 	Sunday    = iota // 0
// 	Monday           // 1
// 	Tuesday          // 2
// 	Wednesday        // 3
// 	Thursday         // 4
// 	Friday           // 5
// 	Saturday         // 6
// )

// const (
// 	Red   = iota + 3 // 1
// 	Green = iota     // 2
// 	Blue             // 3
// )

// const (
// 	_ = iota // 0 (رد شده)
// 	A        // 1
// 	B        // 2
// )

// func main() {
// 	// fmt.Println(Red)   // 1
// 	// fmt.Println(Green) // 4
// 	// fmt.Println(Blue)  // 4

// 	// const (
// 	// 	X = iota + 2 // ?
// 	// 	Y            // ?
// 	// 	Z = 10       // ?
// 	// 	W = iota     // ?
// 	// )
// 	// fmt.Println(X, Y, Z, W)

// 	// const (
// 	// 	A = iota
// 	// 	B
// 	// 	C = 2
// 	// 	D
// 	// )

// 	// fmt.Println(A, B, C, D)

// 	// const (
// 	// 	X = iota + 1
// 	// 	Y
// 	// 	Z
// 	// )
// 	// fmt.Println(X, Y, Z)

// 	// const (
// 	// 	A = iota + 1 //1
// 	// 	B            //2
// 	// 	C = 5        //5
// 	// 	D            //5
// 	// 	E = iota     //4
// 	// )

// 	// fmt.Println(A, B, C, D, E)

// 	// const (
// 	// 	X = iota //0
// 	// 	Y = X    //0
// 	// 	Z = iota //2
// 	// )

// 	// fmt.Println(X, Y, Z)

// 	// const (
// 	// 	A = iota * 2 //0
// 	// 	B = iota * 3 //3
// 	// 	C = 10
// 	// 	D        //10
// 	// 	E = iota //4
// 	// )

// 	// fmt.Println(A, B, C, D, E)

// 	// const (
// 	// 	_  = iota
// 	// 	KB = 1 << (10 * iota)//1024
// 	// 	MB//2^20
// 	// 	GB//2^30
// 	// )

// 	// fmt.Println(KB, MB, GB)

// 	const (
// 		A = iota     //0
// 		B = iota + 2 //3
// 		C = iota * 3 //6
// 	)

// 	fmt.Println(A, B, C)
// }

// // const (
// // 	Read = 1 << iota    // 1 (0001)
// // 	Write               // 2 (0010)
// // 	Execute             // 4 (0100)
// // )

// // func main() {
// // 	perm := Read | Execute
// // 	fmt.Println(perm) // 5 (0001 | 0100 = 0101)
// // }
