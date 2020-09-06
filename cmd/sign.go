package cmd

import (
	"crypto/ed25519"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

// signCmd represents the sign command
var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "Sign a collection of data with a signature",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		sign(args)
	},
}

func init() {
	rootCmd.AddCommand(signCmd)
}

func sign(args []string) {
	if len(args) < 1 {
		fmt.Println("Error: Please provide at least 2 arguments")
		os.Exit(1)
	}
	var privKey ed25519.PrivateKey
	var err error
	privKey, err = ioutil.ReadFile(args[0])
	if err != nil {
		fmt.Println("Error: trying to read private key file")
		os.Exit(1)
	}
	bytecode, errb := ioutil.ReadFile(args[1])
	if errb != nil {
		fmt.Println("Error: trying to read file to sign")
		os.Exit(1)
	}

	sig := ed25519.Sign(privKey, bytecode)
	fmt.Printf("Private Key: %x \n", privKey)
	fmt.Printf("Public Key: %x \n", privKey.Public())
	fmt.Printf("Signature: %x \n", sig)

	err = ioutil.WriteFile("tests/sig.sec", sig, 0644)
	if err != nil {
		fmt.Println("Error:  trying to write sig file")
		os.Exit(1)
	}
	fmt.Println("Done with signing the data")
}
