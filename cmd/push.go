package cmd

import (
	. "gitgud/utils"
	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
	"os"
)

var pushCmd = &cobra.Command{
	Use: "push",
	Short: "Push from current git origin/$branchName",
	Long:"Runs git branch and retrieves the current branch, then runs git push origin $branchName",
}

func init(){
	pushCmd.Run = push
	rootCmd.AddCommand(pushCmd)
}

func push(cmd *cobra.Command, args []string) {
	println("Currently running push")
	println("----------------------")
	cwd, _ := os.Getwd()
	path := cwd

	repo, err := git.PlainOpen(path)
	CheckIfError(err)

	auth := GetKeys()
	err = repo.Push(&git.PushOptions{
		RemoteName: "origin",
		Auth: auth})
	CheckIfError(err)
	println("Done!")
}