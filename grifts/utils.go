package grifts

import (
	"fmt"
	"os"
	"os/exec"
)

func execInShell(command string) error {
	cmd := exec.Command("/bin/bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func runCoverage(output string) error {
	err := execInShell("buffalo test -coverprofile=c.out ./...")
	if err != nil {
		return fmt.Errorf("failed to run tests: %v", err)
	}
	return execInShell(fmt.Sprintf("go tool cover -%s=c.out", output))
}
