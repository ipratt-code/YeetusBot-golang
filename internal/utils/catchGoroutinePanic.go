package utils

import "fmt"

func CatchGoroutinePanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovering from panic:", r)
	}
}