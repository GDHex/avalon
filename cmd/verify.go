package cmd

import (
	"avalon/params"
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
	if len(args) != 3 {
		utils.PrintItems("error", "Please provide three arguments, the public key, the data and the signature")
		return
	}
	printVerifyIntro()
	var pubKey ed25519.PublicKey
	var err error
	pubKey, err = ioutil.ReadFile(args[0])
	utils.Check(err, "Error: Cant read public key file")

	input := args[1]
	fi, err := os.Stat(input)
	utils.Check(err, "Error: Trying to parse the file or directory name")

	var data []byte
	switch mode := fi.Mode(); {
	case mode.IsDir():
		utils.PrintItems("info", "Found directory")
		dir, errd := ioutil.ReadDir(input)
		utils.Check(errd, "Error: trying to read directory")
		for _, file := range dir {
			if strings.Contains(file.Name(), params.Sol) || strings.Contains(file.Name(), params.Pdf) {
				utils.PrintItems("info", "Found sol file: "+file.Name())
				b, errb := ioutil.ReadFile(input + file.Name())
				utils.Check(errb, "Error: trying to read from files in the directory")
				data = append(data[:], b...)
			}
		}
	case mode.IsRegular():
		utils.PrintItems("info", "Found single file")
		data, err = ioutil.ReadFile(input)
		utils.Check(err, "Error: trying to read file to sign")
	}

	sig, err := ioutil.ReadFile(args[2])
	utils.Check(err, "Error: Cant read signature file")
	out := ed25519.Verify(pubKey, data, sig)
	color.HiBlue("Public Key: %x \n", pubKey)
	printVerifyOutro(len(data), sig, out)
}

func printVerifyIntro() {
	utils.PrintItems("line", "---------------------------------------------------------------------------------")
	utils.PrintItems("action", "                      Welcome to Avalon Verify tool")
	utils.PrintItems("line", "---------------------------------------------------------------------------------")
	utils.PrintItems("action", "Verifying signature against data with given public key...")
}

func printVerifyOutro(lenght int, sig []byte, valid bool) {

	color.HiBlue("Data lenght: %d", lenght)
	color.HiCyan("Signature: %x \n", sig)
	if valid {
		utils.PrintItems("line", "---------------------------------------------------------------------------------")
		color.HiGreen("Result: Success on validating this signature against data and public key")
		utils.PrintItems("line", "---------------------------------------------------------------------------------")
		return
	}
	color.Red("---------------------------------------------------------------------------------")
	color.Red("Result: Failure to validate the signature against data and public key")
	color.Red("---------------------------------------------------------------------------------")
}
