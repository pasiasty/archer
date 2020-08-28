package grifts

import (
	"fmt"
	"os"
	"os/exec"

	. "github.com/markbates/grift/grift"
)

func execInShell(command string) error {
	cmd := exec.Command("/bin/bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

var _ = Desc("coverage", "Task Description")
var _ = Add("coverage", func(c *Context) error {
	err := execInShell("buffalo test -coverprofile=c.out ./...")
	if err != nil {
		return fmt.Errorf("failed to run tests: %v", err)
	}
	return execInShell("go tool cover -func=c.out")
})
