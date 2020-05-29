package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: "gitgud",
	Short: "Less git typing",
	Long: "A way to get current branches and affix them to git commands to minimize on typing",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		er(err)
		os.Exit(1)
	}
}

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(1)
}
