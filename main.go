package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/vdemeester/ko2dockerfile/internal/commands"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	if err := commands.Root.ExecuteContext(ctx); err != nil {
		os.Exit(1)
	}
}
