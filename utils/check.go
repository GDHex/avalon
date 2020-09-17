package utils

import (
	"avalon/params"
	"errors"
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

// ValidatePublicKey validates a public key
func ValidatePublicKey(key []byte) {
	if len(key) != 32 {
		PrintItems("error", "Public key lenght is not correct")
		os.Exit(1)
	}
}

// ValidatePrivateKey validates a private key
func ValidatePrivateKey(key []byte) {
	if len(key) != 64 {
		PrintItems("error", "Private key lenght is not correct")
		os.Exit(1)
	}
}

// ValidateSignature validates a signature
func ValidateSignature(sig []byte) {
	if len(sig) != 64 {
		PrintItems("error", "Signature lenght is not correct")
		os.Exit(1)
	}
}

// ValidateFileInput validates a file input
func ValidateFileInput(src string) error {
	err := errors.New("Error: Filetype not supported")
	for _, ft := range params.FileTypes {
		if strings.Contains(src, ft) {
			return nil
		}
	}
	return err
}
