package cmd

import (
	"avalon/utils"
	"crypto/rand"
	"fmt"
	"io/ioutil"

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
		getKeyPair(args)
	},
}

func init() {
	rootCmd.AddCommand(keysCmd)

}

func getKeyPair(args []string) {
	if len(args) == 0 {
		fmt.Println("Error: please provide a name for the keys")
		return
	}
	name := args[0]
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	utils.Check(err, "Error: generating keypair")

	err = ioutil.WriteFile("keys/"+name+"_pblk.sec", pub, 0644)
	utils.Check(err, "Error: trying to write public key file")

	err = ioutil.WriteFile("keys/"+name+"_prvk.sec", priv, 0644)
	utils.Check(err, "Error: trying to write private key file")

	fmt.Println("Done! Keys are saved for user: ", name)
	fmt.Printf("Public Key: %x \n", pub)
	fmt.Printf("Private Key: %x \n", priv)
}
