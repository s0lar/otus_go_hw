package main

import (
	"log"
	"os"
	"os/exec"
	"syscall"
)

const defaultFailedCode = 1

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	for name, envEl := range env {
		if envEl.NeedRemove {
			if err := os.Unsetenv(name); err != nil {
				log.Fatal(err)
				return 1
			}
		} else {
			if err := os.Setenv(name, envEl.Value); err != nil {
				log.Fatal(err)
				return 1
			}
		}
	}

	command := exec.Command(cmd[0], cmd[1:]...) // #nosec G204
	command.Stdout = os.Stdout

	exitCode := 0

	err := command.Run()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok { //nolint:errorlint
			ws := exitError.Sys().(syscall.WaitStatus)
			exitCode = ws.ExitStatus()
		} else {
			exitCode = defaultFailedCode
		}
	}

	return exitCode
}
