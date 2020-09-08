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

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "Verify a signature against a public key and data",
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
		color.Red("Error: Please provide some arguments")
		return
	}
	printVerifyIntro()
	var pubKey ed25519.PublicKey
	var err error
	pubKey, err = ioutil.ReadFile(args[0])
	utils.Check(err, "Error: Cant read public key file")

	color.HiBlue("Public Key: %x \n", pubKey)
	input := args[1]
	fi, err := os.Stat(input)
	utils.Check(err, "Error: trying to parse the file or directory name")

	var msg []byte
	switch mode := fi.Mode(); {
	case mode.IsDir():
		color.HiYellow("Info: Found directory")
		dir, errd := ioutil.ReadDir(input)
		utils.Check(errd, "Error: trying to read directory")
		for _, file := range dir {
			if strings.Contains(file.Name(), ".sol") || strings.Contains(file.Name(), ".pdf") {
				color.HiYellow("Info: Found sol file: ", input+file.Name())
				b, errb := ioutil.ReadFile(input + file.Name())
				utils.Check(errb, "Error: trying to read from files in the directory")
				msg = append(msg[:], b...)
			}
		}
	case mode.IsRegular():
		color.HiYellow("Info: Found single file")
		msg, err = ioutil.ReadFile(input)
		utils.Check(err, "Error: trying to read file to sign")
	}

	sig, err := ioutil.ReadFile(args[2])
	utils.Check(err, "Error: Cant read signature file")
	out := ed25519.Verify(pubKey, msg, sig)
	printVerifyOutro(sig, out)
}

func printVerifyIntro() {
	color.Green("Welcome to Avalon Verify tool")
	color.Green("Verifying signature against data with given public key...")
}

func printVerifyOutro(sig []byte, valid bool) {
	color.HiBlue("Signature: %x \n", sig)
	if valid {
		color.HiGreen("Result: Success on validating this signature against data and public key")
		return
	}
	color.Red("Result: Failure to validate the signature against data and public key")
}
