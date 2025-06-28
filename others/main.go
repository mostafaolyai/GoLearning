// package main

// import "fmt"

// func main() {
// 	// var age int
// 	// var name string
// 	// var score float64
// 	// var active bool

// 	// fmt.Println("Age:", age)
// 	// fmt.Println("Name:", name)
// 	// fmt.Println("Score:", score)
// 	// fmt.Println("Active:", active)

// 	// // حالا مقدار جدید بده و چاپ کن
// 	// age = 28
// 	// name = "Kamran"
// 	// score = 98.5
// 	// active = true

// 	// fmt.Println("Updated age:", age)
// 	// fmt.Println("Updated name:", name)
// 	// fmt.Println("Updated score:", score)
// 	// fmt.Println("Updated active:", active)

// 	// var score float32
// 	// score = 2.5
// 	// score = score + 2
// 	// fmt.Println(score)

// 	// // age := 20
// 	// if (score >= 12) {
// 	//     fmt.Println("you score is greather equal than 12")
// 	// } else if score> 9 {
// 	//     fmt.Println("you score is greather than 9")
// 	// } else {
// 	//     fmt.Println("you score is less than 9")
// 	// }

// 	// if x := getNumber(); x > 0 {
// 	//     fmt.Println("عدد + است")
// 	// } else {
// 	//     fmt.Println("عدد 0 یا - است")
// 	// }

// 	// for initialization; condition; post {
// 	//     // بدنه‌ی حلقه
// 	// }

// 	// for i := 0; i < 5; i++ {
// 	// 	fmt.Println(i)
// 	// }
// 	// i := 0
// 	// for i < 5 {
// 	// 	fmt.Println(i)
// 	// 	i++
// 	// }
// 	// for {
// 	// 	fmt.Println("این حلقه تا ابد اجرا می‌شه")
// 	// }

// 	// for {
// 	// 	fmt.Print("Enter a number: ")
// 	// 	var n int
// 	// 	fmt.Scanln(&n)
// 	// 	if n == 0 {
// 	// 		break
// 	// 	}
// 	// }

// 	// for i := 0; i < 5; i++ {
// 	// 	if i == 2 {
// 	// 		break
// 	// 	}
// 	// 	fmt.Println(i)
// 	// }

// 	// nums := []int{10, 20, 30}
// 	// for i, val := range nums {
// 	// 	fmt.Println("index:", i, "value:", val)
// 	// }

// 	// for i := 1; i <= 5; i++ {
// 	// 	if i%2 == 0 {
// 	// 		continue
// 	// 	}
// 	// 	fmt.Println(i)
// 	// }

// 	// age := 17

// 	// if age >= 18 {
// 	// 	fmt.Println("you can have liscense")

// 	// 	if age >= 21 {
// 	// 		fmt.Println("you can buy a car.")
// 	// 	}
// 	// } else {
// 	// 	fmt.Println("you are so cute!")
// 	// }

// 	// day := 1

// 	// switch day {
// 	// case 1:
// 	// 	fmt.Println("1")
// 	// 	fallthrough
// 	// case 2:
// 	// 	fmt.Println("2")
// 	// case 3:
// 	// 	fmt.Println("3")
// 	// default:
// 	// 	fmt.Println("X")
// 	// }
// 	// level := 2

// 	// switch level {
// 	// case 1:
// 	// 	fmt.Println("دسترسی به دوره‌های رایگان")
// 	// 	fallthrough
// 	// case 2:
// 	// 	fmt.Println("دسترسی به دوره‌های پولی")
// 	// 	fallthrough
// 	// case 3:
// 	// 	fmt.Println("دسترسی به مشاوره اختصاصی")
// 	// }

// 	// x := 1
// 	// switch x {
// 	// case 1:
// 	// 	fmt.Println("A")
// 	// 	fallthrough
// 	// case 2:
// 	// 	fmt.Println("B")
// 	// default:
// 	// 	fmt.Println("C")
// 	// }

// 	// 	var numbers [3]int
// 	// 	numbers[0] = 10
// 	// 	numbers[1] = 20
// 	// 	numbers[2] = 30
// 	// 	fmt.Println(numbers)

// 	// 	letters := [4]string{"a", "b", "c", "d"}
// 	// 	fmt.Println(letters)

// 	// 	//Slice
// 	// 	nums := []int{1, 2, 3}
// 	// fmt.Println(nums)

// 	// s := make([]int, 3) // طول 3
// 	// fmt.Println(s)

// 	// nums := []int{1, 2, 3}
// 	// nums = append(nums, 4, 5)
// 	// fmt.Println(nums)

// 	// letters := []string{"a", "b", "c", "d", "e"}
// 	// fmt.Println(letters[1:4])

// 	// a := []int{1, 2, 3}
// 	// b := a
// 	// b[0] = 100
// 	// fmt.Println("a: ", a)
// 	// fmt.Println("b: ", b)

// 	// c := [3]int{1, 2, 3}
// 	// d := c
// 	// d[0] = 100
// 	// fmt.Println("c: ", c)
// 	// fmt.Println("d: ", d)

// 	// a := []int{1, 2, 3}
// 	// b := make([]int, len(a))
// 	// copy(b, a)
// 	// b[0] = 100
// 	// fmt.Println(a)
// 	// fmt.Println(b)

// 	//Map
// 	// ages := map[string]int{
// 	// 	"Ali":   25,
// 	// 	"Reza":  30,
// 	// 	"Sarah": 28,
// 	// }

// 	// ages["Mosi"] = 37
// 	// delete(ages, "Sarah")
// 	// fmt.Println(ages)
// 	// fmt.Println(ages["Sarah"])

// 	// value, exists := ages["Sarah"]
// 	// if exists {
// 	// 	fmt.Println("Sarah exists:", value)
// 	// } else {
// 	// 	fmt.Println("Not found")
// 	// }

// 	// for key, value := range ages {
// 	// 	fmt.Println(key, value)
// 	// }

// 	// m := make(map[string]int)
// 	// m["A"] = 1
// 	// m["B"] = 2
// 	// m2 := m
// 	// m2["A"] = 100
// 	// fmt.Println(m["A"])
// 	// a, b := add(1, 2)

// 	// fmt.Println(a, b)

// 	//Quiz of Day 5
// 	// for i := 1; i < 100; i++ {
// 	// 	if i%2 == 0 {
// 	// 		fmt.Println("4")
// 	// 	} else {
// 	// 		fmt.Println("5")
// 	// 	}
// 	// }

// 	// for {
// 	// 	fmt.Print("Enter a number: ")
// 	// 	var n int
// 	// 	fmt.Scanln(&n)
// 	// 	if n < 0 {
// 	// 		break
// 	// 	}
// 	// 	fmt.Println(n)
// 	// }

// 	// nums := []int{1, 4, 7, 10, 13}

// 	// for index, value := range nums {
// 	// 	if value%2 == 0 {
// 	// 		fmt.Println(index, value)
// 	// 	}
// 	// }

// 	// a := 10
// 	// b := test(a)
// 	// fmt.Println(b)

// 	//Pointers
// 	// var x int = 10
// 	// var p *int = &x // p یک اشاره‌گر به x است

// 	// fmt.Println(x, p)

// 	// x := 10
// 	// p := &x

// 	// fmt.Println("x =", x)
// 	// fmt.Println("p =", p)   // آدرس x
// 	// fmt.Println("*p =", *p) // مقدار x از طریق آدرس

// 	// *p = 20 // مقدار x رو عوض کردیم از طریق اشاره‌گر

// 	// fmt.Println("x بعد از تغییر =", x)

// 	// x := 5
// 	// addOne(&x)
// 	// fmt.Println(x) // 6

// 	//Practice Day6
// 	// day := 0

// 	// fmt.Println("Enter your Day: ")
// 	// fmt.Scanln(&day)

// 	// switch day {
// 	// case 0:
// 	// 	fmt.Println("Saturday")
// 	// case 1:
// 	// 	fmt.Println("Sunday")
// 	// case 2:
// 	// 	fmt.Println("Monday")
// 	// case 3:
// 	// 	fmt.Println("Tusday")
// 	// case 4:
// 	// 	fmt.Println("Wednsday")
// 	// case 5:
// 	// 	fmt.Println("Tursday")
// 	// case 6:
// 	// 	fmt.Println("Friday")
// 	// default:
// 	// 	fmt.Println("Add valid Day")
// 	// }
// 	//-----
// 	// num := -8

// 	// if num > 0 {
// 	// 	if num%2 == 0 {
// 	// 		fmt.Println("/2")
// 	// 	} else {
// 	// 		fmt.Println("!/2")
// 	// 	}
// 	// } else {
// 	// 	fmt.Println("-")
// 	// }

// 	//Day7 tell them 11
// 	// array := [5]int{1, 2, 5, 8, 9}
// 	// for _, value := range array {
// 	// 	fmt.Println(value)
// 	// }

// 	// friends := []string{"Sirus", "Peyman"}

// 	// friends = append(friends, "Mamad")

// 	// fmt.Println(friends)

// 	// nums := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// 	// for _, val := range nums {
// 	// 	if val%2 == 0 {
// 	// 		fmt.Println(val)
// 	// 	}
// 	// }

// 	//Day 8
// 	// countries := map[string]string{
// 	// 	"Iran": "Tehran",
// 	// 	"Iraq": "Baqdad",
// 	// 	"UK":   "London",
// 	// }

// 	// countries["US"] = "NewYork"

// 	// if countries["UK"] != "" {
// 	// 	fmt.Println(countries["UK"])
// 	// }

// 	// for key, val := range countries {
// 	// 	fmt.Println(key, val)
// 	// }

// 	// x := 10
// 	// y := 20
// 	// test(x, &y)
// 	// fmt.Println(x, y)
// 	// p1 := Person{name: "Ali", age: 30}
// 	// fmt.Println(p1.name) // Ali
// 	// fmt.Println(p1.age)  // 30

// 	// var p2 Person
// 	// p2.name = "Sara"
// 	// p2.age = 25
// 	// fmt.Println(p2.name)
// 	// fmt.Println(p2.age)

// 	// p1 := Person{name: "Ali", age: 30}
// 	// p2 := p1
// 	// p2.name = "Reza"

// 	// fmt.Println(p1.name)

// 	// p := &Person{name: "Omid", age: 40}
// 	// // p.age = 41
// 	// // fmt.Println(p.name)
// 	// // p1 := p
// 	// // p1.name = "Mosi"
// 	// // fmt.Println(p.name)
// 	// // fmt.Println(p1.name)

// 	// printPerson(p)

// 	// nums := [3]int{1, 2, 3}
// 	// for index, value := range nums {
// 	// 	fmt.Println(value)
// 	// }

// 	// u := User{name: "Original", age: 22}
// 	// changeName(u)
// 	// fmt.Println(u.name)

// 	//Interfaces
// 	// var s Speaker
// 	// s = Human{name: "Ali"}
// 	// s.Speak()

// 	// u := User{name: "Reza"}
// 	// fmt.Println(u) // User: Reza

// 	// var w Writer
// 	// w = Pen{}
// 	// w.Write()

// 	//Day 13

// 	// f := add
// 	// fmt.Println(f(3, 4)) // خروجی: 7

// 	// result := apply(multiply, 2, 3)
// 	// fmt.Println(result) // خروجی: 6خروجی:

// 	// double := makeMultiplier(2)

// 	// nums := []int{1, 2, 3, 4, 5}
// 	// for _, num := range nums {
// 	// 	fmt.Println(double(num))
// 	// }
// 	add := func(a, b int) int {
// 		return a + b
// 	}
// 	fmt.Println(add(10, 5))
// }

// func makeMultiplier(factor int) func(int) int {
// 	return func(x int) int {
// 		return x * factor
// 	}
// }

// // func apply(op func(int, int) int, a int, b int) int {
// // 	return op(a, b)
// // }

// // func multiply(x, y int) int {
// // 	return x * y
// // }

// // func add(a, b int) int {
// // 	return a + b
// // }

// // type Writer interface {
// // 	Write()
// // }

// // type Pen struct{}

// // func (p Pen) Write() {
// // 	fmt.Println("Writing with pen")
// // }

// // type User struct {
// // 	name string
// // }

// // func (u User) String() string {
// // 	return "User: " + u.name
// // }

// type Speaker interface {
// 	Speak()
// }

// // type Human struct {
// // 	name string
// // }

// // func (h Human) Speak() {
// // 	fmt.Println(h.name, "says hello")
// // }

// // type User struct {
// // 	name string
// // 	age  int
// // }

// // func changeName(u User) {
// // 	u.name = "Changed"
// // }

// // type Person struct {
// // 	name string
// // 	age  int
// // }

// // func printPerson(p *Person) {
// // 	fmt.Println(p.name, p.age)
// // }

// // func test(a int, b *int) {
// // 	a = a + 5
// // 	*b = *b + 5
// // }

// // func addOne(n *int) {
// // 	*n = *n + 1
// // }

// // func test(x int) int {
// // 	x = x + 5
// // 	return x
// // }

// // func add(a, b int) (x int, y int) {
// // 	return a + b, b - a
// // }

// func getNumber() int {
// 	return -12
// }
