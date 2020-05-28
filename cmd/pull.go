package cmd

import (
	. "fmt"
	"strings"
	"os/exec"
	"github.com/spf13/cobra"
)

func init(){
	rootCmd.AddCommand(pullCmd)
}

var pullCmd = &cobra.Command{
	Use: "pull",
	Short: "Pull from current git origin/$branchName",
	Long:"Runs git branch and retrieves the current branch, then runs git pull origin $branchName",
	Run: func(cmd *cobra.Command, args []string){
		Println("Currently running pull")
		var shellCmd = exec.Command("git", "branch")
		var stdout, err = shellCmd.Output()

		if err != nil {
			Println("Git err response on branch query: ", err.Error())
			return
		}
		outTxt := strings.Split(string(stdout), " ")
		_, branch := outTxt[0], outTxt[1]

		Printf("Currently pulling from %s", branch)
		shellCmd = exec.Command("git", "pull", "origin", branch)
		stdout, err = shellCmd.Output()
		if err != nil {
			Println("Git err response on pull: ", err.Error())
			return
		}
		Println("Git response: ", string(stdout))
	},
}
