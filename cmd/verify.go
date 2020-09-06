package cmd

import (
	"avalon/utils"
	"crypto/ed25519"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify a signature against public key and data",
	Long:  ``,
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
		return
	}
	var pubKey ed25519.PublicKey
	var err error
	pubKey, err = ioutil.ReadFile(args[0])
	utils.Check(err, "Error: Cant read public key file")

	fmt.Printf("Public Key: %x \n", pubKey)
	input := args[1]
	fi, err := os.Stat(input)
	utils.Check(err, "Error: trying to parse the file or directory name")

	var msg []byte
	switch mode := fi.Mode(); {
	case mode.IsDir():
		dir, errd := ioutil.ReadDir(input)
		utils.Check(errd, "Error: trying to read directory")
		for _, file := range dir {
			b, errb := ioutil.ReadFile(file.Name())
			utils.Check(errb, "Error: trying to read directory files")
			msg = append(msg[:], b...)
		}
		fmt.Println("Directory")
	case mode.IsRegular():
		fmt.Println("File")
		msg, err = ioutil.ReadFile(input)
		utils.Check(err, "Error: trying to read file to sign")
	}

	sig, err := ioutil.ReadFile(args[2])
	utils.Check(err, "Error: Cant read signature file")
	fmt.Printf("Signature: %x \n", sig)
	out := ed25519.Verify(pubKey, msg, sig)
	fmt.Println("Is this signature valid? -> ", out)
}
