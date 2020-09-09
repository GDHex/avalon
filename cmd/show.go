package cmd

import (
	"avalon/params"
	"avalon/utils"
	"crypto/ed25519"
	"io/ioutil"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show will load private and public key for a username given",
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
		utils.PrintItems("error", "Please provide at least one argument, the name of the user")
		return
	}
	printShowIntro(args[0])
	pubKey, err := ioutil.ReadFile(params.KeyDir + args[0] + params.PublicKeySuffix)
	utils.Check(err, "Error: trying to read the public key file")
	privKey, errp := ioutil.ReadFile(params.KeyDir + args[0] + params.PrivateKeySuffix)
	utils.Check(errp, "Error: trying to read the public key file")

	printShowOutro(pubKey, privKey)
}

func printShowIntro(user string) {
	utils.PrintItems("line", "---------------------------------------------------------------------------------")
	color.Green("                                 Welcome to Avalon keypair show tool")
	utils.PrintItems("line", "---------------------------------------------------------------------------------")
	color.Green("Selected user : " + user)
	color.Green("Loading keypair from files...")
}

func printShowOutro(pub ed25519.PublicKey, priv ed25519.PrivateKey) {
	utils.PrintItems("line", "---------------------------------------------------------------------------------")
	color.HiBlue("Public Key: %x \n", pub)
	color.HiBlue("Private Key: %x \n", priv)
	utils.PrintItems("line", "---------------------------------------------------------------------------------")
}
