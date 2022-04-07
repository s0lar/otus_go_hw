package main

import (
	"log"
	"os"
)

func main() {
	if env, err := ReadDir(os.Args[1]); err != nil {
		log.Fatal(err)
	} else {
		RunCmd(os.Args[2:], env)
	}
}
