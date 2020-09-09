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
	if len(args) != 2 {
		color.Red("Error: Please provide two arguments, the private key file and the data to sign")
		os.Exit(1)
	}
	printSignIntro()
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
		color.HiYellow("Info: Found directory")
		dir, errd := ioutil.ReadDir(input)
		utils.Check(errd, "Error: trying to read directory")
		for _, file := range dir {
			if strings.Contains(file.Name(), ".sol") || strings.Contains(file.Name(), ".pdf") { // TODO add types here
				color.HiYellow("Info: Found sol file: ", input+file.Name())
				b, errb := ioutil.ReadFile(input + file.Name())
				utils.Check(errb, "Error: trying to read from files in the directory")
				bytecode = append(bytecode[:], b...)
			}
		}
	case mode.IsRegular():
		color.HiYellow("Info: Found single file")
		bytecode, err = ioutil.ReadFile(input)
		utils.Check(err, "Error: trying to read file to sign")
	}

	sig := ed25519.Sign(privKey, bytecode)
	name := strings.Split(input, ".pdf")
	s := strings.TrimPrefix(name[0], "./data/") // Hack for now

	err = ioutil.WriteFile("./signatures/"+s+"_sig.sec", sig, 0644)
	utils.Check(err, "Error:  trying to write sig file")
	color.Green("Done with signing the data")
	printSignOutro(sig)
}

func printSignIntro() {
	color.Green("Welcome to Avalon Sign tool")
	color.Green("Signing data with given private key...")
}

func printSignOutro(sig []byte) {
	color.HiBlue("Signature: %x \n", sig)
}
