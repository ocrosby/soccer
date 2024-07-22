package common

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func ExecuteCommand(command string) error {
	var (
		shell       string
		shellOption string
	)

	shellOption = "-c"
	if runtime.GOOS == "darwin" {
		shell = "bash"
	} else if runtime.GOOS == "windows" {
		shell = "cmd"
	} else {
		shell = "bash"
	}

	fmt.Printf("Running command: '%s'\n", command)

	cmd := exec.Command(shell, shellOption, command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("Error running command: %v", err)
	}

	return nil
}
