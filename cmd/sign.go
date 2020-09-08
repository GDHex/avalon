package cmd

import (
	"avalon/utils"
	"crypto/ed25519"
	"io/ioutil"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// signCmd represents the sign command
var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "Create a signature from a collection of data signed with a private key",
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
		color.Red("Error: Please provide at least 2 arguments")
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
		color.Green("Found directory")
		dir, errd := ioutil.ReadDir(input)
		utils.Check(errd, "Error: trying to read directory")
		for _, file := range dir {
			if strings.Contains(file.Name(), ".sol") || strings.Contains(file.Name(), ".pdf") {
				color.Green("Found sol file: ", input+file.Name())
				b, errb := ioutil.ReadFile(input + file.Name())
				utils.Check(errb, "Error: trying to read from files in the directory")
				bytecode = append(bytecode[:], b...)
			}
		}
	case mode.IsRegular():
		color.Green("Found single file")
		bytecode, err = ioutil.ReadFile(input)
		utils.Check(err, "Error: trying to read file to sign")
	}
	printSignIntro()
	sig := ed25519.Sign(privKey, bytecode)

	err = ioutil.WriteFile("signatures/sig.sec", sig, 0644)
	utils.Check(err, "Error:  trying to write sig file")
	color.Green("Done with signing the data")
	printSignOutro(privKey, sig)
}

func printSignIntro() {
	color.Green("Welcome to Avalon Sign tool")
	color.Green("Signing data with given private key...")
}

func printSignOutro(priv ed25519.PrivateKey, sig []byte) {
	color.HiBlue("Public Key: %x \n", priv.Public())
	color.HiBlue("Private Key: %x \n", priv)
	color.HiBlue("Signature: %x \n", sig)
}
