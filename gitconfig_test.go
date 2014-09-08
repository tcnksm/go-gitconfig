package gitconfig

import (
	"fmt"
	. "github.com/onsi/gomega"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestGlobal(t *testing.T) {
	RegisterTestingT(t)

	var (
		err error
	)

	username, err := Global("user.name")
	Expect(err).NotTo(HaveOccurred())
	fmt.Println("user.name: ", username)
}

func TestLocal(t *testing.T) {
	RegisterTestingT(t)

	var (
		err error
	)

	url, err := Local("remotes.origin.url")
	Expect(err).NotTo(HaveOccurred())
	fmt.Println("remotes.origin.url: ", url)
}

func TestExecGitConfig(t *testing.T) {
	RegisterTestingT(t)

	reset := withGitConfigFile(`
[user]
    name  = deeeet
    email = deeeet@example.com
`)

	defer reset()

	var (
		err error
	)

	username, err := execGitConfig("user.name")
	Expect(err).NotTo(HaveOccurred())
	Expect(username).To(Equal("deeeet"))

	email, err := execGitConfig("user.email")
	Expect(err).NotTo(HaveOccurred())
	Expect(email).To(Equal("deeeet@example.com"))
}

func withGitConfigFile(content string) func() {
	tmpdir, err := ioutil.TempDir("", "go-gitconfig-test")
	if err != nil {
		panic(err)
	}

	tmpGitConfigFile := filepath.Join(tmpdir, "gitconfig")

	ioutil.WriteFile(
		tmpGitConfigFile,
		[]byte(content),
		0777,
	)

	prevGitConfigEnv := os.Getenv("GIT_CONFIG")
	os.Setenv("GIT_CONFIG", tmpGitConfigFile)

	return func() {
		os.Setenv("GIT_CONFIG", prevGitConfigEnv)
	}
}
