package main

import (
	"errors"
	"log"
	"os"
	"os/exec"
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
		var e *exec.ExitError
		if errors.As(err, &e) {
			exitCode = e.ExitCode()
		} else {
			exitCode = defaultFailedCode
		}
	}

	return exitCode
}
