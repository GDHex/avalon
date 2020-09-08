package cmd

import (
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
	if len(args) == 0 {
		color.Red("Error: please provide one argument, a name for the keypair")
		return
	}
	printIntro()

	name := args[0]
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	utils.Check(err, "Error: generating keypair")

	if ioutil.ReadFile("keys/" + name + "_pblk.sec"); err == nil {
		color.Red("Error: File already exists, please provide a different name for the keypair")
		return
	}

	err = ioutil.WriteFile("keys/"+name+"_pblk.sec", pub, 0644)
	utils.Check(err, "Error: trying to write public key file")

	if ioutil.ReadFile("keys/" + name + "_prvk.sec"); err == nil {
		color.Red("Error: File already exists, please provide a different name for the keypair")
		return
	}

	err = ioutil.WriteFile("keys/"+name+"_prvk.sec", priv, 0644)
	utils.Check(err, "Error: trying to write private key file")

	printOutro(pub, priv)
}

func printIntro() {
	color.HiGreen("Welcome to Avalon Keypair Generator")
	color.HiGreen("Creating ed25519 keypair...")
}

func printOutro(pub ed25519.PublicKey, priv ed25519.PrivateKey) {
	color.Green("Done! Keys are saved under keys folder")
	color.Blue("Public Key: %x \n", pub)
	color.Blue("Private Key: %x \n", priv)
}
