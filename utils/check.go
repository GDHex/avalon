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
	fmt.Println(err)
	os.Exit(1)
}
