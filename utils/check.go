package utils

import (
	"os"
	"strings"

	"github.com/fatih/color"
)

// Check represents our exit function
func Check(err error, msg string) {
	if err == nil {
		return
	}
	if strings.Contains(msg, "Error:") {
		color.Yellow(msg)
		color.Red(err.Error())
		os.Exit(1)
	}
	color.Cyan(msg)
	color.Red(err.Error())
	os.Exit(1)
}
