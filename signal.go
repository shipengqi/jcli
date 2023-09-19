package jcli

import (
	"os"
	"os/signal"
	"syscall"
)

type SignalReceiver func(os.Signal)

var setonce = make(chan struct{})

var sigc chan os.Signal

var defaultShutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}

// setupSignalHandler SIGTERM and SIGINT are registered by default.
// Register other signals via the signal parameter.
func (a *App) setupSignalHandler(receiver SignalReceiver, signals ...os.Signal) {
	close(setonce) // channel cannot be closed repeatedly, so panic occurs when called twice.

	if len(signals) == 0 {
		signals = defaultShutdownSignals
	}
	sigc = make(chan os.Signal)

	signal.Notify(sigc, signals...)

	go func() {
		for {
			if sig, ok := <-sigc; ok {
				a.logger.Debugf("%s received signal: %s", progressMessage, sig.String())
				receiver(sig)
			} else {
				a.logger.Debugf("%s signal channel closed", progressMessage)
				break
			}
		}
	}()
}
