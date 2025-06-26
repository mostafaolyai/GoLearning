// package main

// import (
// 	"fmt"
// 	"time"
// )

// type Food struct {
// 	Name  string
// 	Price int
// }

// type Order struct {
// 	Foods []Food
// }

// type Menu map[string]Food

// func createMenu() Menu {
// 	menu := make(Menu)
// 	menu["burger"] = Food{"Burger", 80}
// 	menu["pizza"] = Food{"Pizza", 120}
// 	menu["salad"] = Food{"Salad", 50}
// 	return menu
// }

// func totalPriceCalculator() func(Food) int {
// 	total := 0
// 	return func(f Food) int {
// 		total += f.Price
// 		return total
// 	}
// }

// func takeOrder(orderChan chan Order, menu Menu) {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("Something went wrong during order:", r)
// 		}
// 	}()

// 	var foodNames = []string{"burger", "pizza"}
// 	order := Order{}

// 	for _, name := range foodNames {
// 		food, ok := menu[name]
// 		if !ok {
// 			panic("Food not in menu: " + name)
// 		}
// 		order.Foods = append(order.Foods, food)
// 	}

// 	orderChan <- order
// }

// func main() {
// 	menu := createMenu()
// 	orderChan := make(chan Order)

// 	go takeOrder(orderChan, menu)

// 	order := <-orderChan

// 	calc := totalPriceCalculator()
// 	for _, food := range order.Foods {
// 		fmt.Println("Ordered:", food.Name, "-", food.Price, "Toman")
// 	}

// 	// قیمت کل
// 	for _, food := range order.Foods {
// 		calc(food)
// 	}
// 	fmt.Println("Total price:", calc(Food{Price: 0}), "Toman")

// 	time.Sleep(1 * time.Second)
// }
