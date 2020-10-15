package cmd

import (
	"avalon/utils"
	"crypto/ed25519"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

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
			if form.File["data"] == nil {
				utils.PrintItems("error", "Please provide data needed to procced")
				return
			}
			files := form.File["data"]
			filename := ""
			data := make([]byte, 0)
			for _, file := range files {
				b, err := file.Open()
				utils.Check(err, "Error: trying to open the uploaded file")
				filename = file.Filename
				x, err := ioutil.ReadAll(b)
				utils.Check(err, "Error: trying to read the uploaded file")
				data = append(data, x...)
			}

			keyDirTag := "./keys/"
			// TODO verify again data
			publicKeyFiles, err := ioutil.ReadDir(keyDirTag)
			utils.Check(err, "Error: trying to read the keys directory")

			sigDirTag := "./signatures/"
			signatureFiles, err := ioutil.ReadDir(sigDirTag)
			utils.Check(err, "Error: trying to read the signatures directory")
			isFound := false
		Loop:
			for _, sigFile := range signatureFiles {
				filenameTrimmed := filename
				filenameTrimmed = filename[:len(filename)-4]
				sigNameTrimmed := sigFile.Name()
				sigNameTrimmed = sigNameTrimmed[:len(sigNameTrimmed)-8]
				fmt.Println("Signame Trimmed", sigNameTrimmed)
				if strings.Contains(sigNameTrimmed, filenameTrimmed) {
					for _, pubKFile := range publicKeyFiles {

						pubK, err := ioutil.ReadFile(keyDirTag + pubKFile.Name())
						utils.Check(err, "Error: trying to read the public key file")
						sig, err := ioutil.ReadFile(sigDirTag + sigFile.Name())
						utils.Check(err, "Error: trying to read the signature key file")
						b := ed25519.Verify(pubK, data, sig)
						if b {
							author := pubKFile.Name()
							author = author[:len(author)-9]
							fmt.Println("Is this file valid and signed by ", author)
							c.Send("Is this file valid and signed by ", author)
							isFound = true
							break Loop
						}
					}
				}
			}
			if !isFound {
				c.Send("Unable to verify.There is no record of this report!")
				fmt.Println("Unable to verify.There is no record of this report!")
			}
		} else {
			fmt.Println("Error: trying to parse the form")
		}
	})
	port, err := strconv.Atoi(args[0])
	utils.Check(err, "Error: trying to parse port")
	log.Fatal(app.Listen(port))
}
