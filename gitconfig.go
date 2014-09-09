// Package gitconfig enables you to use `~/.gitconfig` values in Golang.
//
// For a full guide visit http://github.com/tcnksm/go-gitconfig
//
//  package main
//
//  import (
//    "github.com/tcnksm/go-gitconfig"
//    "fmt"
//  )
//
//  func main() {
//    user, err := gitconfig.Global("user.name")
//    if err == nil {
//      fmt.Println(user)
//    }
//  }
//
package gitconfig

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
	"syscall"
)

// Global extracts configuration value from `$HOME/.gitconfig` file or `$GIT_CONFIG`.
func Global(key string) (string, error) {
	return execGitConfig("--global", key)
}

// Local extracts configuration value from current project repository.
func Local(key string) (string, error) {
	return execGitConfig("--local", key)
}

// Username extracts git user name from `$HOME/.gitconfig` file or `$GIT_CONFIG`.
// This is same as Global("user.name")
func Username() (string, error) {
	return Global("user.name")
}

// Email extracts git user email from `$HOME/.gitconfig` file or `$GIT_CONFIG`.
// This is same as Global("user.email")
func Email() (string, error) {
	return Global("user.email")
}

// OriginURL extract remote origin url from current project repository.
// This is same as Local("remote.origin.url")
func OriginURL() (string, error) {
	return Local("remote.origin.url")
}

func execGitConfig(args ...string) (string, error) {
	gitArgs := append([]string{"config", "--get", "--null"}, args...)
	var stdout bytes.Buffer
	cmd := exec.Command("git", gitArgs...)
	cmd.Stdout = &stdout
	cmd.Stderr = ioutil.Discard

	err := cmd.Run()
	if exitError, ok := err.(*exec.ExitError); ok {
		if waitStatus, ok := exitError.Sys().(syscall.WaitStatus); ok {
			if waitStatus.ExitStatus() == 1 {
				return "", fmt.Errorf("the key was not found")
			}
		}
		return "", err
	}

	return strings.TrimRight(stdout.String(), "\000"), nil
}
