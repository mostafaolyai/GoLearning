package main

import "fmt"

type ConfigKeyNotFoundError struct {
	Key string
}

func (e *ConfigKeyNotFoundError) Error() string {
	return fmt.Sprintf("config key not found: %s", e.Key)
}