package cmd

import (
	"avalon/params"
	"avalon/utils"
	"crypto/ed25519"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve starts a service given a port number",
	Long:  `Serve starts a service given a port number that under /verify can verify signatures against data and public key`,
	Run: func(cmd *cobra.Command, args []string) {
		serve(args)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve(args []string) {
	if len(args) != 1 {
		utils.PrintItems("error", "Please provide a argument for port")
		return
	}
	app := fiber.New()

	app.Post("/verify", func(c *fiber.Ctx) {
		if form, err := c.MultipartForm(); err == nil {
			if form.File["pubkey"] == nil || form.File["sig"] == nil || form.File["data"] == nil {
				utils.PrintItems("error", "Please provide all the data needed to procced, public key, data and the signature to check against")
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
				c.Send("Failed to verify the data provided, please check the inputs")
				time.Sleep(1000 * time.Millisecond)
				return
			}
			b := ed25519.Verify(publicKey, data, signature)
			c.Send("Is the signature valid: ", b)
		}

	})

	port, err := strconv.Atoi(args[0])
	utils.Check(err, "Error: trying to parse port")
	log.Fatal(app.Listen(port))
}

// TODO
func verifyJSON(pub, data, sig []byte) bool {
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
