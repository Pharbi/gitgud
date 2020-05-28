package cmd

import (
	. "fmt"
	"github.com/spf13/cobra"

	"os/exec"
)

func init(){
	rootCmd.AddCommand(pullCmd)
}

var pullCmd = &cobra.Command{
	Use: "pull",
	Short: "Pull from current git origin/$branchName",
	Long:"Runs git branch -v and retrieves the current branch, then runs git pull origin $branchName",
	Run: func(cmd *cobra.Command, args []string){
		Println("Currently running pull")
		shellCmd := exec.Command("git", "branch", "-v")
		stdout, err := shellCmd.Output()

		if err != nil {
			Println(err.Error())
			return
		}
		outTxt := string(stdout)
		Println(outTxt)
	},
}
