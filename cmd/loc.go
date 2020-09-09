package cmd

import (
	"avalon/utils"
	"os/exec"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

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

}

func getLoc(args []string) {
	if len(args) != 1 {
		color.Red("Error: please provide a name for the keys")
		return
	}
	printLocIntro()
	input := args[0]
	command := exec.Command("bash", "-c", "/usr/bin/find "+input+" -name '*.sol' | xargs wc -l | sort -nr")
	stdout, err := command.Output()
	utils.Check(err, "Error: trying to run command")
	printLocOutro(string(stdout))
}

func printLocIntro() {
	color.Green("Welcome to Avalon lines of code counter")
	color.Green("Counting lines of code...")
}

func printLocOutro(msg string) {
	color.Green("Printing out loc for sol files in directory")
	color.HiBlue(string(msg))
}
