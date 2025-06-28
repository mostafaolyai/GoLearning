package main

import "fmt"

type Role int

const (
	Guest Role = iota
	User
	Admin
)

func (r Role) String() string {
	return [...]string{"Guest", "User", "Admin"}[r]
}

func PrintRole(role Role) {
	fmt.Println("🎖 Role:", role)
}