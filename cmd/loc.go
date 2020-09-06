package cmd

import (
	"avalon/utils"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// keysCmd represents the keys command
var locCmd = &cobra.Command{
	Use:   "loc",
	Short: "Loc will return lines of code of the codebase",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		getLoc(args)
	},
}

func init() {
	rootCmd.AddCommand(locCmd)

}

func getLoc(args []string) {
	if len(args) == 0 {
		fmt.Println("Error: please provide a name for the keys")
		return
	}
	input := args[0]
	command := exec.Command("bash", "-c", "/usr/bin/find "+input+" -name '*.sol' | xargs wc -l | sort -nr")
	stdout, err := command.Output()
	utils.Check(err, "Error: trying to run command")
	fmt.Println("Printing out loc for sol file or files in directory")
	fmt.Println(string(stdout))
}
