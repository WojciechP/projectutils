package gitinit

import (
	"fmt"
	"os/exec"
)

// LocalGit represents a local Git repository
type LocalGit struct {
	Path string
}

// Command executes a git command in the git directory (` cd REPO && git ...`)
func (git *LocalGit) Command(prog string, args ...string) *exec.Cmd {
	allArgs := make([]string, len(args)+1)
	allArgs[0] = prog
	for i := range args {
		allArgs[i+1] = args[i]
	}
	cmd := exec.Command("git", allArgs...)
	cmd.Dir = git.Path
	return cmd
}

// Init local Git repository (`git init`)
func (git *LocalGit) Init() error {
	out, err := git.Command("init", ".").Output()
	if err != nil {
		fmt.Printf("Failed to init repository at %s: %s\n%s\n", git.Path, err, out)
		return err
	}
	return nil
}

// AddRemote performs `git remote add`
func (git *LocalGit) AddRemote(name, url string) error {
	out, err := git.Command("remote", "add", name, url).Output()
	if err != nil {
		fmt.Printf("Failed to add remote repo: %s\n%s\n", err, out)
	}
	return nil

}
