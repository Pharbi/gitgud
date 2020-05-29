package utils

import (
	. "fmt"
	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/mitchellh/go-homedir"
	"io/ioutil"
	"os"
	"path/filepath"
)

func GetKeys() *ssh.PublicKeys{
	println("Retrieving the keys")
	println("Expanding the directory path")
	println("----------------------------")
	home, _ := homedir.Dir()
	path := filepath.Join(home, "/.ssh/id_rsa")
	sshKey, err := ioutil.ReadFile(path)
	CheckIfError(err)
	auth, err := ssh.NewPublicKeys("git", sshKey, "")
	CheckIfError(err)
	return auth
}

func CheckIfError(err error) {
	if err == nil {
		return
	}

	Printf("\x1b[31;1m%s\x1b[0m\n", Sprintf("Git err response on pull: %s", err))
	os.Exit(1)
}