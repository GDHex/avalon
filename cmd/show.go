package cmd

import (
	"avalon/utils"
	"crypto/ed25519"
	"fmt"
	"io/ioutil"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show will load private and public key for a username given and show them in a hex format",
	Long:  `Show will take a argument of username and will return the keypair if he can find keypairs under the ./keys folder`,
	Run: func(cmd *cobra.Command, args []string) {
		show(args)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}

func show(args []string) {
	if len(args) != 1 {
		fmt.Println("Please provide at least one argument, the name of the user")
		return
	}
	printShowIntro(args[0])
	pubKey, err := ioutil.ReadFile("keys/" + args[0] + "_pblk.sec")
	utils.Check(err, "Error: trying to read the public key file")
	privKey, errp := ioutil.ReadFile("keys/" + args[0] + "_prvk.sec")
	utils.Check(errp, "Error: trying to read the public key file")

	printShowOutro(pubKey, privKey)
}

func printShowIntro(user string) {
	color.Green("Welcome to Avalon keypair show tool")
	color.Green("Selected user : " + user)
	color.Green("Loading keypair from files...")
}

func printShowOutro(pub ed25519.PublicKey, priv ed25519.PrivateKey) {
	color.HiBlue("Public Key: %x \n", pub)
	color.HiBlue("Private Key: %x \n", priv)
}
