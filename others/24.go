// // package main

// // import (
// // 	"fmt"
// // )

// // type NotFoundError struct {
// // 	Resource string
// // }

// // func (e NotFoundError) Error() string {
// // 	return fmt.Sprintf("%s not found", e.Resource)
// // }

// // func getUser(id int) error {
// // 	if id != 1 {
// // 		return NotFoundError{Resource: "User"}
// // 	}
// // 	return nil
// // }

// // func main() {
// // 	err := getUser(2)
// // 	if err != nil {
// // 		fmt.Println("Error:", err)
// // 	}
// // }

// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func doSomething() error {
// 	return errors.New("original error")
// }

// func main() {
// 	err := doSomething()
// 	if err != nil {
// 		wrappedErr := fmt.Errorf("operation failed: %w", err)
// 		fmt.Println(wrappedErr)

// 		// Unwrap
// 		unwrapped := errors.Unwrap(wrappedErr)
// 		fmt.Println("Unwrapped:", unwrapped)
// 	}
// }
