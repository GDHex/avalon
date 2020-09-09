package cmd

import (
	"avalon/params"
	"avalon/utils"
	"io/ioutil"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/crypto/ed25519"
)

var signBool bool
var verifyBool bool

// batchCmd represents the batch command
var batchCmd = &cobra.Command{
	Use:   "batch",
	Short: "Batch will batch sign or verify multiple data/signatures",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if signBool {
			batchSign(args)
		}
		if verifyBool {
			color.HiYellow("Batch verify")
		}
	},
}

func init() {
	rootCmd.AddCommand(batchCmd)

	batchCmd.Flags().BoolVarP(&signBool, "sign", "s", false, "Batch sign files")
	batchCmd.Flags().BoolVarP(&verifyBool, "verify", "v", false, "Batch verify files")
	viper.BindPFlag("sign", rootCmd.PersistentFlags().Lookup("sign"))
	viper.BindPFlag("verify", rootCmd.PersistentFlags().Lookup("verify"))
}

func batchSign(args []string) {
	if len(args) != 2 {
		utils.PrintItems("error", "Please provide two arguments, the private key file and the directory")
		return
	}
	utils.PrintItems("action", "Starting up batch signing process...")
	var privKey ed25519.PrivateKey
	var err error
	privKey, err = ioutil.ReadFile(args[0])
	utils.Check(err, "Error: trying to read private key file")
	utils.ValidatePrivateKey(privKey)

	input := args[1]
	fi, err := os.Stat(input)
	utils.Check(err, "Error: trying to parse the directory name")

	switch mode := fi.Mode(); {
	case mode.IsRegular():
		utils.PrintItems("error", "You need to provide with a directory to read")
		return
	case mode.IsDir():
		utils.PrintItems("info", "Found directory")
		dir, errd := ioutil.ReadDir(input)
		utils.Check(errd, "Error: trying to read directory")
		for _, file := range dir {
			if strings.Contains(file.Name(), params.Pdf) { // TODO add types here
				utils.PrintItems("info", "Found pdf file: "+file.Name())
				b, errb := ioutil.ReadFile(input + file.Name())
				utils.Check(errb, "Error: trying to read from files in the directory")
				sig := ed25519.Sign(privKey, b)
				name := strings.Split(file.Name(), params.Pdf)

				s := strings.TrimPrefix(name[0], params.DirPrefix) // Hack for now
				err = ioutil.WriteFile(params.SignatureDir+s+params.SignatureSuffix, sig, 0644)
				utils.Check(err, "Error:  trying to write sig file")
				printBatchSignOutro(file.Size(), sig)
			}
		}
	}
}

func printBatchSignOutro(lenght int64, sig []byte) {
	utils.PrintItems("line", "---------------------------------------------------------------------------------")
	utils.PrintItems("success", "Done with signing the data")
	color.HiBlue("Data lenght: %d", lenght)
	color.HiCyan("Signature: %x \n", sig)
	utils.PrintItems("line", "---------------------------------------------------------------------------------")
}
