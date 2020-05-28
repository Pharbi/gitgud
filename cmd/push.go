package cmd

import (
	. "fmt"
	"strings"
	"os/exec"
	"github.com/spf13/cobra"
)

func init(){
	rootCmd.AddCommand(pushCmd)
}

var pushCmd = &cobra.Command{
	Use: "push",
	Short: "Push from current git origin/$branchName",
	Long:"Runs git branch and retrieves the current branch, then runs git push origin $branchName",
	Run: func(cmd *cobra.Command, args []string){
		Println("Currently running pull")
		var shellCmd = exec.Command("git", "branch")
		var stdout, err = shellCmd.Output()

		if err != nil {
			Println(err.Error())
			return
		}
		outTxt := strings.Split(string(stdout), "*")
		_, branch := outTxt[0], outTxt[1]
		Println(branch)

		Printf("Currently pushing to %s", branch)
		shellCmd = exec.Command("git", "push", "origin", branch)
		stdout, err = shellCmd.Output()
		if err != nil {
			Println(err.Error())
			return
		}
		Println("Git response: ", string(stdout))
	},
}