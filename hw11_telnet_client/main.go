package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "")
	flag.Parse()

	if flag.NArg() != 2 {
		log.Fatalln("Не задан хост и порт")
	}

	addr := strings.Join(flag.Args(), ":")
	tc := NewTelnetClient(addr, *timeout, os.Stdin, os.Stdout)

	if err := tc.Connect(); err != nil {
		log.Fatalln(err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	go func() {
		if err := tc.Receive(); err != nil {
			log.Fatalln(err)
		}
		fmt.Fprintln(os.Stderr, "...Connection was closed by peer")
		cancel()
	}()

	go func() {
		if err := tc.Send(); err != nil {
			log.Fatalln(err)
		}
		fmt.Fprintln(os.Stderr, "...EOF")
		cancel()
	}()

	<-ctx.Done()
}
