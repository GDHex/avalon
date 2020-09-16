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
	if len(args) != 2 {
		utils.PrintItems("error", "Please provide two arguments, the private key file and the data to sign")
		return
	}
	printSignIntro()
	var privKey ed25519.PrivateKey
	var err error
	privKey, err = ioutil.ReadFile(args[0])
	utils.Check(err, "Error: Trying to read private key file")
	utils.ValidatePrivateKey(privKey)

	input := args[1]
	// err = utils.ValidateFileInput(input)
	// utils.Check(err, err.Error())

	fi, err := os.Stat(input)
	utils.Check(err, "Error: Trying to parse the file or directory name")

	var bytecode []byte
	switch mode := fi.Mode(); {
	case mode.IsDir():
		utils.PrintItems("info", "Found directory")
		dir, errd := ioutil.ReadDir(input)
		utils.Check(errd, "Error: trying to read directory")
		for _, file := range dir {
			if strings.Contains(file.Name(), params.Sol) || strings.Contains(file.Name(), params.Pdf) { // TODO add types here
				utils.PrintItems("info", "Found sol file: "+file.Name())
				b, errb := ioutil.ReadFile(input + file.Name())
				utils.Check(errb, "Error: trying to read from files in the directory")
				bytecode = append(bytecode, b...)
			}
		}
	case mode.IsRegular():
		utils.PrintItems("info", "Found single file")
		bytecode, err = ioutil.ReadFile(input)
		utils.Check(err, "Error: trying to read file to sign")
	}

	sig := ed25519.Sign(privKey, bytecode)
	name := strings.Split(input, params.Pdf)
	s := strings.TrimPrefix(name[0], params.DirPrefix) // Hack for now

	err = ioutil.WriteFile(params.SignatureDir+s+params.SignatureSuffix, sig, 0644)
	utils.Check(err, "Error:  trying to write sig file")

	printSignOutro(len(bytecode), sig)
}

func printSignIntro() {
	utils.PrintItems("line", "---------------------------------------------------------------------------------")
	utils.PrintItems("action", "                        Welcome to Avalon Sign tool")
	utils.PrintItems("line", "---------------------------------------------------------------------------------")
	utils.PrintItems("action", "Signing data with given private key...")
}

func printSignOutro(lenght int, sig []byte) {
	color.HiBlue("Data lenght: %d", lenght)
	color.HiCyan("Signature: %x \n", sig)
	utils.PrintItems("line", "---------------------------------------------------------------------------------")
	utils.PrintItems("success", "Done with signing the data")
	utils.PrintItems("line", "---------------------------------------------------------------------------------")
}
