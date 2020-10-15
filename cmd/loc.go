package cmd

import (
	"avalon/utils"
	"fmt"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var fileType string

// keysCmd represents the keys command
var locCmd = &cobra.Command{
	Use:   "loc",
	Short: "Loc will return lines of code of the codebase in directory",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		getLoc(args)
	},
}

func init() {
	rootCmd.AddCommand(locCmd)
	batchCmd.Flags().StringP("filetype", "f", fileType, "file type to scan for")
	viper.BindPFlag("filetype", rootCmd.PersistentFlags().Lookup("filetype"))
}

func getLoc(args []string) {
	if len(args) != 3 {
		utils.PrintItems("error", "Please provide a directory")
		return
	}
	printLocIntro()
	fileType = args[1]
	input := args[2]

	err := utils.ValidateFileInput(fileType)
	utils.Check(err, "Error: Filetype not supported")

	command := exec.Command("bash", "-c", "/usr/bin/find "+input+" -name '*"+fileType+"' | xargs wc -l | sort -nr")
	fmt.Println(command)
	stdout, err := command.Output()
	utils.Check(err, "Error: Trying to run command")
	printLocOutro(string(stdout))
}

func printLocIntro() {
	utils.PrintItems("line", "---------------------------------------------------------------------------------")
	utils.PrintItems("action", "                Welcome to Avalon lines of code counter")
	utils.PrintItems("line", "---------------------------------------------------------------------------------")
	utils.PrintItems("action", "Counting lines of code...")
}

func printLocOutro(msg string) {
	utils.PrintItems("line", "---------------------------------------------------------------------------------")
	utils.PrintItems("action", "Printing out loc for sol files in directory")
	color.HiBlue(string(msg))
	utils.PrintItems("line", "---------------------------------------------------------------------------------")
}
