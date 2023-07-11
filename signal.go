package jcli

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/shipengqi/log"
)

type SignalReceiver func(os.Signal)

var setonce = make(chan struct{})

var sigc chan os.Signal

var defaultShutdownSignals = []os.Signal{os.Interrupt, syscall.SIGTERM}

// setupSignalHandler SIGTERM and SIGINT are registered by default.
// Register other signals via the signal parameter.
func setupSignalHandler(receiver SignalReceiver, signals ...os.Signal) {
	close(setonce) // channel cannot be closed repeatedly, so panic occurs when called twice.

	if len(signals) == 0 {
		signals = defaultShutdownSignals
	}
	sigc = make(chan os.Signal)

	signal.Notify(sigc, signals...)

	go func() {
		for {
			if sig, ok := <-sigc; ok {
				log.Debugf("%s received signal: %s", progressMessage, sig.String())
				receiver(sig)
			} else {
				log.Debugf("%s signal channel closed", progressMessage)
				break
			}
		}
	}()
}
