package cmd

import (
	. "gitgud/utils"
	"github.com/go-git/go-git/v5"
	"github.com/spf13/cobra"
	"os"
)

var pullCmd = &cobra.Command{
	Use: "pull",
	Short: "Pull from current git origin/$branchName",
	Long:"Runs git branch and retrieves the current branch, then runs git pull origin $branchName",
}

func init(){
	pullCmd.Run = pull
	rootCmd.AddCommand(pullCmd)
}

func pull(cmd *cobra.Command, args []string){
	println("Currently running pull")
	println("----------------------")
	cwd, _ := os.Getwd()
	path := cwd

	repo, err := git.PlainOpen(path)
	CheckIfError(err)

	currDir, err := repo.Worktree()
	CheckIfError(err)

	auth := GetKeys()
	err = currDir.Pull(&git.PullOptions{
		RemoteName: "origin",
		Auth: auth})
	CheckIfError(err)

	ref, err := repo.Head()
	CheckIfError(err)

	lastCommit, err := repo.CommitObject(ref.Hash())
	CheckIfError(err)
	println("Done! Currently on this commit:", lastCommit)
}
