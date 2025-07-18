// type Invoice struct {
// 	Amount float64
// }

// func (i Invoice) CalculateTax() float64 {
// 	return i.Amount * 0.1
// }

// func (i Invoice) Print() {
// 	fmt.Println("Invoice amount:", i.Amount)
// }

// type Invoice struct {
// 	Amount float64
// }

// func (i Invoice) CalculateTax() float64 {
// 	return i.Amount * 0.1
// }

// type InvoicePrinter struct{}

// func (p InvoicePrinter) Print(i Invoice) {
// 	fmt.Println("Invoice:", i.Amount)
// }

// type Discount interface {
// 	Apply(price float64) float64
// }

// type BlackFriday struct{}
// func (b BlackFriday) Apply(price float64) float64 {
// 	return price * 0.8
// }

// type NewYear struct{}
// func (n NewYear) Apply(price float64) float64 {
// 	return price * 0.85
// }

// func Calculate(p float64, d Discount) float64 {
// 	return d.Apply(p)
// }
// type Flyer interface {
// 	Fly()
// }

// type Walker interface {
// 	Walk()
// }
// type Bird interface {
// 	Fly()
// }

// type Sparrow struct{}
// func (s Sparrow) Fly() {
// 	fmt.Println("Sparrow flying")
// }

// type Ostrich struct{}
// func (o Ostrich) Fly() {
// 	// Ostrich can't fly!
// 	panic("Ostrich can't fly!")
// }

// type Printer interface {
// 	Print()
// }
// type Scanner interface {
// 	Scan()
// }
// type Fax interface {
// 	Scan()
// }
// type Machine interface {
// 	Print()
// 	Scan()
// 	Fax()
// }

// type SimplePrinter struct{}
// func (p SimplePrinter) Print() { fmt.Println("Printed") }


type Notifier interface {
	Send(to, msg string)
}

type Email struct{}
func (e Email) Send(to, msg string) {
	fmt.Println("Email:", to, msg)
}

type UserService struct {
	notifier Notifier
}

func NewUserService(n Notifier) *UserService {
	return &UserService{notifier: n}
}

func (u *UserService) Register(name, email string) {
	u.notifier.Send(email, "Welcome, "+name)
}