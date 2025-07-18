// type EmailService struct{}

// func (e EmailService) SendEmail(to, message string) {
// 	fmt.Println("Sending Email to:", to, "=>", message)
// }

// type UserService struct {
// 	email EmailService
// }

// func (u UserService) RegisterUser(name, email string) {
// 	fmt.Println("Registering user:", name)
// 	u.email.SendEmail(email, "Welcome!")
// }
// ///DI

// // وابستگی به صورت interface تعریف میشه
// type Notifier interface {
// 	Send(to, message string)
// }

// // پیاده‌سازی Email
// type EmailService struct{}
// func (e EmailService) Send(to, message string) {
// 	fmt.Println("Email to", to, ":", message)
// }

// // سرویس اصلی که وابستگی از بیرون تزریق میشه
// type UserService struct {
// 	notifier Notifier
// }

// func NewUserService(n Notifier) *UserService {
// 	return &UserService{notifier: n}
// }

// func (u *UserService) RegisterUser(name, email string) {
// 	fmt.Println("Registering:", name)
// 	u.notifier.Send(email, "Welcome!")
// }