package cmd

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ed25519"
)

// keysCmd represents the keys command
var keysCmd = &cobra.Command{
	Use:   "gen-keys",
	Short: "Gen-keys will return a ed25519 keypair",
	Long: `Gen-keys returns a new keypair and stores it into 2 files, argument is the name 
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
		os.Exit(1)
	}
	name := args[0]
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Println("Error: generating keypair")
		os.Exit(1)
	}
	err = ioutil.WriteFile("tests/"+name+"_pblk.sec", pub, 0644)
	if err != nil {
		fmt.Println("Error: trying to write public key file")
		os.Exit(1)
	}
	err = ioutil.WriteFile("tests/"+name+"_prvk.sec", priv, 0644)
	if err != nil {
		fmt.Println("Error: trying to write private key file")
		os.Exit(1)
	}
	fmt.Println("Done! Keys are saved for user: ", name)
	fmt.Printf("Public Key: %x \n", pub)
	fmt.Printf("Private Key: %x \n", priv)
}
