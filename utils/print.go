package utils

import "github.com/fatih/color"

// PrintItems prints the message with a coresponding color match
func PrintItems(itemType string, msg string) {
	switch {
	case itemType == "info":
		color.HiMagenta("Info: " + msg)
	case itemType == "success":
		color.HiGreen("Success: " + msg)
	case itemType == "error":
		color.HiRed("Error: " + msg)
	case itemType == "action":
		color.Green(msg)
	case itemType == "data":
		color.HiCyan(msg)
	case itemType == "line":
		color.Green(msg)
	}
}
