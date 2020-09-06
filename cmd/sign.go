package cmd

import (
	"avalon/utils"
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
	utils.Check(err, "Error: trying to read private key file")

	input := args[1]
	fi, err := os.Stat(input)
	utils.Check(err, "Error: trying to parse the file or directory name")

	var bytecode []byte
	switch mode := fi.Mode(); {
	case mode.IsDir():
		fmt.Println("Directory")
		fmt.Println(input)
		dir, errd := ioutil.ReadDir(input)
		utils.Check(errd, "Error: trying to read directory")
		for _, file := range dir {
			b, errb := ioutil.ReadFile(file.Name())
			utils.Check(errb, "Error: trying to read directory files")
			bytecode = append(bytecode[:], b...)
		}
	case mode.IsRegular():
		fmt.Println("File")
		bytecode, err = ioutil.ReadFile(input)
		utils.Check(err, "Error: trying to read file to sign")
	}

	sig := ed25519.Sign(privKey, bytecode)
	fmt.Printf("Private Key: %x \n", privKey)
	fmt.Printf("Public Key: %x \n", privKey.Public())
	fmt.Printf("Signature: %x \n", sig)

	err = ioutil.WriteFile("tests/sig.sec", sig, 0644)
	utils.Check(err, "Error:  trying to write sig file")
	fmt.Println("Done with signing the data")
}
