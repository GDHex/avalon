package cmd

import (
	"avalon/utils"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

// showCmd represents the show command
var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show will load private and public key from files and show them in a hex format",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		show(args)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
}

func show(args []string) {
	if len(args) < 2 {
		fmt.Println("Please provide at least 2 arguments")
		return
	}
	pubKey, err := ioutil.ReadFile(args[0])
	utils.Check(err, "Error: trying to read the public key file")
	privKey, errp := ioutil.ReadFile(args[1])
	utils.Check(errp, "Error: trying to read the public key file")
	fmt.Printf("Public Key: %x \n", pubKey)
	fmt.Printf("Private Key: %x \n", privKey)
}
