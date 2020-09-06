package cmd

import (
	"crypto/ed25519"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		verify(args)
	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)
}

func verify(args []string) {
	if len(args) < 2 {
		fmt.Println("Error: Please provide some arguments")
		os.Exit(1)
	}
	var pubKey ed25519.PublicKey
	var err error
	pubKey, err = ioutil.ReadFile(args[0])
	if err != nil {
		fmt.Println("Error: Cant read public key file")
		os.Exit(1)
	}
	fmt.Printf("Public Key: %x \n", pubKey)
	msg, err := ioutil.ReadFile(args[1])
	if err != nil {
		fmt.Println("Error: Cant read data file")
		os.Exit(1)
	}
	sig, err := ioutil.ReadFile(args[2])
	if err != nil {
		fmt.Println("Error: Cant read signature file")
		os.Exit(1)
	}
	fmt.Printf("Signature: %x \n", sig)
	out := ed25519.Verify(pubKey, msg, sig)
	fmt.Println("Is this signature valid? -> ", out)
}
