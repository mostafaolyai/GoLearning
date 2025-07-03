// package main

// import "fmt"

// // Product Interface
// type Animal interface {
// 	Speak() string
// }

// // Concrete types
// type Dog struct{}
// func (d Dog) Speak() string { return "Woof!" }

// type Cat struct{}
// func (c Cat) Speak() string { return "Meow!" }

// // Factory function
// func AnimalFactory(animalType string) Animal {
// 	if animalType == "dog" {
// 		return Dog{}
// 	} else if animalType == "cat" {
// 		return Cat{}
// 	}
// 	return nil
// }

// func main() {
// 	animal := AnimalFactory("dog")
// 	fmt.Println(animal.Speak()) // Woof!
// }

// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// )

// type config struct {
// 	AppName string
// }

// var instance *config
// var once sync.Once

// func GetConfig() *config {
// 	once.Do(func() {
// 		fmt.Println("Initializing config... "+time.Now().Format(time.RFC3339))
// 		instance = &config{AppName: "GoApp"}
// 	})
// 	return instance
// }

// func main() {
// 	c1 := GetConfig()
// 	c2 := GetConfig()
// 	fmt.Println(c1 == c2)         // true
// 	fmt.Println(c1.AppName)       // GoApp
// }

// package main

// import "fmt"

// // Strategy interface
// type PaymentStrategy interface {
// 	Pay(amount int)
// }

// // Concrete strategies
// type CreditCard struct{}
// func (c CreditCard) Pay(amount int) {
// 	fmt.Println("Paying", amount, "with Credit Card")
// }

// type PayPal struct{}
// func (p PayPal) Pay(amount int) {
// 	fmt.Println("Paying", amount, "with PayPal")
// }

// // Context
// type PaymentProcessor struct {
// 	strategy PaymentStrategy
// }

// func (p *PaymentProcessor) SetStrategy(strategy PaymentStrategy) {
// 	p.strategy = strategy
// }

// func (p *PaymentProcessor) Pay(amount int) {
// 	p.strategy.Pay(amount)
// }

// func main() {
// 	pp := &PaymentProcessor{}

// 	pp.SetStrategy(CreditCard{})
// 	pp.Pay(100)

// 	pp.SetStrategy(PayPal{})
// 	pp.Pay(200)
// }


package main

import "fmt"

// Component
type Notifier interface {
	Send(message string)
}

// Concrete component
type EmailNotifier struct{}
func (e EmailNotifier) Send(message string) {
	fmt.Println("Sending email:", message)
}

// Decorator
type SMSDecorator struct {
	Notifier
}

func (s SMSDecorator) Send(message string) {
	s.Notifier.Send(message)
	fmt.Println("Sending SMS:", message)
}

func main() {
	email := EmailNotifier{}
	notifier := SMSDecorator{Notifier: email}
	notifier.Send("Hello")
}