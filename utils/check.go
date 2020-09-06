package utils

import (
	"fmt"
	"os"
)

// Check represents our exit function
func Check(err error, msg string) {
	if err == nil {
		return
	}
	fmt.Println(msg)
	os.Exit(1)
}
