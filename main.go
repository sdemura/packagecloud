package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/edgeworx/packagecloud/cmd"
)

func main() {
	ctx, cancelFn := context.WithCancel(context.Background())
	defer cancelFn() // nolint:gocritic

	go func() {
		stopCh := make(chan os.Signal, 1)
		signal.Notify(stopCh, os.Interrupt)

		<-stopCh
		cancelFn()
	}()

	err := cmd.Execute(ctx)
	if err != nil {
		cancelFn()
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1) // nolint:gocritic
	}
}
