package jcli

import (
	"os"
	"os/signal"
	"syscall"
)

var setonce = make(chan struct{})

var shutdownc chan os.Signal

var defaultShutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}

// SetupSignalHandler SIGTERM and SIGINT are registered by default.
// Register other signals via the signal parameter.
// A stop channel is returned which is closed on one of these signals.
// If a second signal is captured, the program will exit with code 1.
func SetupSignalHandler(signals ...os.Signal) <-chan struct{} {
	close(setonce) // channel cannot be closed repeatedly, so there will throw a panic when called twice

	if len(signals) == 0 {
		signals = defaultShutdownSignals
	}
	shutdownc = make(chan os.Signal, 2)

	stop := make(chan struct{})

	signal.Notify(shutdownc, signals...)

	go func() {
		<-shutdownc
		close(stop)
		<-shutdownc
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop
}
