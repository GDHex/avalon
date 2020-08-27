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
	"github.com/vuvuzela/crypto/bls"
)

// signCmd represents the sign command
var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "Sign a collection of data with BLS signature",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		sign(args)
	},
}

func init() {
	rootCmd.AddCommand(signCmd)
}

func sign(args []string) {
	addr := []byte(args[0])
	bytecode := make([]byte, 64)
	combine := append(addr[:], bytecode[:]...)
	pub, priv, err := bls.GenerateKey(rand.Reader)
	if err != nil {
		fmt.Println("Error generating new key pair")
		os.Exit(1)
	}
	fmt.Printf("Public Key: %v \n", pub)
	fmt.Printf("Private Key: %p \n", priv)
	sig := bls.Sign(priv, combine)
	fmt.Printf("This is the produced signature: %x", sig)
}
