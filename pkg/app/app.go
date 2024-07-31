package app

import (
	"context"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Controller struct {
	AppContext   *context.Context
	AppWaitGroup *sync.WaitGroup
	AppLogger    *zap.Logger
}

// InitialApp register signal handler and return app context and wait group
func InitialApp(options *LoggerOption) *Controller {
	var (
		Logger     *zap.Logger
		WaitGroup  sync.WaitGroup
		AppContext context.Context
		QuitApp    context.CancelFunc
	)

	// read options
	if options == nil {
		Logger = zap.NewExample()
	} else {
		Logger = createAppLogger(options)
	}

	// initialize app controller
	AppContext, QuitApp = context.WithCancel(context.Background())

	// monitor signal to exit app
	SignalChannel := make(chan os.Signal, 1)
	signal.Notify(SignalChannel, syscall.SIGINT, syscall.SIGTERM)
	WaitGroup.Add(1)
	go listenQuitSignal(SignalChannel, &WaitGroup, Logger, QuitApp)

	return &Controller{
		AppContext:   &AppContext,
		AppWaitGroup: &WaitGroup,
		AppLogger:    Logger,
	}
}

// listenQuitSignal listen quit signal and quit app
func listenQuitSignal(signal chan os.Signal, group *sync.WaitGroup, logger *zap.Logger, quit context.CancelFunc) {
	defer group.Done()
	defer syncLogger(logger)
	select {
	case <-signal:
		quit()
	}
}

// createAppLogger create customized app logger
