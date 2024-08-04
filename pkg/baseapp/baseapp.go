package baseapp

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Cleaner interface {
	Stop()
}

type Controller struct {
	signalListener chan os.Signal
	cleaners       []Cleaner
	wg             *sync.WaitGroup
}

// InitApp handle the life cycle of the entire baseapp
func InitApp() (app *Controller) {
	listener := make(chan os.Signal, 1)
	signal.Notify(listener, syscall.SIGINT, syscall.SIGTERM)
	app = &Controller{
		signalListener: listener,
		wg:             &sync.WaitGroup{},
	}
	app.wg.Add(1)
	go func() {
		<-listener
		app.Exit()
	}()
	return
}

func (app *Controller) AppendCleaners(cleaners ...Cleaner) {
	for _, cleaner := range cleaners {
		app.cleaners = append(app.cleaners, cleaner)
	}
}

func (app *Controller) runAtExit() {
	defer app.wg.Done()
	for _, cleaner := range app.cleaners {
		cleaner.Stop()
	}
}

func (app *Controller) Exit() {
	app.runAtExit()
	os.Exit(0)
}

func (app *Controller) Loop() {
	app.wg.Wait()
}
