package main

import (
	"fmt"
	"github.com/JMTyler/battlesnake/_/config"
	"github.com/getsentry/sentry-go"
	"time"
)

func initSentry() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:         config.Get("sentry_dsn", ""),
		Environment: config.Get("environment", ""),
	})

	if err != nil {
		panic(err)
	}
}

func recoverWithSentry() {
	sentry.Recover()

	fmt.Println("Flushing Sentry...")
	sentry.Flush(time.Second)
}
