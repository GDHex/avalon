package cmd

import (
	"avalon/params"
	"avalon/utils"
	"crypto/rand"
	"io/ioutil"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ed25519"
)

// keysCmd represents the keys command
var keysCmd = &cobra.Command{
	Use:   "gen-keys",
	Short: "Gen-keys will return a ed25519 keypair",
	Long: `Gen-keys returns a new keypair and stores it into 2 files under the keys folder, argument is the name 
	of the user and its added as a prefix for the generated files`,
	Run: func(cmd *cobra.Command, args []string) {
		newKeyPair(args)
	},
}

func init() {
	rootCmd.AddCommand(keysCmd)
}

func newKeyPair(args []string) {
	if len(args) != 1 {
		utils.PrintItems("error", "Please provide one argument, a name for the keypair")
		return
	}
	printIntro()

	name := args[0]
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	utils.Check(err, "Error: Generating keypair")

	if utils.FileExists(params.KeyDir + name + params.PublicKeySuffix) {
		utils.PrintItems("error", "File already exists, please provide a different name for the keypair")
		return
	}

	err = ioutil.WriteFile(params.KeyDir+name+params.PublicKeySuffix, pub, 0644)
	utils.Check(err, "Error: Trying to write public key file")

	if utils.FileExists(params.KeyDir + name + params.PrivateKeySuffix) {
		utils.PrintItems("error", "File already exists, please provide a different name for the keypair")
		return
	}

	err = ioutil.WriteFile(params.KeyDir+name+params.PrivateKeySuffix, priv, 0644)
	utils.Check(err, "Error: Trying to write private key file")

	printOutro(pub, priv)
}

func printIntro() {
	utils.PrintItems("line", "---------------------------------------------------------------------------------")
	utils.PrintItems("action", "                  Welcome to Avalon Keypair Generator")
	utils.PrintItems("line", "---------------------------------------------------------------------------------")
	utils.PrintItems("action", "Creating ed25519 keypair...")
}

func printOutro(pub ed25519.PublicKey, priv ed25519.PrivateKey) {
	color.Blue("Public Key: %x \n", pub)
	color.Blue("Private Key: %x \n", priv)
	utils.PrintItems("line", "---------------------------------------------------------------------------------")
	utils.PrintItems("success", "               Done! Keys are saved under keys folder")
	utils.PrintItems("line", "---------------------------------------------------------------------------------")
}
