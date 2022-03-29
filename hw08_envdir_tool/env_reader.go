package main

import (
	"bytes"
	"io/ioutil"
	"strings"
)

type Environment map[string]EnvValue

// EnvValue helps to distinguish between empty files and files with the first empty line.
type EnvValue struct {
	Value      string
	NeedRemove bool
}

func ReadDir(dir string) (Environment, error) {
	env := make(Environment)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return env, err
	}

	for _, file := range files {
		if strings.Contains(file.Name(), "=") {
			continue
		}
		if envValue, err := ReadFile(dir + "/" + file.Name()); err == nil {
			env[file.Name()] = envValue
		}
	}

	return env, nil
}

func ReadFile(fileName string) (EnvValue, error) {
	envValue := EnvValue{"", false}

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return envValue, err
	}

	if len(file) == 0 {
		envValue.NeedRemove = true
	}

	body := strings.Split(string(file), "\n")
	envValue.Value = strings.TrimRight(body[0], " 	")
	envValue.Value = string(bytes.Replace([]byte(envValue.Value), []byte("\x00"), []byte("\n"), 1))

	return envValue, nil
}
