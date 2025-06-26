// package main

// import "fmt"

// // نوع‌هایی که می‌خوایم قبول کنیم
// type SupportedTypes interface {
// 	int | float64
// }

// // تابع جنریک که با int یا float64 یا string کار می‌کنه
// func Max[T SupportedTypes](a, b T) T {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func main() {
// 	fmt.Println(Max(3, 4))     // 7
// 	fmt.Println(Max(2.5, 1.5)) // 4
// 	// fmt.Println(Max("Go", "lang")) // Golang

// 	// fmt.Println(Max(true, false)) // ❌ کامپایل نمی‌شه چون bool پشتیبانی نمی‌شه
// }
