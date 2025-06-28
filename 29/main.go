package main

import (
	"fmt"
	"reflect"
)

func init() {
	fmt.Println("🔥 Config Manager Starting...")
}

func PrintTypeAndValue[T any](key string, val T) {
	fmt.Printf("🔑 Key: %s\n", key)
	fmt.Println("📦 Type:", reflect.TypeOf(val))
	fmt.Println("📄 Value:", reflect.ValueOf(val))
}

type Supported interface {
	string | int | bool
}

func GetConfigValue[T Supported](cfg map[string]interface{}, key string) (T, error) {
	val, ok := cfg[key]
	if !ok {
		return *new(T), &ConfigKeyNotFoundError{Key: key}
	}
	typedVal, ok := val.(T)
	if !ok {
		return *new(T), fmt.Errorf("type mismatch for key %s: expected %T but got %T", key, typedVal, val)
	}
	return typedVal, nil
}


func main() {
	cfg := LoadConfig()

	// Print all config values with type
	for k, v := range cfg {
		PrintTypeAndValue(k, v)
	}

	// Use generic getter
	value, err := GetConfigValue[string](cfg, "APP_NAM")
	if err != nil {
		fmt.Println("❌", err)
	} else {
		fmt.Println("✅ App Name:", value)
	}
}