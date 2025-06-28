//go:build dev
// +build dev

package main

func LoadConfig() map[string]interface{} {
	return map[string]interface{}{
		"APP_NAME": "MyApp (Dev)",
		"PORT":     3000,
		"DEBUG":    true,
	}
}
