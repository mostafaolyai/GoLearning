//go:build prod
// +build prod

package main

func GetEnvironment() string {
	return "Production"
}
