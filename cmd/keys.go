/*
Copyright © 2020 Georgios Delkos georgios.delkos@certik.io

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"crypto/rand"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ed25519"
)

// keysCmd represents the keys command
var keysCmd = &cobra.Command{
	Use:   "keys",
	Short: "Keys will return a ed25519 keypair",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		getKeyPair(args)
	},
}

func init() {
	rootCmd.AddCommand(keysCmd)

}

func getKeyPair(args []string) {
	pub, priv, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Println("Error generating keypair")
		os.Exit(1)
	}
	fmt.Printf("Public Key: %X \n", pub)
	fmt.Printf("Private Key: %X \n", priv)
}
