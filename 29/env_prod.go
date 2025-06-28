//go:build prod
// +build prod

package main

func LoadConfig() map[string]interface{} {
	return map[string]interface{}{
		"APP_NAME": "MyApp (Prod)",
		"PORT":     80,
		"DEBUG":    false,
	}
}