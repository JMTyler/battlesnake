package main

import (
	"os"
	"os/signal"
	"syscall"
)

func WaitForKillSignal() {
	// Catch signal so we can shutdown gracefully
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
	<-sigCh
}
