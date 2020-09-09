package utils

import (
	"os"
	"strings"

	"github.com/fatih/color"
)

// Check represents our exit function with error print outs if a err is present
func Check(err error, msg string) {
	if err == nil {
		return
	}
	if strings.Contains(msg, "Error:") {
		color.Yellow(msg)
		color.Red("Inner Error: " + err.Error())
		os.Exit(1)
	}
	color.HiMagenta(msg)
	color.Red("Inner Error: " + err.Error())
	os.Exit(1)
}

func ValidatePublicKey(key []byte) {
	if len(key) != 32 {
		PrintItems("error", "Public key lenght is not correct")
		os.Exit(1)
	}
}

func ValidatePrivateKey(key []byte) {
	if len(key) != 64 {
		PrintItems("error", "Private key lenght is not correct")
		os.Exit(1)
	}
}

func ValidateSignature(sig []byte) {
	if len(sig) != 64 {
		PrintItems("error", "Signature lenght is not correct")
		os.Exit(1)
	}
}
