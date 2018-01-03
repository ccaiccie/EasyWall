package main

import (
	"fmt"
	"time"
)

// Log - Write Log into Console and Log File
func Log(message string) {
	fmt.Println(time.Now().Format(time.RFC3339), message)
}
