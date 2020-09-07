package cmd

import (
	"avalon/params"
	"avalon/utils"
	"crypto/ed25519"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/gofiber/fiber"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		serve(args)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type payload struct {
	PubKey    []byte `json:"pubkey"`
	Data      []byte `json:"data"`
	Signature []byte `json:"sig"`
}

func serve(args []string) {
	if len(args) == 0 {
		fmt.Println("Please provide a argument for port")
		return
	}
	app := fiber.New()

	app.Post("/verify", func(c *fiber.Ctx) {
		if form, err := c.MultipartForm(); err == nil {
			if form.File["pubkey"] == nil || form.File["sig"] == nil || form.File["data"] == nil {
				fmt.Println("Please provide all the data needed to procced")
				return
			}
			publicKeyHeader := form.File["pubkey"]
			signatureHeader := form.File["sig"]
			files := form.File["data"]

			data := make([]byte, 0)
			for _, file := range files {
				b, err := file.Open()
				utils.Check(err, "Error: trying to open the uploaded file")
				x, err := ioutil.ReadAll(b)
				utils.Check(err, "Error: trying to read the uploaded file")
				data = append(data[:], x...)
			}

			publicKeyFile, err := publicKeyHeader[0].Open()
			utils.Check(err, "Error: trying to open the public key uploaded file")
			publicKey, err := ioutil.ReadAll(publicKeyFile)
			utils.Check(err, "Error: trying to read the public key uploaded file")
			signatureFile, err := signatureHeader[0].Open()
			utils.Check(err, "Error: trying to open the signature uploaded file")
			signature, err := ioutil.ReadAll(signatureFile)
			utils.Check(err, "Error: trying to read the signature uploaded file")
			if !verifyJSON(publicKey, data, signature) {
				return
			}
			b := ed25519.Verify(publicKey, data, signature)
			fmt.Println("Is the signature valid: ", b)
			c.Send("Is the signature valid: ", b)
		}

	})

	port, err := strconv.Atoi(args[0])
	utils.Check(err, "Error: trying to parse port")
	log.Fatal(app.Listen(port))
}

// TODO
func verifyJSON(pub, data, sig []byte) bool {
	fmt.Println(len(pub), len(data), len(sig))
	if len(pub) != 32 {
		fmt.Println("Public Key is not correct")
		return false
	}
	if len(data) == 0 {
		fmt.Println("No Data to process")
		return false
	}
	if len(data) > params.MaxDataSize {
		fmt.Println("Data is to big to verify")
		return false
	}
	if len(sig) != 64 {
		fmt.Println("Signature is not correct")
		return false
	}
	return true
}
