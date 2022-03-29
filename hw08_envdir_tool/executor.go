package main

import (
	"log"
	"os"
	"os/exec"
)

// RunCmd runs a command + arguments (cmd) with environment variables from env.
func RunCmd(cmd []string, env Environment) (returnCode int) {
	for name, envEl := range env {
		if envEl.NeedRemove {
			if err := os.Unsetenv(name); err != nil {
				log.Fatal(err)
			}
		} else {
			if err := os.Setenv(name, envEl.Value); err != nil {
				log.Fatal(err)
			}
		}
	}

	command := exec.Command(cmd[0], cmd[1:]...) // #nosec G204
	command.Stdout = os.Stdout
	if err := command.Run(); err != nil {
		log.Fatal(err)
	}
	return
}
