package gitconfig

import (
	. "github.com/onsi/gomega"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestGlobal(t *testing.T) {
	RegisterTestingT(t)

	reset := withGlobalGitConfigFile(`
[user]
    name  = deeeet
    email = deeeet@example.com
`)
	defer reset()

	var err error
	username, err := Global("user.name")
	Expect(err).NotTo(HaveOccurred())
	Expect(username).To(Equal("deeeet"))

	email, err := Global("user.email")
	Expect(err).NotTo(HaveOccurred())
	Expect(email).To(Equal("deeeet@example.com"))

	nothing, err := Local("nothing.return")
	Expect(err).To(HaveOccurred())
	Expect(nothing).To(Equal(""))
}

func TestLocal(t *testing.T) {
	RegisterTestingT(t)

	reset := withLocalGitConfigFile("remote.origin.url", "git@github.com:tcnksm/go-test-gitconfig.git")
	defer reset()

	var err error
	url, err := Local("remote.origin.url")
	Expect(err).NotTo(HaveOccurred())
	Expect(url).To(Equal("git@github.com:tcnksm/go-test-gitconfig.git"))

	nothing, err := Local("nothing.return")
	Expect(err).To(HaveOccurred())
	Expect(nothing).To(Equal(""))
}

func withGlobalGitConfigFile(content string) func() {
	tmpdir, err := ioutil.TempDir("", "go-gitconfig-test")
	if err != nil {
		panic(err)
	}

	tmpGitConfigFile := filepath.Join(tmpdir, ".gitconfig")

	ioutil.WriteFile(
		tmpGitConfigFile,
		[]byte(content),
		0777,
	)

	prevGitConfigEnv := os.Getenv("HOME")
	os.Setenv("HOME", tmpdir)

	return func() {
		os.Setenv("HOME", prevGitConfigEnv)
	}
}

func withLocalGitConfigFile(key string, value string) func() {
	var err error
	tmpdir, err := ioutil.TempDir(".", "go-gitconfig-test")
	if err != nil {
		panic(err)
	}

	prevDir, err := filepath.Abs(".")
	if err != nil {
		panic(err)
	}

	os.Chdir(tmpdir)

	gitInit := exec.Command("git", "init")
	gitInit.Stderr = ioutil.Discard
	if err = gitInit.Run(); err != nil {
		panic(err)
	}

	gitAddConfig := exec.Command("git", "config", "--local", key, value)
	gitAddConfig.Stderr = ioutil.Discard
	if err = gitAddConfig.Run(); err != nil {
		panic(err)
	}

	return func() {
		os.Chdir(prevDir)
		os.RemoveAll(tmpdir)
	}
}
