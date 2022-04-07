package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunCmd(t *testing.T) {
	t.Run("Errors", func(t *testing.T) {
		tests := []struct {
			cmd        []string
			env        Environment
			returnCode int
		}{
			{cmd: []string{"1", "2"}, env: nil, returnCode: 1},
			{cmd: []string{"lsls", "ls"}, env: nil, returnCode: 1},
			{cmd: []string{"ls"}, env: nil, returnCode: 0},
			{cmd: []string{"echo", "1"}, env: nil, returnCode: 0},
			{cmd: []string{"pwd"}, env: nil, returnCode: 0},
		}

		for _, test := range tests {
			exitCode := RunCmd(test.cmd, test.env)
			assert.Equal(t, exitCode, test.returnCode)
		}
	})
}
